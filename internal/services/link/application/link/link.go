/*
Link Service. Application layer
*/
package link

import (
	"context"
	"errors"
	"fmt"
	"io"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/shortlink-org/shortlink/internal/pkg/auth/session"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	"github.com/shortlink-org/shortlink/internal/pkg/saga"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud"
	queryStore "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
	metadata_rpc "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
)

type Service struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[v1.Link]

	// Security
	permission *authzed.Client

	// Delivery
	mq             mq.MQ
	MetadataClient metadata_rpc.MetadataServiceClient

	// Repository
	store crud.Repository

	log logger.Logger
}

func New(log logger.Logger, dataBus mq.MQ, metadataService metadata_rpc.MetadataServiceClient, store crud.Repository, permissionClient *authzed.Client) (*Service, error) {
	service := &Service{
		log: log,

		// Security
		permission: permissionClient,

		// Delivery
		mq:             dataBus,
		MetadataClient: metadataService,

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
func (s *Service) Get(ctx context.Context, hash string) (*v1.Link, error) {
	const (
		SAGA_NAME                  = "GET_LINK"
		SAGA_STEP_CHECK_PERMISSION = "SAGA_STEP_CHECK_PERMISSION"
		SAGA_STEP_GET_FROM_STORE   = "SAGA_STEP_GET_FROM_STORE"
	)

	sess := session.GetSession(ctx)
	resp := &v1.Link{}

	// create a new saga for a get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaGetLink.AddStep(SAGA_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			relationship := &permission.CheckPermissionRequest{
				Resource:   &permission.ObjectReference{ObjectType: "link", ObjectId: hash},
				Permission: "reader",
				Subject:    &permission.SubjectReference{Object: &permission.ObjectReference{ObjectType: "user", ObjectId: sess.GetId()}},
			}

			_, err := s.permission.PermissionsServiceClient.CheckPermission(ctx, relationship)
			if err != nil {
				return err
			}

			return nil
		}).Reject(func(ctx context.Context, thenErr error) error {
		return &v1.PermissionDeniedError{Err: thenErr}
	}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaGetLink.AddStep(SAGA_STEP_GET_FROM_STORE).
		Needs(SAGA_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			var err error
			resp, err = s.store.Get(ctx, hash)
			if err != nil {
				return err
			}

			return nil
		}).Reject(func(ctx context.Context, thenErr error) error {
		return &v1.PermissionDeniedError{Err: thenErr}
	}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaGetLink.Play(nil)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: hash}}
	}

	return resp, nil
}

