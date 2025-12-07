package link

import (
	"context"
	"errors"
	"log/slog"
	"time"

	permission "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/shortlink-org/go-sdk/auth/session"
	"github.com/shortlink-org/go-sdk/saga"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
		permissionWriteTimeout           = 5 * time.Second
	)

	// Observability
	NewLinkHistogramObserve(ctx)

	userID, err := session.GetUserID(ctx)
	if err != nil {
		uc.log.ErrorWithContext(ctx, "failed to get user ID from session",
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
			ctx, span := otel.Tracer("link.uc.store").Start(ctx, "link.store.add",
				trace.WithSpanKind(trace.SpanKindInternal),
			)
			defer span.End()

			span.SetAttributes(
				attribute.String("saga.step", SAGA_STEP_SAVE_TO_STORE),
				attribute.String("link.hash", in.GetHash()),
				attribute.String("user.id", userID),
			)

			_, err := uc.store.Add(ctx, in)
			if err != nil {
				span.RecordError(err)
				span.SetStatus(otelcodes.Error, err.Error())
				return err
			}

			return nil
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

			permissionCtx, cancel := context.WithTimeout(ctx, permissionWriteTimeout)
			defer cancel()

			permissionCtx, span := otel.Tracer("link.uc.permission").Start(permissionCtx, "authzed.api.v1.PermissionsService/WriteRelationships",
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

			_, err := uc.permission.PermissionsServiceClient.WriteRelationships(permissionCtx, relationship)
			if err != nil {
				st, ok := status.FromError(err)
				if ok {
					span.SetAttributes(attribute.String("grpc.code", st.Code().String()))

					switch st.Code() {
					case codes.AlreadyExists:
						span.SetAttributes(attribute.Bool("permission.already_exists", true))
						span.AddEvent("permission relationship already existed")

						return nil
					default:
						if int32(st.Code()) == int32(permission.ErrorReason_ERROR_REASON_TOO_MANY_PRECONDITIONS_IN_REQUEST) {
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
			linkData := &dto.LinkData{
				URL:       in.GetUrl().String(),
				Hash:      in.GetHash(),
				Describe:  in.GetDescribe(),
				CreatedAt: in.GetCreatedAt().GetTime(),
				UpdatedAt: in.GetUpdatedAt().GetTime(),
			}

			// Convert LinkData to LinkCreated event using DTO
			event := dto.ToLinkCreatedEvent(linkData)
			if event == nil {
				uc.log.ErrorWithContext(ctx, "Link creation event is nil",
					slog.String("saga.step", SAGA_STEP_PUBLISH_EVENT_NEW_LINK),
					slog.String("link_hash", in.GetHash()),
				)
				return domain.NewInternalError("link creation event is nil")
			}

			// Publish event using EventBus
			// Producer span will be created automatically by instrumentation (otelsarama/watermill)
			err := uc.eventBus.Publish(ctx, event)
			if err != nil {
				uc.log.ErrorWithContext(ctx, "Failed to publish link creation event",
					slog.String("error", err.Error()),
					slog.String("saga.step", SAGA_STEP_PUBLISH_EVENT_NEW_LINK),
					slog.String("event_type", domain.LinkCreatedTopic),
					slog.String("link_hash", in.GetHash()),
					slog.String("user.id", userID),
				)

				return err
			}

			uc.log.InfoWithContext(ctx, "Link creation event published successfully",
				slog.String("saga.step", SAGA_STEP_PUBLISH_EVENT_NEW_LINK),
				slog.String("event_type", domain.LinkCreatedTopic),
				slog.String("link_hash", in.GetHash()),
				slog.String("user.id", userID),
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
