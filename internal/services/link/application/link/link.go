/*
Link Service. Application layer
*/
package link

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/notify"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud"
	queryStore "github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	"github.com/batazor/shortlink/pkg/saga"
)

type Service struct {
	// Observer interface for subscribe on system event
	notify.Subscriber

	// Delivery
	MetadataClient metadata_rpc.MetadataServiceClient

	// Repository
	store *crud.Store

	logger logger.Logger
}

func New(logger logger.Logger, metadataService metadata_rpc.MetadataServiceClient, store *crud.Store) (*Service, error) {
	service := &Service{
		MetadataClient: metadataService,

		store: store,

		logger: logger,
	}

	return service, nil
}

func errorHelper(ctx context.Context, logger logger.Logger, errs []error) error {
	if len(errs) > 0 {
		errList := field.Fields{}
		for index := range errs {
			errList[fmt.Sprintf("stack error: %d", index)] = errs[index]
		}

		logger.ErrorWithContext(ctx, "Error create a new link", errList)
		return fmt.Errorf("Error create a new link")
	}

	return nil
}

func (s *Service) Get(ctx context.Context, hash string) (*v1.Link, error) {
	const (
		SAGA_NAME           = "GET_LINK"
		SAGA_STEP_STORE_GET = "SAGA_STEP_STORE_GET"
	)

	resp := &v1.Link{}

	// create a new saga for get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: get link from store
	_, errs = sagaGetLink.AddStep(SAGA_STEP_STORE_GET).
		Then(func(ctx context.Context) error {
			var err error
			resp, err = s.store.Get(ctx, hash)
			return err
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaGetLink.Play(nil)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: hash}, Err: fmt.Errorf("Not found links")}
	}

	return resp, nil
}

func (s *Service) List(ctx context.Context, in string) (*v1.Links, error) {
	// Parse args
	filter := queryStore.Filter{}

	if in != "" {
		errJsonUnmarshal := json.Unmarshal([]byte(in), &filter)
		if errJsonUnmarshal != nil {
			return nil, errors.New("error parse payload as string")
		}
	}

	const (
		SAGA_NAME            = "LIST_LINK"
		SAGA_STEP_STORE_LIST = "SAGA_STEP_STORE_LIST"
	)

	resp := &v1.Links{}

	// create a new saga for get list of link
	sagaListLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: get link from store
	_, errs = sagaListLink.AddStep(SAGA_STEP_STORE_LIST).
		Then(func(ctx context.Context) error {
			var err error
			resp, err = s.store.List(ctx, &filter)
			return err
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaListLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) Add(ctx context.Context, in *v1.Link) (*v1.Link, error) {
	const (
		SAGA_NAME                        = "ADD_LINK"
		SAGA_STEP_STORE_SAVE             = "SAGA_STEP_STORE_SAVE"
		SAGA_STEP_METADATA_GET           = "SAGA_STEP_METADATA_GET"
		SAGA_STEP_PUBLISH_EVENT_NEW_LINK = "SAGA_STEP_PUBLISH_EVENT_NEW_LINK"
	)

	// saga for create a new link
	sagaAddLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: save to Store
	_, errs = sagaAddLink.AddStep(SAGA_STEP_STORE_SAVE).
		Then(func(ctx context.Context) error {
			var err error
			in, err = s.store.Add(ctx, in)
			return err
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: request to metadata
	_, errs = sagaAddLink.AddStep(SAGA_STEP_METADATA_GET).
		Then(func(ctx context.Context) error {
			_, err := s.MetadataClient.Set(ctx, &metadata_rpc.MetadataServiceSetRequest{
				Id: in.Url,
			})
			return err
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: publish event by this service
	_, errs = sagaAddLink.AddStep(SAGA_STEP_PUBLISH_EVENT_NEW_LINK).
		Then(func(ctx context.Context) error {
			notify.Publish(ctx, v1.METHOD_ADD, in, nil)
			return nil
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaAddLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Service) Update(ctx context.Context, in *v1.Link) (*v1.Link, error) {
	return nil, nil
}

func (s *Service) Delete(ctx context.Context, hash string) (*v1.Link, error) {
	const (
		SAGA_NAME              = "DELETE_LINK"
		SAGA_STEP_STORE_DELETE = "SAGA_STEP_STORE_DELETE"
	)

	// create a new saga for delete link by hash
	sagaDeleteLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: get link from store
	_, errs = sagaDeleteLink.AddStep(SAGA_STEP_STORE_DELETE).
		Then(func(ctx context.Context) error {
			return s.store.Delete(ctx, hash)
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaDeleteLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
