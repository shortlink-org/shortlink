package sitemap

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/segmentio/encoding/json"

	link "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/sitemap/v1"
	http_client "github.com/shortlink-org/shortlink/pkg/http/client"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
)

type Service struct {
	log logger.Logger

	// Delivery
	mq mq.MQ
}

func New(log logger.Logger, dataBus mq.MQ) (*Service, error) {
	service := &Service{
		log: log,

		// Delivery
		mq: dataBus,
	}

	return service, nil
}

func (s *Service) Parse(ctx context.Context, url string) error {
	// Request the HTML page.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	client := http_client.New()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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

		errPublish := s.mq.Publish(ctx, link.MQ_EVENT_LINK_NEW, nil, data)
		if errPublish != nil {
			return errPublish
		}
	}

	return nil
}
