package link

import (
	"context"
	"errors"

	"github.com/segmentio/encoding/json"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"google.golang.org/grpc/status"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/pkg/auth/session"
	"github.com/shortlink-org/shortlink/pkg/pattern/saga"
)

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
