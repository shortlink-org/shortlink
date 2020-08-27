package application

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	rpc "github.com/batazor/shortlink/internal/metadata/domain"
	"github.com/batazor/shortlink/internal/store"
)

type Repository struct {
	DB store.DB
}

func (r *Repository) Get(ctx context.Context, hash string) (*rpc.Meta, error) {
	// TODO: Correct read store
	link, err := r.DB.Get(ctx, hash)
	if err != nil {
		return nil, err
	}

	return &rpc.Meta{
		Id:          hash,
		Description: link.Describe,
	}, nil
}

func (r *Repository) Set(ctx context.Context, url string) (*rpc.Meta, error) {
	meta := &rpc.Meta{
		Id: url,
	}

	// Request the HTML page.
	resp, err := http.Get(url) // nolint gosec
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

	// TODO: Write to DB
	//_, err = r.DB.Add(ctx, &link.Link{
	//	Url:       url,
	//})
	//if err != nil {
	//	return nil, err
	//}

	return meta, nil
}
