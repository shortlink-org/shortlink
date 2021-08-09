package postgres

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	"github.com/batazor/shortlink/internal/services/metadata/domain"
)

// LinkUpdate ...
func (s *Store) MetadataUpdate(ctx context.Context, in *domain.Meta) (*domain.Meta, error) {
	// query builder
	metadata := psql.Update("shortlink.link_view").
		Set("image_url", in.ImageURL).
		Set("meta_description", in.Description).
		Set("meta_keywords", in.Keywords).
		Where(squirrel.Eq{"url": in.Id})

	q, args, err := metadata.ToSql()
	if err != nil {
		return nil, err
	}

	row := s.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&in.ImageURL, &in.Description, &in.Keywords)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return in, nil
	}
	if errScan.Error() != "" {
		return nil, errScan
	}

	return in, nil
}
