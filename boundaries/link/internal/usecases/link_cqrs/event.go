package link_cqrs

import (
	"context"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	infraerrors "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/errors"
)

// EventHandlers subscribes to CQRS events using ProtoMarshaler for automatic deserialization
func (s *Service) EventHandlers(ctx context.Context) error {
	if err := s.subscribeToLinkCreated(ctx); err != nil {
		return err
	}

	if err := s.subscribeToLinkUpdated(ctx); err != nil {
		return err
	}

	if err := s.subscribeToLinkDeleted(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Service) subscribeToLinkCreated(ctx context.Context) error {
	return subscribe(ctx, s, domain.LinkCreatedTopic, func() *domain.LinkCreated {
		return &domain.LinkCreated{}
	}, s.handleLinkCreated)
}

func (s *Service) subscribeToLinkUpdated(ctx context.Context) error {
	return subscribe(ctx, s, domain.LinkUpdatedTopic, func() *domain.LinkUpdated {
		return &domain.LinkUpdated{}
	}, s.handleLinkUpdated)
}

func (s *Service) subscribeToLinkDeleted(ctx context.Context) error {
	return subscribe(ctx, s, domain.LinkDeletedTopic, func() *domain.LinkDeleted {
		return &domain.LinkDeleted{}
	}, s.handleLinkDeleted)
}

func subscribe[T proto.Message](
	ctx context.Context,
	s *Service,
	topic string,
	newEvent func() T,
	handler func(context.Context, T) error,
) error {
	messages, err := s.subscriber.Subscribe(ctx, topic)
	if err != nil {
		s.log.ErrorWithContext(ctx, "Failed to subscribe to events",
			slog.String("error", err.Error()),
			slog.String("topic", topic),
		)

		return err
	}

	go func(ctx context.Context) {
		for msg := range messages {
			msgCtx := msg.Context() //nolint:contextcheck // reuse message context from Watermill
			if msgCtx == nil {
				msgCtx = ctx
			}

			event := newEvent()

			watermillMsg := message.NewMessage(msg.UUID, msg.Payload)
			watermillMsg.Metadata = msg.Metadata
			watermillMsg.SetContext(msgCtx) //nolint:contextcheck // inherit context from message or parent

			if err := s.marshaler.Unmarshal(watermillMsg, event); err != nil {
				s.log.ErrorWithContext(msgCtx, "Failed to unmarshal event",
					slog.String("error", err.Error()),
					slog.String("topic", topic),
				)
				msg.Nack()

				continue
			}

			if err := handler(msgCtx, event); err != nil {
				dto := infraerrors.FromDomainError("link.cqrs", err)

				s.log.ErrorWithContext(msgCtx, "Failed to handle event",
					slog.String("error_code", dto.Code),
					slog.Bool("retryable", dto.Retryable),
					slog.String("error", dto.Message),
					slog.String("topic", topic),
				)

				if dto.Retryable {
					msg.Nack()
				} else {
					msg.Ack()
				}

				continue
			}

			msg.Ack()
		}
	}(ctx)

	return nil
}
