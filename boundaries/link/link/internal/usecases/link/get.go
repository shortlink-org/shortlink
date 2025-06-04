package link

import (
	"context"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/pkg/auth/session"
	"github.com/shortlink-org/shortlink/pkg/pattern/saga"
)

// Get - get a link by hash
//
// Saga:
// 1. Check permission
// 2. Get a link from store
func (uc *UC) Get(ctx context.Context, hash string) (*domain.Link, error) {
	const (
		SAGA_NAME                  = "GET_LINK"
		SAGA_STEP_CHECK_PERMISSION = "SAGA_STEP_CHECK_PERMISSION"
		SAGA_STEP_GET_FROM_STORE   = "SAGA_STEP_GET_FROM_STORE"
	)

	userID := session.GetUserID(ctx)
	resp := &domain.Link{}

	// create a new saga for a get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.SetLogger(uc.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaGetLink.AddStep(SAGA_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			relationship := &permission.CheckPermissionRequest{
				Resource:   &permission.ObjectReference{ObjectType: "link", ObjectId: hash},
				Permission: "view",
				Subject:    &permission.SubjectReference{Object: &permission.ObjectReference{ObjectType: "user", ObjectId: userID}},
			}

			_, err := uc.permission.PermissionsServiceClient.CheckPermission(ctx, relationship)
			if err != nil {
				return err
			}

			return nil
		}).Reject(func(ctx context.Context, thenErr error) error {
		return &domain.PermissionDeniedError{Err: thenErr}
	}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaGetLink.AddStep(SAGA_STEP_GET_FROM_STORE).
		Needs(SAGA_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			var err error
			resp, err = uc.store.Get(ctx, hash)
			if err != nil {
				return err
			}

			return nil
		}).Reject(func(ctx context.Context, thenErr error) error {
		return &domain.PermissionDeniedError{Err: thenErr}
	}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaGetLink.Play(nil)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, NotFoundByHash{Hash: hash}
	}

	return resp, nil
}
