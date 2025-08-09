package postgres

import (
	"context"

	"github.com/segmentio/encoding/json"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	repository_err "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/error"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres/schema/crud"
)

// Update - update link
func (s *Store) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	payload, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	row, err := s.query.UpdateLink(ctx, crud.UpdateLinkParams{
		Url:      in.GetUrl().String(),
		Hash:     in.GetHash(),
		Describe: in.GetDescribe(),
		Column4:  payload,
	})
	if err != nil {
		return nil, err
	}

	if row.RowsAffected() == 0 {
		return nil, repository_err.ErrNotFound
	}

	return in, nil
}
