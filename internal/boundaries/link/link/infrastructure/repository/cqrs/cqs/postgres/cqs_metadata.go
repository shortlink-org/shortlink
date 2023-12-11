package postgres

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	v1 "github.com/shortlink-org/shortlink/internal/services/metadata/domain/metadata/v1"
)

func (s *Store) MetadataUpdate(ctx context.Context, in *v1.Meta) (*v1.Meta, error) {
	// query builder
	metadata := psql.Update("link.link_view").
		Set("image_url", in.GetImageUrl()).
		Set("meta_description", in.GetDescription()).
		Set("meta_keywords", in.GetKeywords()).
		Where(squirrel.Eq{"url": in.GetId()})

	q, args, err := metadata.ToSql()
	if err != nil {
		return nil, err
	}

	row := s.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&in.ImageUrl, &in.Description, &in.Keywords)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return in, nil
	}
	if errScan.Error() != "" {
		return nil, errScan
	}

	return in, nil
}
