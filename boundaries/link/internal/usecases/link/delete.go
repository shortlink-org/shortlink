package link

import (
	"context"
	"log/slog"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/shortlink-org/go-sdk/auth/session"
	"github.com/shortlink-org/go-sdk/saga"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/dto"
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

	userID, err := session.GetUserID(ctx)
	if err != nil {
		uc.log.ErrorWithContext(ctx, "failed to get user ID from session",
			slog.String("error", err.Error()),
		)

		return nil, err
	}

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
		return domain.ErrPermissionDenied(thenErr)
	}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaDeleteLink.AddStep(SAGA_STEP_DELETE_FROM_STORE).
		Needs(SAGE_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			return uc.store.Delete(ctx, hash)
		}).Reject(func(ctx context.Context, err error) error {
		return err
	}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err = sagaDeleteLink.Play(nil)
	if err != nil {
		return nil, err
	}

	// Publish LinkDeleted event
	event := dto.ToLinkDeletedEvent(hash)
	if err := uc.eventBus.Publish(ctx, event); err != nil {
		uc.log.ErrorWithContext(ctx, "Failed to publish link deleted event",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkDeletedTopic),
			slog.String("link_hash", hash),
		)
		// Don't fail the delete if event publishing fails
	} else {
		uc.log.InfoWithContext(ctx, "Link deleted event published successfully",
			slog.String("event_type", domain.LinkDeletedTopic),
			slog.String("link_hash", hash),
		)
	}

	return nil, nil
}
