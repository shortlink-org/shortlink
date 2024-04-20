package link

import (
	"context"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/pkg/auth/session"
	"github.com/shortlink-org/shortlink/pkg/pattern/saga"
)

// Delete - delete link
//
// Saga:
// 1. Check permission
// 2. Delete from store
func (uc *UC) Delete(ctx context.Context, hash string) (*domain.Link, error) {
	const (
		SAGA_NAME                   = "DELETE_LINK"
		SAGE_STEP_CHECK_PERMISSION  = "SAGE_STEP_CHECK_PERMISSION"
		SAGA_STEP_DELETE_FROM_STORE = "SAGA_STEP_DELETE_FROM_STORE"
	)

	userID := session.GetUserID(ctx)

	// create a new saga for a delete link by hash
	sagaDeleteLink, errs := saga.New(SAGA_NAME, saga.SetLogger(uc.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaDeleteLink.AddStep(SAGE_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			_, err := uc.permission.DeleteRelationships(ctx, &permission.DeleteRelationshipsRequest{
				RelationshipFilter: &permission.RelationshipFilter{
					ResourceType:       "link",
					OptionalResourceId: hash,
					OptionalRelation:   "writer",
					OptionalSubjectFilter: &permission.SubjectFilter{
						SubjectType:       "user",
						OptionalSubjectId: userID,
					},
				},
			})
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

	_, errs = sagaDeleteLink.AddStep(SAGA_STEP_DELETE_FROM_STORE).
		Needs(SAGE_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			return uc.store.Delete(ctx, hash)
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaDeleteLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