// List - get a list of links
//
// Saga:
// 1. Check permission
// 2. Get a list of links from store
func (s *Service) List(ctx context.Context, filter queryStore.Filter) (*v1.Links, error) {
	const (
		SAGA_NAME                     = "LIST_LINK"
		SAGA_STEP_LOOKUP              = "SAGA_STEP_LOOKUP"
		SAGA_STEP_GET_LIST_FROM_STORE = "SAGA_STEP_GET_LIST_FROM_STORE"
	)

	userID := session.GetUserID(ctx)
	links := &v1.Links{}

	// create a new saga for a get list of a link
	sagaListLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaListLink.AddStep(SAGA_STEP_LOOKUP).
		Then(func(ctx context.Context) error {
			relationship := &permission.LookupResourcesRequest{
				ResourceObjectType: "link",
				Permission:         "reader",
				Subject:            &permission.SubjectReference{Object: &permission.ObjectReference{ObjectType: "user", ObjectId: userID}},
			}

			stream, err := s.permission.PermissionsServiceClient.LookupResources(ctx, relationship)
			if err != nil {
				return err
			}

			resources := []*permission.LookupResourcesResponse{}
			for {
				resp, errRead := stream.Recv()
				if errRead != nil {
					if errors.Is(errRead, io.EOF) {
						return nil
					}

					return errRead
				}

				resources = append(resources, resp) //nolint:staticcheck // use it later
			}

			// TODO: use filter
			// *filter.Link.Contains = list.String()
		}).Reject(func(ctx context.Context, thenErr error) error {
		return &v1.PermissionDeniedError{Err: thenErr}
	}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaListLink.AddStep(SAGA_STEP_GET_LIST_FROM_STORE).
		Then(func(ctx context.Context) error {
			var err error
			links, err = s.store.List(ctx, &filter)

			return err
		}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaListLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return links, nil
}

// Add - create a new link
//
// Saga:
// 1. Save to store
// 2. Add permission
// 3. Get metadata
// 4. Publish event
func (s *Service) Add(ctx context.Context, in *v1.Link) (*v1.Link, error) {
	const (
		SAGA_NAME                        = "ADD_LINK"
		SAGA_STEP_ADD_PERMISSION         = "SAGA_STEP_ADD_PERMISSION"
		SAGA_STEP_SAVE_TO_STORE          = "SAGA_STEP_SAVE_TO_STORE"
		SAGA_STEP_GET_METADATA           = "SAGA_STEP_GET_METADATA"
		SAGA_STEP_PUBLISH_EVENT_NEW_LINK = "SAGA_STEP_PUBLISH_EVENT_NEW_LINK"
	)

	userID := session.GetUserID(ctx)

	// saga for create a new link
	sagaAddLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaAddLink.
		AddStep(SAGA_STEP_SAVE_TO_STORE).
		Then(func(ctx context.Context) error {
			var err error
			_, err = s.store.Add(ctx, in)

			return err
		}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
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

			_, err := s.permission.PermissionsServiceClient.WriteRelationships(ctx, relationship)
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
		err := s.store.Delete(ctx, in.GetHash())
		if err != nil {
			return err
		}

		return nil
	}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaAddLink.AddStep(SAGA_STEP_GET_METADATA).
		Needs(SAGA_STEP_ADD_PERMISSION).
		Then(func(ctx context.Context) error {
			_, err := s.MetadataClient.Set(ctx, &metadata_rpc.MetadataServiceSetRequest{
				Url: in.GetUrl(),
			})
			if err != nil {
				// TODO:
				// 1. Move to metadata service
				// 2. Listen MQ event

				return nil //nolint:nilerr // ignore
			}

			return nil
		}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaAddLink.AddStep(SAGA_STEP_PUBLISH_EVENT_NEW_LINK).
		Then(func(ctx context.Context) error {
			// If mq is nil, then we don't need to publish event
			if s.mq == nil {
				return nil
			}

			data, err := proto.Marshal(in)
			if err != nil {
				return err
			}

			err = s.mq.Publish(ctx, v1.MQ_EVENT_LINK_CREATED, nil, data)
			if err != nil {
				return err
			}

			return nil
		}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaAddLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Service) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete - delete link
//
// Saga:
// 1. Check permission
// 2. Delete from store
func (s *Service) Delete(ctx context.Context, hash string) (*v1.Link, error) {
	const (
		SAGA_NAME                   = "DELETE_LINK"
		SAGE_STEP_CHECK_PERMISSION  = "SAGE_STEP_CHECK_PERMISSION"
		SAGA_STEP_DELETE_FROM_STORE = "SAGA_STEP_DELETE_FROM_STORE"
	)

	sess := session.GetSession(ctx)

	// create a new saga for a delete link by hash
	sagaDeleteLink, errs := saga.New(SAGA_NAME, saga.SetLogger(s.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaDeleteLink.AddStep(SAGE_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			_, err := s.permission.DeleteRelationships(ctx, &permission.DeleteRelationshipsRequest{
				RelationshipFilter: &permission.RelationshipFilter{
					ResourceType:       "link",
					OptionalResourceId: hash,
					OptionalRelation:   "writer",
					OptionalSubjectFilter: &permission.SubjectFilter{
						SubjectType:       "user",
						OptionalSubjectId: sess.GetId(),
					},
				},
			})
			if err != nil {
				return err
			}

			return nil
		}).Reject(func(ctx context.Context, thenErr error) error {
		return &v1.PermissionDeniedError{Err: thenErr}
	}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaDeleteLink.AddStep(SAGA_STEP_DELETE_FROM_STORE).
		Needs(SAGE_STEP_CHECK_PERMISSION).
		Then(func(ctx context.Context) error {
			return s.store.Delete(ctx, hash)
		}).Build()
	if err := errorHelper(ctx, s.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaDeleteLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
