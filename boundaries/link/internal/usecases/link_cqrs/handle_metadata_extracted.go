package link_cqrs

import (
	"context"
	"log/slog"

	metadatapb "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
)

const metadataExtractedTopic = "metadata.metadata.extracted.v1"

func (s *Service) subscribeToMetadataExtracted(ctx context.Context) error {
	return subscribe(ctx, s, metadataExtractedTopic, func() *metadatapb.MetadataExtracted {
		return &metadatapb.MetadataExtracted{}
	}, s.handleMetadataExtracted)
}

func (s *Service) handleMetadataExtracted(ctx context.Context, event *metadatapb.MetadataExtracted) error {
	if err := s.cqsStore.MetadataUpdate(ctx, event.GetId(), event.GetImageUrl(), event.GetDescription(), event.GetKeywords()); err != nil {
		s.log.ErrorWithContext(ctx, "Failed to update metadata in CQRS store",
			slog.String("error", err.Error()),
			slog.String("event_type", metadataExtractedTopic),
			slog.String("url", event.GetId()),
		)

		return err
	}

	s.log.InfoWithContext(ctx, "Processed MetadataExtracted event",
		slog.String("event_type", metadataExtractedTopic),
		slog.String("url", event.GetId()),
	)

	return nil
}
