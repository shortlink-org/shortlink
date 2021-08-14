package sitemap

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"golang.org/x/sync/errgroup"

	"github.com/batazor/shortlink/internal/pkg/logger"
	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	v12 "github.com/batazor/shortlink/internal/services/link/domain/sitemap/v1"
)

type Service struct {
	logger logger.Logger

	// Delivery
	mq mq.MQ
}

func New(logger logger.Logger, mq mq.MQ) (*Service, error) {
	service := &Service{
		mq: mq,

		logger: logger,
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

	var payload v12.Sitemap
	err = xml.Unmarshal(bodyBytes, &payload)
	if err != nil {
		return err
	}

	// send to link_rpc.add
	g := errgroup.Group{}
	for key := range payload.URL {
		g.Go(func() error {
			data, errMarshal := proto.Marshal(&link.Link{Url: payload.URL[key].Loc})
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

			return nil
		})
	}

	err = g.Wait()
	if err != nil {
		return err
	}

	return nil
}
