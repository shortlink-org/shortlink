package sitemap

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/batazor/shortlink/internal/pkg/logger"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	v12 "github.com/batazor/shortlink/internal/services/link/domain/sitemap/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
)

type Service struct {
	logger logger.Logger

	// Delivery
	LinkServiceClient link_rpc.LinkServiceClient
}

func New(logger logger.Logger, linkService link_rpc.LinkServiceClient) (*Service, error) {
	service := &Service{
		LinkServiceClient: linkService,

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

	newCtx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()

	for key := range payload.URL {
		g.Go(func() error {
			_, errAddLink := s.LinkServiceClient.Add(newCtx, &link_rpc.AddRequest{Link: &link.Link{Url: payload.URL[key].Loc}})
			return errAddLink
		})
	}

	err = g.Wait()
	if err != nil {
		return err
	}

	return nil
}
