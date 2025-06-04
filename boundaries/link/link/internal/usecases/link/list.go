package link

import (
	"context"
	"errors"
	"io"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"go.opentelemetry.io/otel/trace"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/auth/session"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/pattern/saga"
)

// List - get a list of links
//
// Saga:
// 1. Check permission
// 2. Get a list of links from store
func (uc *UC) List(ctx context.Context, filter *types.FilterLink, cursor string, limit uint32) (*domain.Links, *string, error) {
	const (
		SAGA_NAME                     = "LIST_LINK"
		SAGA_STEP_LOOKUP              = "SAGA_STEP_LOOKUP"
		SAGA_STEP_GET_LIST_FROM_STORE = "SAGA_STEP_GET_LIST_FROM_STORE"
	)

	// Set default values
	userID := session.GetUserID(ctx)
	links := &domain.Links{}
	nextToken := ""

	if filter == nil {
		filter = &types.FilterLink{}
	}

	if filter.Hash == nil {
		filter.Hash = &types.StringFilterInput{}
	}

	// create a new saga for a get list of a link
	sagaListLink, errs := saga.New(SAGA_NAME, saga.SetLogger(uc.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, nil, err
	}

	_, errs = sagaListLink.AddStep(SAGA_STEP_LOOKUP).
		Then(func(ctx context.Context) error {
			relationship := &permission.LookupResourcesRequest{
				ResourceObjectType: "link",
				Permission:         "view",
				Subject:            &permission.SubjectReference{Object: &permission.ObjectReference{ObjectType: "user", ObjectId: userID}},
				OptionalLimit:      limit,
			}

			if cursor != "" {
				relationship.OptionalCursor = &permission.Cursor{
					Token: cursor,
				}
			}

			stream, err := uc.permission.PermissionsServiceClient.LookupResources(ctx, relationship)
			if err != nil {
				return err
			}

			for {
				resp, errRead := stream.Recv()
				if errRead != nil {
					if errors.Is(errRead, io.EOF) {
						return nil
					}

					// add error to span
					span := trace.SpanFromContext(ctx)
					span.RecordError(errRead)

					return errRead
				}

				// Set token for pagination
				nextToken = resp.GetAfterResultCursor().GetToken()

				// Add hash to filter
				filter.Hash.Contains = append(filter.Hash.Contains, resp.GetResourceObjectId())
			}
		}).Reject(func(ctx context.Context, thenErr error) error {
		return &domain.PermissionDeniedError{Err: thenErr}
	}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, nil, err
	}

	_, errs = sagaListLink.AddStep(SAGA_STEP_GET_LIST_FROM_STORE).
		Needs(SAGA_STEP_LOOKUP).
		Then(func(ctx context.Context) error {
			var err error

			links, err = uc.store.List(ctx, filter)
			if err != nil {
				return err
			}

			return nil
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, nil, err
	}

	// Run saga
	err := sagaListLink.Play(nil)
	if err != nil {
		uc.log.ErrorWithContext(ctx, "Error get list of links", field.Fields{"error": err})
		return nil, nil, err
	}

	return links, &nextToken, nil
}
