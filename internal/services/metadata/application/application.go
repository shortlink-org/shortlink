/*
Metadata Service. Application layer
*/
package metadata

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/metadata/domain"
	meta_store "github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
)

type Service struct {
	MetaStore *meta_store.MetaStore
}

func New(store *meta_store.MetaStore) (*Service, error) {
	return &Service{
		MetaStore: store,
	}, nil
}

func (r *Service) Get(ctx context.Context, hash string) (*domain.Meta, error) {
	meta, err := r.MetaStore.Store.Get(ctx, hash)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (r *Service) Set(ctx context.Context, url string) (*domain.Meta, error) {
	meta := &domain.Meta{
		Id: url,
	}

	// Request the HTML page.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" {
			meta.Description, _ = s.Attr("content")
		}

		if name, _ := s.Attr("name"); name == "keyworlds" {
			meta.Keywords, _ = s.Attr("content")
		}
	})

	// Write to DB
	err = r.MetaStore.Store.Add(ctx, meta)
	if err != nil {
		return nil, err
	}

	// publish event by this service
	notify.Publish(ctx, domain.METHOD_ADD, meta, nil)

	return meta, nil
}
