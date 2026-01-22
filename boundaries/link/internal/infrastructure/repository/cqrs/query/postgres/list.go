package postgres

import (
	"context"
	"database/sql"
	"fmt"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	v13 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	v12 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link_cqrs/v1"
)

// List - list
func (s *Store) List(ctx context.Context, filter *v13.FilterLink) (*v12.LinksView, error) {
	queryTerm := ""
	if filter != nil && filter.URL != nil && len(filter.URL.Contains) > 0 {
		queryTerm = filter.URL.Contains[0]
	}
	if queryTerm == "" {
		return nil, &v1.NotFoundError{Hash: ""}
	}

	links := psql.Select(
		"url",
		"hash",
		"describe",
		"image_url",
		"ts_headline(meta_description, q, 'StartSel=<em>, StopSel=</em>') as meta_description",
		"meta_keywords",
		"created_at",
		"updated_at",
	).
		From(fmt.Sprintf(`link.link_view, to_tsquery('%s') AS q`, queryTerm)).
		Where("make_tsvector_link_view(meta_keywords, meta_description) @@ q").
		OrderBy("ts_rank(make_tsvector_link_view(meta_keywords, meta_description), q) DESC")

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Hash: ""}
	}
	defer rows.Close()

	response := v12.NewLinksView()

	for rows.Next() {
		var (
			urlValue        string
			hashValue       string
			describeValue   string
			imageURLValue   sql.NullString
			metaDescValue   sql.NullString
			metaKeywordsVal sql.NullString
			createdAt       sql.NullTime
			updatedAt       sql.NullTime
		)

		if err := rows.Scan(
			&urlValue,
			&hashValue,
			&describeValue,
			&imageURLValue,
			&metaDescValue,
			&metaKeywordsVal,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, &v1.NotFoundError{Hash: ""}
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
		if createdAt.Valid {
			builder = builder.SetCreatedAt(createdAt.Time)
		}
		if updatedAt.Valid {
			builder = builder.SetUpdatedAt(updatedAt.Time)
		}

		result, err := builder.Build()
		if err != nil {
			return nil, &v1.NotFoundError{Hash: ""}
		}

		response.AddLink(result)
	}

	if rows.Err() != nil {
		return nil, &v1.NotFoundError{Hash: ""}
	}

	return response, nil
}
