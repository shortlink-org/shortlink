package link

import (
	"context"
	"errors"
	"log/slog"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shortlink-org/go-sdk/auth/session"
	"github.com/shortlink-org/go-sdk/saga"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/dto"
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
		SAGA_STEP_PUBLISH_EVENT_NEW_LINK = "SAGA_STEP_PUBLISH_EVENT_NEW_LINK"
	)

	// Observability
	NewLinkHistogramObserve(ctx)

	userID, err := session.GetUserID(ctx)
	if err != nil {
		uc.log.Error("failed to get user ID from session",
			slog.String("error", err.Error()),
		)

		return nil, err
	}

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
					Operation: permission.RelationshipUpdate_OPERATION_TOUCH,
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

			ctx, span := otel.Tracer("link.uc.permission").Start(ctx, "authzed.api.v1.PermissionsService/WriteRelationships",
				trace.WithSpanKind(trace.SpanKindClient),
			)
			defer span.End()

			span.SetAttributes(
				attribute.String("rpc.system", "grpc"),
				attribute.String("rpc.service", "authzed.api.v1.PermissionsService"),
				attribute.String("rpc.method", "WriteRelationships"),
				attribute.String("link.hash", in.GetHash()),
				attribute.String("user.id", userID),
			)

			_, err := uc.permission.PermissionsServiceClient.WriteRelationships(ctx, relationship)
			if err != nil {
				st, ok := status.FromError(err)
				if ok {
					span.SetAttributes(attribute.String("grpc.code", st.Code().String()))
					switch st.Code() {
					case codes.AlreadyExists:
						span.SetAttributes(attribute.Bool("permission.already_exists", true))
						span.AddEvent("permission relationship already existed")
						span.SetStatus(otelcodes.Ok, "relationship already exists")

						return nil
					default:
						if int32(st.Code()) == int32(permission.ErrorReason_ERROR_REASON_TOO_MANY_PRECONDITIONS_IN_REQUEST) {
							span.SetStatus(otelcodes.Ok, "skipped: too many preconditions")

							return nil
						}
					}

					span.RecordError(err)
					span.SetStatus(otelcodes.Error, st.Message())
				} else {
					span.RecordError(err)
					span.SetStatus(otelcodes.Error, err.Error())
				}

				return err
			}

			span.SetAttributes(
				attribute.Bool("permission.already_exists", false),
				attribute.String("grpc.code", codes.OK.String()),
			)
			span.SetStatus(otelcodes.Ok, "relationship created")

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

	_, errs = sagaAddLink.AddStep(SAGA_STEP_PUBLISH_EVENT_NEW_LINK).
		Then(func(ctx context.Context) error {
			// Convert domain Link to LinkData (avoids import cycle)
			linkData := dto.LinkData{
				URL:       in.GetUrl().String(),
				Hash:      in.GetHash(),
				Describe:  in.GetDescribe(),
				CreatedAt: in.GetCreatedAt().GetTime(),
				UpdatedAt: in.GetUpdatedAt().GetTime(),
			}

			// Convert LinkData to LinkCreated event using DTO
			event := dto.ToLinkCreatedEvent(linkData)

			// Publish event using EventBus
			if err := uc.eventBus.Publish(ctx, event); err != nil {
				uc.log.Error("Failed to publish link creation event",
					slog.String("error", err.Error()),
					slog.String("event_type", domain.LinkCreatedTopic),
					slog.String("link_hash", in.GetHash()),
				)
				return err
			}

			uc.log.Info("Link creation event published successfully",
				slog.String("event_type", domain.LinkCreatedTopic),
				slog.String("link_hash", in.GetHash()),
			)

			return nil
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err = sagaAddLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return in, nil
}
