package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

// Init ...
func (s *Store) Init(ctx context.Context, db *db.Store) error {
	// Set configuration
	s.client = db.Store.GetConn().(*pgxpool.Pool)

	return nil
}

// Get ...
func (s *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("shortlink.link_view").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)

	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	if rows.Err() != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response v1.Link
	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
		}
	}

	if response.Hash == "" {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (s *Store) List(ctx context.Context, filter *query.Filter) (*v1.Links, error) {
	// query builder
	links := psql.Select("url, hash, describe, created_at, updated_at").
		From("shortlink.link_view")
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
	}

	response := &v1.Links{
		Link: []*v1.Link{},
	}

	for rows.Next() {
		var result v1.Link
		var (
			created_ad sql.NullTime
			updated_at sql.NullTime
		)
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe, &created_ad, &updated_at)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
		}
		result.CreatedAt = &timestamp.Timestamp{Seconds: int64(created_ad.Time.Second()), Nanos: int32(created_ad.Time.Nanosecond())}
		result.UpdatedAt = &timestamp.Timestamp{Seconds: int64(updated_at.Time.Second()), Nanos: int32(updated_at.Time.Nanosecond())}

		response.Link = append(response.Link, &result)
	}

	return response, nil
}
