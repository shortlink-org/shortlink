/*
Link Service. Application layer
*/
package link

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	"github.com/shortlink-org/shortlink/internal/pkg/saga"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud"
	queryStore "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
	metadata_rpc "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
)

type Service struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[v1.Link]

	// Delivery
	mq             *mq.DataBus
	MetadataClient metadata_rpc.MetadataServiceClient

	// Repository
	store *crud.Store

	logger logger.Logger
}

func New(logger logger.Logger, mq *mq.DataBus, metadataService metadata_rpc.MetadataServiceClient, store *crud.Store) (*Service, error) {
	service := &Service{
		logger: logger,

		// Delivery
		mq:             mq,
		MetadataClient: metadataService,

		// Repository
		store: store,
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

	// create a new saga for a get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.logger)).
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
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: hash}, Err: queryStore.ErrNotFound}
	}

	return resp, nil
}

func (s *Service) List(ctx context.Context, filter queryStore.Filter) (*v1.Links, error) {
	const (
		SAGA_NAME            = "LIST_LINK"
		SAGA_STEP_STORE_LIST = "SAGA_STEP_STORE_LIST"
	)

	resp := &v1.Links{}

	// create a new saga for a get list of a link
	sagaListLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.logger)).
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
	sagaAddLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: save to Store
	_, errs = sagaAddLink.AddStep(SAGA_STEP_STORE_SAVE).
		Then(func(ctx context.Context) error {
			var err error
			_, err = s.store.Add(ctx, in)

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
				Url: in.Url,
			})
			if err != nil {
				return err
			}

			return nil
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: publish event by this service
	_, errs = sagaAddLink.AddStep(SAGA_STEP_PUBLISH_EVENT_NEW_LINK).
		Then(func(ctx context.Context) error {
			data, err := proto.Marshal(in)
			if err != nil {
				return err
			}

			err = s.mq.Publish(ctx, v1.MQ_EVENT_LINK_CREATED, nil, data)
			if err != nil {
				return err
			}

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

func (s *Service) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

func (s *Service) Delete(ctx context.Context, hash string) (*v1.Link, error) {
	const (
		SAGA_NAME              = "DELETE_LINK"
		SAGA_STEP_STORE_DELETE = "SAGA_STEP_STORE_DELETE"
	)

	// create a new saga for a delete link by hash
	sagaDeleteLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.logger)).
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
