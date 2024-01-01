/*
Metadata Service. Application layer
*/

package screenshot

import (
	"context"
	"net/url"

	"github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/infrastructure/repository/media"
)

type UC struct {
	media media.Service
}

func New(ctx context.Context) (*UC, error) {
	return &UC{}, nil
}

func (s *UC) Get(ctx context.Context, linkURL string) (*url.URL, error) {
	return s.media.Get(ctx, linkURL)
}

func (s *UC) Put(ctx context.Context, linkURL string) error {
	screenshot := []byte("screenshot")

	return s.media.Put(ctx, linkURL, screenshot)
}
