package metadata_mq

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	linkpb "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/domain/link/v1"
	"github.com/shortlink-org/go-sdk/cqrs/bus"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"
	"github.com/shortlink-org/go-sdk/logger"
	shortwatermill "github.com/shortlink-org/go-sdk/watermill"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain"
	domainerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/errors"
	infraerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/errors"
)

var (
	errInvalidEvent = errors.New("metadata mq: invalid event payload")
)

const linkCreatedEvent = domain.LinkCreatedTopic // Canonical event name (ADR-0002)

// SubscribeLinkCreated subscribes to link creation events from Kafka
// Uses ProtoMarshaler for automatic deserialization to *linkpb.LinkCreated
// Eliminates manual reflect.New - directly creates typed event instance
// Note: registry parameter kept for backward compatibility but not used (type is known from subscription)
func (e *Event) SubscribeLinkCreated(ctx context.Context, log logger.Logger, registry *bus.TypeRegistry, marshaler cqrsmessage.Marshaler) error {
	messages, err := e.subscriber.Subscribe(ctx, linkCreatedEvent)
	if err != nil {
		log.ErrorWithContext(ctx, "failed to subscribe to link created events",
			slog.String("error", err.Error()),
			slog.String("event", linkCreatedEvent),
		)

		return fmt.Errorf("subscribe to %s: %w", linkCreatedEvent, err)
	}

	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.ErrorWithContext(ctx, "panic in LinkCreated subscriber",
					slog.Any("recover", r),
					slog.String("topic", linkCreatedEvent),
				)
			}
		}()

		for msg := range messages {
			msgCtx := msg.Context() //nolint:contextcheck // inherit context from Watermill message
			if msgCtx == nil {
				msgCtx = ctx
			}

			// Restore trace context from message metadata so downstream spans stay in the same trace.
			msgCtx = shortwatermill.ExtractTrace(msgCtx, msg)
			msg.SetContext(msgCtx) //nolint:contextcheck // ensure downstream unmarshaler sees enriched ctx

			// Validate payload before unmarshaling
			if len(msg.Payload) == 0 {
				// Create span for empty payload error to track problematic messages in traces
				_, span := otel.Tracer("metadata.mq").Start(msgCtx, "metadata.mq.empty_payload_error",
					trace.WithSpanKind(trace.SpanKindConsumer),
				)
				span.SetStatus(otelcodes.Error, "Empty payload received")
				span.SetAttributes(
					attribute.String("messaging.system", "kafka"),
					attribute.String("messaging.destination.name", linkCreatedEvent),
					attribute.String("messaging.destination.kind", "topic"),
					attribute.String("messaging.message.id", msg.UUID),
					attribute.String("messaging.operation", "receive"),
					attribute.String("error.type", "empty_payload"),
				)
				span.End()

				log.ErrorWithContext(msgCtx, "Received empty payload for link created event - nacking for Kafka DLQ",
					slog.String("topic", linkCreatedEvent),
					slog.String("message_uuid", msg.UUID),
				)

				msg.Nack()
				continue
			}

			// Create typed event instance directly - no reflect.New needed
			// We know the type is *linkpb.LinkCreated from the subscription
			event := &linkpb.LinkCreated{}

			unmarshalErr := marshaler.Unmarshal(msg, event)
			if unmarshalErr != nil {
				// Create span for unmarshal error to track problematic messages in traces
				_, span := otel.Tracer("metadata.mq").Start(msgCtx, "metadata.mq.unmarshal_error",
					trace.WithSpanKind(trace.SpanKindConsumer),
				)
				span.RecordError(unmarshalErr)
				span.SetStatus(otelcodes.Error, unmarshalErr.Error())
				span.SetAttributes(
					attribute.String("messaging.system", "kafka"),
					attribute.String("messaging.destination.name", linkCreatedEvent),
					attribute.String("messaging.destination.kind", "topic"),
					attribute.String("messaging.message.id", msg.UUID),
					attribute.String("messaging.operation", "receive"),
					attribute.Int("messaging.message.payload_size_bytes", len(msg.Payload)),
					attribute.Int("messaging.message.metadata_count", len(msg.Metadata)),
					attribute.String("error.type", "unmarshal"),
				)
				span.End()

				// Nack() to allow Watermill DLQ to track retries and move to DLQ after max retries
				// Watermill will automatically move message to DLQ topic after WATERMILL_DLQ_MAX_RETRIES attempts
				// Log all metadata to debug DLQ retry tracking
				log.ErrorWithContext(msgCtx, "Failed to unmarshal event using marshaler - nacking for Kafka DLQ",
					slog.String("error", unmarshalErr.Error()),
					slog.String("topic", linkCreatedEvent),
					slog.Int("payload_size", len(msg.Payload)),
					slog.Int("metadata_count", len(msg.Metadata)),
					slog.String("message_uuid", msg.UUID),
					slog.Any("metadata", msg.Metadata), // Log all metadata to debug DLQ retry tracking
				)

				msg.Nack()
				continue
			}

			// Handle event - event is already typed as *linkpb.LinkCreated
			// msgCtx already contains the consumer span created automatically by otelsarama
			// Create metadata.process span as child of the automatic consumer span
			// This ensures proper trace hierarchy: automatic consumer span -> metadata.process -> saga
			processCtx, processSpan := otel.Tracer("metadata.uc").Start(msgCtx, "metadata.process",
				trace.WithSpanKind(trace.SpanKindInternal),
			)
			defer processSpan.End()

			processSpan.SetAttributes(
				attribute.String("link.url", event.GetUrl()),
				attribute.String("link.hash", event.GetHash()),
			)

			// Pass processCtx (which contains the automatic consumer span and process span) to handleLinkCreated
			handleErr := e.handleLinkCreated(processCtx, event, log) //nolint:contextcheck // metadata handling depends on message context
			if handleErr != nil {
				processSpan.RecordError(handleErr)
				processSpan.SetStatus(otelcodes.Error, handleErr.Error())
				// Automatic consumer span will automatically reflect error status from child span
				var domainErr *domainerrors.Error
				if errors.As(handleErr, &domainErr) {
					dto := infraerrors.FromDomainError("metadata.mq.link_created", domainErr)
					processSpan.SetAttributes(
						attribute.String("error.code", dto.Code),
						attribute.Bool("error.retryable", dto.Retryable),
					)
					log.ErrorWithContext(processCtx, "Failed to handle link created event - nacking for Kafka DLQ",
						slog.String("error_code", dto.Code),
						slog.String("topic", linkCreatedEvent),
						slog.Bool("retryable", dto.Retryable),
						slog.String("message", dto.Message),
					)
				} else {
					processSpan.SetAttributes(
						attribute.Bool("error.retryable", true),
					)
					log.ErrorWithContext(processCtx, "Failed to handle link created event - nacking for Kafka DLQ",
						slog.String("error", handleErr.Error()),
						slog.String("topic", linkCreatedEvent),
						slog.Bool("retryable", true),
					)
				}

				msg.Nack()

				continue
			}

			processSpan.SetStatus(otelcodes.Ok, "Metadata processed successfully")
			// Automatic consumer span status is automatically derived from child span
			msg.Ack()
		}
	}(ctx)

	return nil
}

