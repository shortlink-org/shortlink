package link

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/notify"
	api_domain "github.com/batazor/shortlink/internal/services/api/domain"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
	"github.com/batazor/shortlink/pkg/saga"
)

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
			in, err = s.cqsStore.Store.Add(ctx, in)
			return err
		}).
		Reject(func(ctx context.Context) error {
			return s.cqsStore.Store.Delete(ctx, in.Hash)
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: request to metadata
	_, errs = sagaAddLink.AddStep(SAGA_STEP_METADATA_GET).
		Then(func(ctx context.Context) error {
			_, err := s.MetadataClient.Set(ctx, &metadata_rpc.SetMetaRequest{
				Id: in.Url,
			})
			return err
		}).
		Reject(func(ctx context.Context) error {
			return s.cqsStore.Store.Delete(ctx, in.Hash)
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: publish event by this service
	_, errs = sagaAddLink.AddStep(SAGA_STEP_PUBLISH_EVENT_NEW_LINK).
		Then(func(ctx context.Context) error {
			notify.Publish(ctx, api_domain.METHOD_ADD, in, nil)
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
			return s.cqsStore.Store.Delete(ctx, hash)
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
