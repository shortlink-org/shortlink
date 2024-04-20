/*
Link UC. Application layer
*/
package link

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/status"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/auth/session"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/notify"
	"github.com/shortlink-org/shortlink/pkg/pattern/saga"
)

type UC struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[domain.Link]

	// Security
	permission *authzed.Client

	// Delivery
	mq mq.MQ
	// MetadataClient metadata_rpc.MetadataServiceClient

	// Repository
	store crud.Repository

	log logger.Logger
}

func New(log logger.Logger, dataBus mq.MQ, metadataService any, store crud.Repository, permissionClient *authzed.Client) (*UC, error) {
	service := &UC{
		log: log,

		// Security
		permission: permissionClient,

		// Delivery
		mq: dataBus,
		// MetadataClient: metadataService,

		// Repository
		store: store,
	}

	return service, nil
}

func errorHelper(ctx context.Context, log logger.Logger, errs []error) error {
	if len(errs) > 0 {
		errList := field.Fields{}
		for index := range errs {
			errList[fmt.Sprintf("stack error: %d", index)] = errs[index]
		}

		log.ErrorWithContext(ctx, "Error create a new link", errList)

		return ErrCreateLink
	}

	return nil
}

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

// Add - create a new link
//
// Saga:
// 1. Save to store
// 2. Add permission
// 3. Get metadata
// 4. Publish event
func (uc *UC) Add(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	const (
		SAGA_NAME                        = "ADD_LINK"
		SAGA_STEP_ADD_PERMISSION         = "SAGA_STEP_ADD_PERMISSION"
		SAGA_STEP_SAVE_TO_STORE          = "SAGA_STEP_SAVE_TO_STORE"
		SAGA_STEP_GET_METADATA           = "SAGA_STEP_GET_METADATA"
		SAGA_STEP_PUBLISH_EVENT_NEW_LINK = "SAGA_STEP_PUBLISH_EVENT_NEW_LINK"
	)

	// Observability
	NewLinkHistogramObserve(ctx)

	userID := session.GetUserID(ctx)

	// saga for create a new link
	sagaAddLink, errs := saga.New(SAGA_NAME, saga.SetLogger(uc.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaAddLink.
		AddStep(SAGA_STEP_SAVE_TO_STORE).
		Then(func(ctx context.Context) error {
			var err error
			_, err = uc.store.Add(ctx, in)

			return err
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaAddLink.AddStep(SAGA_STEP_ADD_PERMISSION).
		Needs(SAGA_STEP_SAVE_TO_STORE).
		Then(func(ctx context.Context) error {
			relationship := &permission.WriteRelationshipsRequest{
				Updates: []*permission.RelationshipUpdate{{
					Operation: permission.RelationshipUpdate_OPERATION_CREATE,
					Relationship: &permission.Relationship{
						Resource: &permission.ObjectReference{
							ObjectType: "link",
							ObjectId:   in.GetHash(),
						},
						Relation: "writer",
						Subject: &permission.SubjectReference{
							Object: &permission.ObjectReference{
								ObjectType: "user",
								ObjectId:   userID,
							},
						},
					},
				}},
			}

			_, err := uc.permission.PermissionsServiceClient.WriteRelationships(ctx, relationship)
			if err != nil {
				st, ok := status.FromError(err)
				if ok {
					if int32(st.Code()) == int32(permission.ErrorReason_ERROR_REASON_TOO_MANY_PRECONDITIONS_IN_REQUEST) {
						return nil
					}
				}

				return err
			}

			return nil
		}).Reject(func(ctx context.Context, thenErr error) error {
		err := uc.store.Delete(ctx, in.GetHash())
		if err != nil {
			return errors.Join(thenErr, err)
		}

		return thenErr
	}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaAddLink.AddStep(SAGA_STEP_GET_METADATA).
		Needs(SAGA_STEP_ADD_PERMISSION).
		Then(func(ctx context.Context) error {
			// link := in.GetUrl()
			// _, err := uc.MetadataClient.Set(ctx, &metadata_rpc.MetadataServiceSetRequest{
			// 	Url: link.String(),
			// })
			// if err != nil {
			// 	// TODO:
			// 	// 1. Move to metadata service
			// 	// 2. Listen MQ event
			//
			// 	return nil //nolint:nilerr // ignore
			// }

			return nil
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaAddLink.AddStep(SAGA_STEP_PUBLISH_EVENT_NEW_LINK).
		Then(func(ctx context.Context) error {
			// If mq is a nil, then we don't need to publish event
			if uc.mq == nil {
				return nil
			}

			data, err := json.Marshal(in)
			if err != nil {
				return err
			}

			err = uc.mq.Publish(ctx, domain.MQ_EVENT_LINK_CREATED, nil, data)
			if err != nil {
				return err
			}

			return nil
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaAddLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (uc *UC) Update(_ context.Context, _ *domain.Link) (*domain.Link, error) {
	return nil, nil
}

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
