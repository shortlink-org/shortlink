package postgres

import (
	"context"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// Get - a get link
func (s *Store) Get(ctx context.Context, hash string) (*domain.Link, error) {
	link, err := s.query.GetLinkByHash(ctx, hash)
	if err != nil {
		return nil, &domain.NotFoundError{Hash: hash}
	}

	resp, err := domain.NewLinkBuilder().
		SetURL(link.Url).
		SetDescribe(link.Describe).
		Build()

	return resp, nil
}
