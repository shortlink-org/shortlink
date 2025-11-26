package parsers

import (
	"context"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	http_client "github.com/shortlink-org/go-sdk/http/client"

	domainerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/errors"
	v1 "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
)

func (r *UC) Set(ctx context.Context, url string) (*v1.Meta, error) {
	meta := &v1.Meta{
		Id: url,
	}

	newCtx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()

	// Request the HTML page.
	req, err := http.NewRequestWithContext(newCtx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, domainerrors.NewInvalidURLError(url, err)
	}

	client, err := http_client.New()
	if err != nil {
		return nil, domainerrors.NewMetadataExtractionError(url, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, domainerrors.NewMetadataExtractionError(url, err)
	}
	defer resp.Body.Close() //nolint:errcheck // ignore close error

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, domainerrors.NewMetadataExtractionError(url, err)
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" {
			meta.Description, _ = s.Attr("content")
		}

		if name, _ := s.Attr("name"); name == "keywords" {
			meta.Keywords, _ = s.Attr("content")
		}
	})

	// Write to DB
	err = r.MetaStore.Store.Add(newCtx, meta)
	if err != nil {
		return nil, domainerrors.ProcessingFailed("metadata.parser.store.add", err)
	}

	return meta, nil
}
