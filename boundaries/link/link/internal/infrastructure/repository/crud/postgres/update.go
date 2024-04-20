package postgres

import (
	"context"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres/schema/crud"
)

// Update - update link
func (s *Store) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	link := in.GetUrl()
	_, err := s.query.UpdateLink(ctx, crud.UpdateLinkParams{
		Url:      link.String(),
		Hash:     in.GetHash(),
		Describe: in.GetDescribe(),
		Json:     *in,
	})
	if err != nil {
		return nil, err
	}

	return in, nil
}
