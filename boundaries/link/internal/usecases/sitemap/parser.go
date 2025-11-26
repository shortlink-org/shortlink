package sitemap

import (
	"context"
	"encoding/xml"
	"io"
	"log/slog"
	"net/http"

	"github.com/shortlink-org/go-sdk/cqrs/bus"
	http_client "github.com/shortlink-org/go-sdk/http/client"
	"github.com/shortlink-org/go-sdk/logger"

	link "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/sitemap/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/dto"
)

type Service struct {
	log logger.Logger

	// Delivery
	eventBus *bus.EventBus // CQRS EventBus for publishing events
}

func New(log logger.Logger, eventBus *bus.EventBus) (*Service, error) {
	service := &Service{
		log:      log,
		eventBus: eventBus,
	}

	return service, nil
}

func (s *Service) Parse(ctx context.Context, url string) error {
	// Request the HTML page.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	client, err := http_client.New()
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck // ignore close error

	if resp.StatusCode != http.StatusOK {
		return &IncorrectResponseCodeError{
			StatusCode: resp.StatusCode,
			URL:        url,
		}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var payload domain.Sitemap

	err = xml.Unmarshal(bodyBytes, &payload)
	if err != nil {
		return err
	}

	// Publish LinkCreated events for each URL in sitemap
	for key := range payload.GetUrl() {
		newLink, err := link.NewLinkBuilder().
			SetURL(payload.GetUrl()[key].GetLoc()).
			Build()
		if err != nil {
			s.log.ErrorWithContext(ctx, "Failed to build link from sitemap URL",
				slog.String("error", err.Error()),
				slog.String("url", payload.GetUrl()[key].GetLoc()),
			)

			continue
		}

		// Convert domain Link to LinkData
		linkData := &dto.LinkData{
			URL:       newLink.GetUrl().String(),
			Hash:      newLink.GetHash(),
			Describe:  newLink.GetDescribe(),
			CreatedAt: newLink.GetCreatedAt().GetTime(),
			UpdatedAt: newLink.GetUpdatedAt().GetTime(),
		}

		// Convert LinkData to LinkCreated event using DTO
		event := dto.ToLinkCreatedEvent(linkData)
		if event == nil {
			s.log.ErrorWithContext(ctx, "Failed to build link creation event from sitemap",
				slog.String("link_hash", newLink.GetHash()),
				slog.String("url", payload.GetUrl()[key].GetLoc()),
			)

			continue
		}

		// Publish event using EventBus (canonical name: link.link.created.v1)
		if err := s.eventBus.Publish(ctx, event); err != nil {
			s.log.ErrorWithContext(ctx, "Failed to publish link creation event from sitemap",
				slog.String("error", err.Error()),
				slog.String("event_type", "link.link.created.v1"),
				slog.String("link_hash", newLink.GetHash()),
				slog.String("url", payload.GetUrl()[key].GetLoc()),
			)

			continue
		}

		s.log.InfoWithContext(ctx, "Link creation event published from sitemap",
			slog.String("event_type", "link.link.created.v1"),
			slog.String("link_hash", newLink.GetHash()),
			slog.String("url", payload.GetUrl()[key].GetLoc()),
		)
	}

	return nil
}
