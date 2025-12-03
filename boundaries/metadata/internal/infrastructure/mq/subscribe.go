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
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain"
	domainerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/errors"
)

var (
	errInvalidEvent = errors.New("metadata mq: invalid event payload")
)

const linkCreatedEvent = domain.LinkCreatedTopic // Canonical event name (ADR-0002)

// SubscribeLinkCreated subscribes to link creation events from Kafka
// Uses ProtoMarshaler for automatic deserialization to *linkpb.LinkCreated
// Eliminates manual reflect.New - directly creates typed event instance
// Note: registry parameter kept for backward compatibility but not used (type is known from subscription)
func (e *Event) SubscribeLinkCreated(
	ctx context.Context,
	log logger.Logger,
	registry *bus.TypeRegistry,
	marshaler cqrsmessage.Marshaler,
) error {

	messages, err := e.subscriber.Subscribe(ctx, linkCreatedEvent)
	if err != nil {
		log.ErrorWithContext(ctx, "failed to subscribe to link created events",
			slog.String("error", err.Error()),
			slog.String("event", linkCreatedEvent),
		)
		return fmt.Errorf("subscribe to %s: %w", linkCreatedEvent, err)
	}

	go func() {
		for msg := range messages {
			// 1) Get message context (contains consumer span created by otelsarama)
			// otelsarama automatically extracts traceparent from Kafka RecordHeaders
			// and creates a consumer span linked to the producer span
			msgCtx := msg.Context()
			if msgCtx == nil {
				msgCtx = ctx
			}

			// 2) Extract trace context from message metadata
			// This ensures proper parent-child relationship between producer and consumer spans
			// shortwatermill.ExtractTrace handles the complete trace propagation
			msgCtx = shortwatermill.ExtractTrace(msgCtx, msg)
			msg.SetContext(msgCtx)

			// 3) Validate payload
			if len(msg.Payload) == 0 {
				_, span := otel.Tracer("metadata.mq").Start(
					msgCtx,
					"metadata.empty_payload",
					trace.WithSpanKind(trace.SpanKindInternal),
				)
				span.SetStatus(otelcodes.Error, "empty payload")
				span.End()

				msg.Nack()
				continue
			}

			// 4) Parse event payload
			event := &linkpb.LinkCreated{}
			if err := marshaler.Unmarshal(msg, event); err != nil {
				_, span := otel.Tracer("metadata.mq").Start(
					msgCtx,
					"metadata.unmarshal_error",
					trace.WithSpanKind(trace.SpanKindInternal),
				)
				span.RecordError(err)
				span.SetStatus(otelcodes.Error, err.Error())
				span.End()

				msg.Nack()
				continue
			}

			// 5) Start processing span as a child of the otelsarama consumer span
			processCtx, processSpan := otel.Tracer("metadata.process").Start(
				msgCtx,
				"metadata.process",
				trace.WithSpanKind(trace.SpanKindConsumer),
			)

			// 6) Execute metadata use case (fix: Add returns (value, error))
			_, err = e.metadataUC.Add(processCtx, event.GetUrl())
			if err != nil {
				processSpan.RecordError(err)
				processSpan.SetStatus(otelcodes.Error, err.Error())
				processSpan.End()

				msg.Nack()
				continue
			}

			// 7) Success
			processSpan.SetStatus(otelcodes.Ok, "")
			processSpan.End()

			msg.Ack()
		}
	}()

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
