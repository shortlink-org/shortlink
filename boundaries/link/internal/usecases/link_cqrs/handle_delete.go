package link_cqrs

import (
	"context"
	"log/slog"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// handleLinkDeleted processes LinkDeleted events
// Event is already typed as *domain.LinkDeleted
func (s *Service) handleLinkDeleted(ctx context.Context, event *domain.LinkDeleted) error {
	if err := s.cqsStore.LinkDelete(ctx, event.GetHash()); err != nil {
		s.log.ErrorWithContext(ctx, "Failed to delete link from CQRS store",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkDeletedTopic),
			slog.String("link_hash", event.GetHash()),
		)

		return err
	}

	s.log.InfoWithContext(ctx, "Processed LinkDeleted event",
		slog.String("event_type", domain.LinkDeletedTopic),
		slog.String("link_hash", event.GetHash()),
	)

	return nil
}
