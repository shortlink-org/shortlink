package link_cqrs

import (
	"context"
	"fmt"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/saga"
	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link_cqrs/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

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

func (s *Service) Get(ctx context.Context, hash string) (*domain.LinkView, error) {
	const (
		SAGA_NAME           = "GET_LINK_CQRS"
		SAGA_STEP_STORE_GET = "SAGA_STEP_STORE_GET_CQRS"
	)

	resp := &domain.LinkView{}

	// create a new saga for get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: get a link from store
	_, errs = sagaGetLink.AddStep(SAGA_STEP_STORE_GET).
		Then(func(ctx context.Context) error {
			var err error
			resp, err = s.queryStore.Get(ctx, hash)

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
		return nil, &link.NotFoundError{Link: &link.Link{Hash: hash}, Err: query.ErrNotFound}
	}

	return resp, nil
}

func (s *Service) List(ctx context.Context, filter *query.Filter) (*domain.LinksView, error) {
	const (
		SAGA_NAME           = "GET_LINKS_CQRS"
		SAGA_STEP_STORE_GET = "SAGA_STEP_STORE_GET_CQRS"
	)

	resp := &domain.LinksView{}

	// create a new saga for get link by hash
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
			resp, err = s.queryStore.List(ctx, filter)

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
		return nil, &link.NotFoundError{Link: &link.Link{Hash: ""}, Err: query.ErrNotFound}
	}

	return resp, nil
}
