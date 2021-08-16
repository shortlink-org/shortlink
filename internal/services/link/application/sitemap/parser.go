package sitemap

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogo/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/logger"
	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	domain "github.com/batazor/shortlink/internal/services/link/domain/sitemap/v1"
)

type Service struct {
	logger logger.Logger

	// Delivery
	mq mq.MQ
}

func New(logger logger.Logger, mq mq.MQ) (*Service, error) {
	service := &Service{
		logger: logger,

		// Delivery
		mq: mq,
	}

	return service, nil
}

func (s *Service) Parse(ctx context.Context, url string) error {
	// Request the HTML page.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(`Incorrect response code: %d for %s`, resp.StatusCode, url)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var payload domain.Sitemap
	err = xml.Unmarshal(bodyBytes, &payload)
	if err != nil {
		return err
	}

	// send to link_rpc.add
	for key := range payload.Url {
		data, errMarshal := proto.Marshal(&link.Link{Url: payload.Url[key].Loc})
		if errMarshal != nil {
			return errMarshal
		}

		errPublish := s.mq.Publish(ctx, link.MQ_EVENT_LINK_NEW, query.Message{
			Key:     nil,
			Payload: data,
		})
		if errPublish != nil {
			return errPublish
		}
	}

	return nil
}
