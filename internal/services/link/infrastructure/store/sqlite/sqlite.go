package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *sql.DB
}

// Init ...
func (s *Store) Init(_ context.Context, db *db.Store) error {
	s.client = db.Store.GetConn().(*sql.DB)

	return nil
}

// Get ...
func (lite *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := lite.client.Prepare(query)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	defer stmt.Close() // nolint errcheck

	var response link.Link
	err = stmt.QueryRowContext(ctx, args...).Scan(&response.Url, &response.Hash, &response.Describe)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (lite *Store) List(ctx context.Context, _ *query.Filter) (*link.Links, error) {
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links")
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := lite.client.QueryContext(ctx, query, args...)
	if err != nil || rows.Err() != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}
	defer rows.Close() // nolint errcheck

	response := &link.Links{
		Link: []*link.Link{},
	}

	for rows.Next() {
		var result link.Link
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		response.Link = append(response.Link, &result)
	}

	return response, nil
}

// Add ...
func (lite *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	err := link.NewURL(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := squirrel.Insert("links").
		Columns("url", "hash", "describe").
		Values(source.Url, source.Hash, source.Describe)

	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = lite.client.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

// Update ...
func (lite *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (lite *Store) Delete(ctx context.Context, id string) error {
	// query builder
	links := squirrel.Delete("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = lite.client.ExecContext(ctx, query, args...)
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Hash: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}
