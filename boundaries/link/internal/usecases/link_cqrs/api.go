package link_cqrs

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/saga"
	link "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link_cqrs/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/types/v1"
)

func errorHelper(ctx context.Context, log logger.Logger, errs []error) error {
	if len(errs) > 0 {
		attrs := make([]slog.Attr, 0, len(errs))
		for index, err := range errs {
			attrs = append(attrs, slog.Any(fmt.Sprintf("stack error: %d", index), err))
		}

		log.ErrorWithContext(ctx, "Error create a new link", attrs...)

		return ErrCreateLink
	}

	return nil
}

func (s *Service) Get(ctx context.Context, hash string) (*domain.LinkView, error) {
	const (
		SAGA_NAME           = "GET_LINK_CQRS"
		SAGA_STEP_STORE_GET = "SAGA_STEP_STORE_GET_CQRS"
	)

	resp := &domain.LinkView{}

	// create a new saga for a get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
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
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaGetLink.Play(nil)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, &link.NotFoundByHashError{Hash: hash}
	}

	return resp, nil
}

func (s *Service) List(ctx context.Context, filter *v1.FilterLink) (*domain.LinksView, error) {
	const (
		SAGA_NAME           = "GET_LINKS_CQRS"
		SAGA_STEP_STORE_GET = "SAGA_STEP_STORE_GET_CQRS"
	)

	resp := &domain.LinksView{}

	// create a new saga for a get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
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
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaGetLink.Play(nil)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, ErrNotFound
	}

	return resp, nil
}
