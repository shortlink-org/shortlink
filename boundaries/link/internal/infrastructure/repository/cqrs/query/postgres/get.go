package postgres

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	v12 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link_cqrs/v1"
)

// Get - get
func (s *Store) Get(ctx context.Context, id string) (*v12.LinkView, error) {
	links := psql.Select("url, hash, describe", "image_url", "meta_description", "meta_keywords").
		From("link.link_view").
		Where(squirrel.Eq{"hash": id})

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Hash: id}
	}
	defer rows.Close()

	if !rows.Next() {
		if rows.Err() != nil {
			return nil, &v1.NotFoundError{Hash: id}
		}

		return nil, &v1.NotFoundError{Hash: id}
	}

	var (
		urlValue        string
		hashValue       string
		describeValue   string
		imageURLValue   sql.NullString
		metaDescValue   sql.NullString
		metaKeywordsVal sql.NullString
	)

	if err := rows.Scan(&urlValue, &hashValue, &describeValue, &imageURLValue, &metaDescValue, &metaKeywordsVal); err != nil {
		return nil, &v1.NotFoundError{Hash: id}
	}

	builder := v12.NewLinkViewBuilder().
		SetURL(urlValue).
		SetHash(hashValue).
		SetDescribe(describeValue)

	if imageURLValue.Valid {
		builder = builder.SetImageUrl(imageURLValue.String)
	}
	if metaDescValue.Valid {
		builder = builder.SetMetaDescription(metaDescValue.String)
	}
	if metaKeywordsVal.Valid {
		builder = builder.SetMetaKeywords(metaKeywordsVal.String)
	}

	response, err := builder.Build()
	if err != nil {
		return nil, &v1.NotFoundError{Hash: id}
	}

	return response, nil
}
