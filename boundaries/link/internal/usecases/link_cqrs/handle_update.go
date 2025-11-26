package link_cqrs

import (
	"context"
	"log/slog"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// handleLinkUpdated processes LinkUpdated events
// Event is already typed as *domain.LinkUpdated
func (s *Service) handleLinkUpdated(ctx context.Context, event *domain.LinkUpdated) error {
	linkBuilder := domain.NewLinkBuilder().
		SetURL(event.GetUrl()).
		SetDescribe(event.GetDescribe())

	if createdAt := event.GetCreatedAt(); createdAt != nil {
		linkBuilder = linkBuilder.SetCreatedAt(createdAt.AsTime())
	}

	if updatedAt := event.GetUpdatedAt(); updatedAt != nil {
		linkBuilder = linkBuilder.SetUpdatedAt(updatedAt.AsTime())
	}

	link, err := linkBuilder.Build()
	if err != nil {
		s.log.ErrorWithContext(ctx, "Failed to build domain Link from link updated event",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkUpdatedTopic),
			slog.String("link_hash", event.GetHash()),
		)

		return err
	}

	if _, err := s.cqsStore.LinkUpdate(ctx, link); err != nil {
		s.log.ErrorWithContext(ctx, "Failed to update link in CQRS store",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkUpdatedTopic),
			slog.String("link_hash", event.GetHash()),
		)

		return err
	}

	s.log.InfoWithContext(ctx, "Processed LinkUpdated event",
		slog.String("event_type", domain.LinkUpdatedTopic),
		slog.String("link_hash", event.GetHash()),
	)

	return nil
}
