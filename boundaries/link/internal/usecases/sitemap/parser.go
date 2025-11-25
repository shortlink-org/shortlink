package sitemap

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	http_client "github.com/shortlink-org/go-sdk/http/client"
	"github.com/shortlink-org/go-sdk/logger"
	link "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/sitemap/v1"
)

type Service struct {
	log logger.Logger

	// Delivery
	publisher message.Publisher
}

func New(log logger.Logger, publisher message.Publisher) (*Service, error) {
	service := &Service{
		log: log,

		// Delivery
		publisher: publisher,
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

	// send to link_rpc.add
	for key := range payload.GetUrl() {
		newLink, err := link.NewLinkBuilder().
			SetURL(payload.GetUrl()[key].GetLoc()).
			Build()
		if err != nil {
			return err
		}

		data, errMarshal := json.Marshal(newLink)
		if errMarshal != nil {
			return errMarshal
		}

		msg := message.NewMessage(watermill.NewUUID(), data)
		msg.Metadata.Set("event_type", link.MQ_EVENT_LINK_NEW)

		errPublish := s.publisher.Publish(link.MQ_EVENT_LINK_NEW, msg)
		if errPublish != nil {
			return errPublish
		}
	}

	return nil
}
