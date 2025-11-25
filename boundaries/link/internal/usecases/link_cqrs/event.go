package link_cqrs

import (
	"context"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// EventHandlers subscribes to CQRS events using ProtoMarshaler for automatic deserialization
func (s *Service) EventHandlers(ctx context.Context) error {
	// Subscribe to LinkCreated events
	if err := s.subscribeToLinkCreated(ctx); err != nil {
		return err
	}

	// Subscribe to LinkUpdated events
	if err := s.subscribeToLinkUpdated(ctx); err != nil {
		return err
	}

	// Subscribe to LinkDeleted events
	if err := s.subscribeToLinkDeleted(ctx); err != nil {
		return err
	}

	return nil
}

// subscribeToLinkCreated subscribes to LinkCreated events
// Uses ProtoMarshaler for automatic deserialization to *linkpb.LinkCreated
// Eliminates manual reflect.New - directly creates typed event instance
func (s *Service) subscribeToLinkCreated(ctx context.Context) error {
	messages, err := s.subscriber.Subscribe(ctx, domain.LinkCreatedTopic)
	if err != nil {
		s.log.Error("Failed to subscribe to LinkCreated events",
			slog.String("error", err.Error()),
			slog.String("topic", domain.LinkCreatedTopic),
		)
		return err
	}

	go func() {
		for msg := range messages {
			// Create typed event instance directly - no reflect.New needed
			event := &linkpb.LinkCreated{}

			// Unmarshal using ProtoMarshaler (handles metadata extraction)
			watermillMsg := message.NewMessage(msg.UUID, msg.Payload)
			watermillMsg.Metadata = msg.Metadata
			watermillMsg.SetContext(msg.Context())

			if err := s.marshaler.Unmarshal(watermillMsg, event); err != nil {
				s.log.Error("Failed to unmarshal LinkCreated event",
					slog.String("error", err.Error()),
					slog.String("topic", domain.LinkCreatedTopic),
				)
				msg.Nack()
				continue
			}

			// Handle event - event is already typed as *linkpb.LinkCreated
			if err := s.handleLinkCreated(msg.Context(), event); err != nil {
				s.log.Error("Failed to handle LinkCreated event",
					slog.String("error", err.Error()),
					slog.String("topic", domain.LinkCreatedTopic),
				)
				msg.Nack()
				continue
			}

			msg.Ack()
		}
	}()

	return nil
}

// subscribeToLinkUpdated subscribes to LinkUpdated events
// Uses ProtoMarshaler for automatic deserialization to *linkpb.LinkUpdated
func (s *Service) subscribeToLinkUpdated(ctx context.Context) error {
	messages, err := s.subscriber.Subscribe(ctx, domain.LinkUpdatedTopic)
	if err != nil {
		s.log.Error("Failed to subscribe to LinkUpdated events",
			slog.String("error", err.Error()),
			slog.String("topic", domain.LinkUpdatedTopic),
		)
		return err
	}

	go func() {
		for msg := range messages {
			// Create typed event instance directly
			event := &linkpb.LinkUpdated{}

			// Unmarshal using ProtoMarshaler
			watermillMsg := message.NewMessage(msg.UUID, msg.Payload)
			watermillMsg.Metadata = msg.Metadata
			watermillMsg.SetContext(msg.Context())

			if err := s.marshaler.Unmarshal(watermillMsg, event); err != nil {
				s.log.Error("Failed to unmarshal LinkUpdated event",
					slog.String("error", err.Error()),
					slog.String("topic", domain.LinkUpdatedTopic),
				)
				msg.Nack()
				continue
			}

			// Handle event
			if err := s.handleLinkUpdated(msg.Context(), event); err != nil {
				s.log.Error("Failed to handle LinkUpdated event",
					slog.String("error", err.Error()),
					slog.String("topic", domain.LinkUpdatedTopic),
				)
				msg.Nack()
				continue
			}

			msg.Ack()
		}
	}()

	return nil
}

// subscribeToLinkDeleted subscribes to LinkDeleted events
// Uses ProtoMarshaler for automatic deserialization to *linkpb.LinkDeleted
func (s *Service) subscribeToLinkDeleted(ctx context.Context) error {
	messages, err := s.subscriber.Subscribe(ctx, domain.LinkDeletedTopic)
	if err != nil {
		s.log.Error("Failed to subscribe to LinkDeleted events",
			slog.String("error", err.Error()),
			slog.String("topic", domain.LinkDeletedTopic),
		)
		return err
	}

	go func() {
		for msg := range messages {
			// Create typed event instance directly
			event := &linkpb.LinkDeleted{}

			// Unmarshal using ProtoMarshaler
			watermillMsg := message.NewMessage(msg.UUID, msg.Payload)
			watermillMsg.Metadata = msg.Metadata
			watermillMsg.SetContext(msg.Context())

			if err := s.marshaler.Unmarshal(watermillMsg, event); err != nil {
				s.log.Error("Failed to unmarshal LinkDeleted event",
					slog.String("error", err.Error()),
					slog.String("topic", domain.LinkDeletedTopic),
				)
				msg.Nack()
				continue
			}

			// Handle event
			if err := s.handleLinkDeleted(msg.Context(), event); err != nil {
				s.log.Error("Failed to handle LinkDeleted event",
					slog.String("error", err.Error()),
					slog.String("topic", domain.LinkDeletedTopic),
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
// Event is already typed as *linkpb.LinkCreated
func (s *Service) handleLinkCreated(ctx context.Context, event *linkpb.LinkCreated) error {

	// Convert protobuf event to domain Link
	// Hash is calculated automatically from URL in SetURL
	linkBuilder := domain.NewLinkBuilder().
		SetURL(event.GetUrl()).
		SetDescribe(event.GetDescribe())

	// Set timestamps if available
	if event.GetCreatedAt() != nil {
		linkBuilder = linkBuilder.SetCreatedAt(event.GetCreatedAt().AsTime())
	}
	if event.GetUpdatedAt() != nil {
		linkBuilder = linkBuilder.SetUpdatedAt(event.GetUpdatedAt().AsTime())
	}

	link, err := linkBuilder.Build()
	if err != nil {
		s.log.Error("Failed to build domain Link from event",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkCreatedTopic),
			slog.String("link_hash", event.GetHash()),
		)
		return err
	}

	// Store in CQRS store
	_, err = s.cqsStore.LinkAdd(ctx, link)
	if err != nil {
		s.log.Error("Failed to add link to CQRS store",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkCreatedTopic),
			slog.String("link_hash", event.GetHash()),
		)
		return err
	}

	s.log.Info("Processed LinkCreated event",
		slog.String("event_type", domain.LinkCreatedTopic),
		slog.String("link_hash", event.GetHash()),
	)

	return nil
}

// handleLinkUpdated processes LinkUpdated events
// Event is already typed as *linkpb.LinkUpdated
func (s *Service) handleLinkUpdated(ctx context.Context, event *linkpb.LinkUpdated) error {

	// Convert protobuf event to domain Link
	// Hash is calculated automatically from URL in SetURL
	linkBuilder := domain.NewLinkBuilder().
		SetURL(event.GetUrl()).
		SetDescribe(event.GetDescribe())

	// Set timestamps if available
	if event.GetCreatedAt() != nil {
		linkBuilder = linkBuilder.SetCreatedAt(event.GetCreatedAt().AsTime())
	}
	if event.GetUpdatedAt() != nil {
		linkBuilder = linkBuilder.SetUpdatedAt(event.GetUpdatedAt().AsTime())
	}

	link, err := linkBuilder.Build()
	if err != nil {
		s.log.Error("Failed to build domain Link from event",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkUpdatedTopic),
			slog.String("link_hash", event.GetHash()),
		)
		return err
	}

	// Update in CQRS store
	_, err = s.cqsStore.LinkUpdate(ctx, link)
	if err != nil {
		s.log.Error("Failed to update link in CQRS store",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkUpdatedTopic),
			slog.String("link_hash", event.GetHash()),
		)
		return err
	}

	s.log.Info("Processed LinkUpdated event",
		slog.String("event_type", domain.LinkUpdatedTopic),
		slog.String("link_hash", event.GetHash()),
	)

	return nil
}

// handleLinkDeleted processes LinkDeleted events
// Event is already typed as *linkpb.LinkDeleted
func (s *Service) handleLinkDeleted(ctx context.Context, event *linkpb.LinkDeleted) error {

	if err := s.cqsStore.LinkDelete(ctx, event.GetHash()); err != nil {
		s.log.Error("Failed to delete link from CQRS store",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkDeletedTopic),
			slog.String("link_hash", event.GetHash()),
		)
		return err
	}

	s.log.Info("Processed LinkDeleted event",
		slog.String("event_type", domain.LinkDeletedTopic),
		slog.String("link_hash", event.GetHash()),
	)

	return nil
}
