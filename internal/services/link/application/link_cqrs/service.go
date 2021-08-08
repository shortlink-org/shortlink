package link_cqrs

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/notify"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/cqs"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/query"
	"github.com/batazor/shortlink/pkg/saga"
)

type Service struct {
	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event

	// Repository
	cqsStore   *cqs.Store
	queryStore *query.Store

	logger logger.Logger
}

func New(logger logger.Logger, cqsStore *cqs.Store, queryStore *query.Store) (*Service, error) {
	service := &Service{
		cqsStore:   cqsStore,
		queryStore: queryStore,

		logger: logger,
	}

	// Subscribe to event
	service.EventHandler()

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
		SAGA_NAME           = "GET_LINK_CQRS"
		SAGA_STEP_STORE_GET = "SAGA_STEP_STORE_GET_CQRS"
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
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: hash}, Err: fmt.Errorf("Not found links")}
	}

	return resp, nil
}
