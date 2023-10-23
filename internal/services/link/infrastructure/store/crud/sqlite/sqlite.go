package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

// Store implementation of db interface
type Store struct {
	client *sql.DB
}

// New store
func New(ctx context.Context, db *db.Store) (*Store, error) {
	s := &Store{
		client: db.Store.GetConn().(*sql.DB),
	}

	// Migration
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS links (
			id integer not null primary key,
			url      varchar(255) not null,
			hash     varchar(255) not null,
			describe text
		);
	`

	if _, err := s.client.Exec(sqlStmt); err != nil {
		return nil, err
	}

	return s, nil
}

// Get ...
func (lite *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
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
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	defer stmt.Close() //nolint:errcheck

	var response v1.Link
	err = stmt.QueryRowContext(ctx, args...).Scan(&response.Url, &response.Hash, &response.Describe)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (lite *Store) List(ctx context.Context, _ *query.Filter) (*v1.Links, error) {
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links")
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := lite.client.QueryContext(ctx, q, args...)
	if err != nil || rows.Err() != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: query.ErrNotFound}
	}
	defer rows.Close() //nolint:errcheck

	response := &v1.Links{
		Link: []*v1.Link{},
	}

	for rows.Next() {
		var result v1.Link
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: query.ErrNotFound}
		}

		response.Link = append(response.GetLink(), &result)
	}

	return response, nil
}

// Add ...
func (lite *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := squirrel.Insert("links").
		Columns("url", "hash", "describe").
		Values(source.GetUrl(), source.GetHash(), source.GetDescribe())

	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = lite.client.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.GetUrl())}
	}

	return source, nil
}

// Update ...
func (lite *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
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
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}