// handleLinkCreated processes LinkCreated events
// Event is typed as *linkpb.LinkCreated
// Note: metadata.process span is already created in SubscribeLinkCreated as child of automatic consumer span
// This function receives ctx that contains the automatic consumer span and metadata.process span
func (e *Event) handleLinkCreated(ctx context.Context, event *linkpb.LinkCreated, log logger.Logger) error {
	linkURL := event.GetUrl()
	if linkURL == "" {
		log.ErrorWithContext(ctx, "Link URL is empty in event")
		return domainerrors.NewInvalidURLError("event.url", fmt.Errorf("empty url: %w", errInvalidEvent))
	}

	linkHash := event.GetHash()

	// Process metadata for the link URL
	// The context ctx already contains metadata.process span created in SubscribeLinkCreated
	_, err := e.metadataUC.Add(ctx, linkURL)
	if err != nil {
		log.ErrorWithContext(ctx, "Error processing metadata for link",
			slog.String("error", err.Error()),
			slog.String("url", linkURL),
			slog.String("hash", linkHash),
		)

		return err
	}

	log.InfoWithContext(ctx, "Successfully processed metadata for link",
		slog.String("url", linkURL),
		slog.String("hash", linkHash),
		slog.String("event_type", domain.LinkCreatedTopic),
	)

	return nil
}
