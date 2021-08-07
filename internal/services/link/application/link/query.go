package link

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
	"github.com/batazor/shortlink/pkg/saga"
)

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
			resp, err = s.Store.Get(ctx, hash)
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
	filter := &query.Filter{}

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
			resp, err = s.Store.List(ctx, filter)
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
