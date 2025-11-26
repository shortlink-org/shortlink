package metadata_mq

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	linkpb "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/domain/link/v1"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/shortlink-org/go-sdk/cqrs/bus"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"
	"github.com/shortlink-org/go-sdk/logger"

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
		for msg := range messages {
			msgCtx := msg.Context() //nolint:contextcheck // inherit context from Watermill message
			if msgCtx == nil {
				msgCtx = ctx
			}

			// Create typed event instance directly - no reflect.New needed
			// We know the type is *linkpb.LinkCreated from the subscription
			event := &linkpb.LinkCreated{}

			// Unmarshal using ProtoMarshaler (handles metadata extraction)
			watermillMsg := message.NewMessage(msg.UUID, msg.Payload)
			watermillMsg.Metadata = msg.Metadata
			watermillMsg.SetContext(msgCtx) //nolint:contextcheck // watermill message already carries derived context

			unmarshalErr := marshaler.Unmarshal(watermillMsg, event)
			if unmarshalErr != nil {
				log.ErrorWithContext(msgCtx, "Failed to unmarshal event using marshaler", //nolint:contextcheck // logging must use message context
					slog.String("error", unmarshalErr.Error()),
					slog.String("topic", linkCreatedEvent),
				)
				e.archiveRawEvent(msgCtx, msg, linkCreatedEvent, fmt.Sprintf("unmarshal: %v", unmarshalErr))
				msg.Nack()

				continue
			}

			// Handle event - event is already typed as *linkpb.LinkCreated
			handleErr := e.handleLinkCreated(msgCtx, event, log) //nolint:contextcheck // metadata handling depends on message context
			if handleErr != nil {
				var domainErr *domainerrors.Error
				if errors.As(handleErr, &domainErr) {
					dto := infraerrors.FromDomainError("metadata.mq.link_created", domainErr)
					log.ErrorWithContext(msgCtx, "Failed to handle link created event",
						slog.String("error_code", dto.Code),
						slog.String("topic", linkCreatedEvent),
						slog.Bool("retryable", dto.Retryable),
						slog.String("message", dto.Message),
					)
					e.archiveRawEvent(msgCtx, msg, linkCreatedEvent, fmt.Sprintf("domain: %s", dto.Message))
				} else {
					log.ErrorWithContext(msgCtx, "Failed to handle link created event",
						slog.String("error", handleErr.Error()),
						slog.String("topic", linkCreatedEvent),
					)
					e.archiveRawEvent(msgCtx, msg, linkCreatedEvent, fmt.Sprintf("handler: %v", handleErr))
				}

				msg.Nack()

				continue
			}

			msg.Ack()
		}
	}(ctx)

	return nil
}

// handleLinkCreated processes LinkCreated events
// Event is typed as *linkpb.LinkCreated - use proto reflection for field access
// (buf.build generated code may not have direct getters)
func (e *Event) handleLinkCreated(ctx context.Context, event *linkpb.LinkCreated, log logger.Logger) error {
	// Use proto reflection to access fields (buf.build may not generate getters)
	eventReflect := event.ProtoReflect()
	urlField := eventReflect.Descriptor().Fields().ByName("url")
	hashField := eventReflect.Descriptor().Fields().ByName("hash")

	if urlField == nil {
		log.ErrorWithContext(ctx, "URL field not found in LinkCreated event")
		return domainerrors.NewInvalidURLError("event.url", fmt.Errorf("missing url field: %w", errInvalidEvent))
	}

	linkURL := eventReflect.Get(urlField).String()
	if linkURL == "" {
		log.ErrorWithContext(ctx, "Link URL is empty in event")
		return domainerrors.NewInvalidURLError("event.url", fmt.Errorf("empty url: %w", errInvalidEvent))
	}

	var linkHash string
	if hashField != nil {
		linkHash = eventReflect.Get(hashField).String()
	}

	// Process metadata for the link URL
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

func (e *Event) archiveRawEvent(ctx context.Context, msg *message.Message, topic, reason string) {
	if e == nil || e.rawEvents == nil || msg == nil {
		return
	}

	_ = e.rawEvents.Save(ctx, RawEventRecord{
		Topic:    topic,
		Payload:  cloneBytes(msg.Payload),
		Metadata: cloneMetadata(msg.Metadata),
		Error:    reason,
	})
}
