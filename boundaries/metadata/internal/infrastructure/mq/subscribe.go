package metadata_mq

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/cqrs/bus"

	// Import Link events from buf.build
	linkpb "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/domain/link/v1"

	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain"
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
		return err
	}

	go func() {
		for msg := range messages {
			// Create typed event instance directly - no reflect.New needed
			// We know the type is *linkpb.LinkCreated from the subscription
			event := &linkpb.LinkCreated{}

			// Unmarshal using ProtoMarshaler (handles metadata extraction)
			watermillMsg := message.NewMessage(msg.UUID, msg.Payload)
			watermillMsg.Metadata = msg.Metadata
			watermillMsg.SetContext(msg.Context())

			if err := marshaler.Unmarshal(watermillMsg, event); err != nil {
				log.ErrorWithContext(msg.Context(), "Failed to unmarshal event using marshaler",
					slog.String("error", err.Error()),
					slog.String("topic", linkCreatedEvent),
				)
				msg.Nack()
				continue
			}

			// Handle event - event is already typed as *linkpb.LinkCreated
			if err := e.handleLinkCreated(msg.Context(), event, log); err != nil {
				log.ErrorWithContext(msg.Context(), "Failed to handle link created event",
					slog.String("error", err.Error()),
					slog.String("topic", linkCreatedEvent),
				)
				msg.Nack()
				continue
			}

			msg.Ack()
		}
	}()

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
		return errors.New("URL field not found")
	}

	linkURL := eventReflect.Get(urlField).String()
	if linkURL == "" {
		log.ErrorWithContext(ctx, "Link URL is empty in event")
		return errors.New("link URL is empty")
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
