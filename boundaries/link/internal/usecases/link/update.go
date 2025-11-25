package link

import (
	"context"
	"log/slog"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/dto"
)

func (uc *UC) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	resp, err := uc.store.Update(ctx, in)
	if err != nil {
		return nil, err
	}

	// Publish LinkUpdated event
	linkData := dto.LinkData{
		URL:       resp.GetUrl().String(),
		Hash:      resp.GetHash(),
		Describe:  resp.GetDescribe(),
		CreatedAt: resp.GetCreatedAt().GetTime(),
		UpdatedAt: resp.GetUpdatedAt().GetTime(),
	}

	event := dto.ToLinkUpdatedEvent(linkData)
	if err := uc.eventBus.Publish(ctx, event); err != nil {
		uc.log.Error("Failed to publish link updated event",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.LinkUpdatedTopic),
			slog.String("link_hash", resp.GetHash()),
		)
		// Don't fail the update if event publishing fails
	} else {
		uc.log.Info("Link updated event published successfully",
			slog.String("event_type", domain.LinkUpdatedTopic),
			slog.String("link_hash", resp.GetHash()),
		)
	}

	return resp, nil
}
