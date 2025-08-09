package postgres

import (
	"context"

	"github.com/lib/pq"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres/filter"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
)

// List - return list links
func (s *Store) List(ctx context.Context, params *v1.FilterLink) (*domain.Links, error) {
	request := psql.Select("url", "hash", "describe", "created_at", "updated_at").
		From("link.links")

	// Build filter
	request = filter.NewFilter(params).BuildFilter(request)

	q, args, err := request.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil || rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	links := domain.NewLinks()
	for rows.Next() {
		var (
			url       string
			hash      string
			describe  string
			createdAt pq.NullTime
			updatedAt pq.NullTime
		)

		err = rows.Scan(&url, &hash, &describe, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		link, errBuilder := domain.NewLinkBuilder().
			SetURL(url).
			SetDescribe(describe).
			SetCreatedAt(createdAt.Time).
			SetUpdatedAt(updatedAt.Time).
			Build()

		if errBuilder != nil {
			return nil, errBuilder
		}

		links.Push(link)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return links, nil
}
