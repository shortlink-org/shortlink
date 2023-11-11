package sqlite

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// Store implementation of db interface
type Store struct {
	client *sql.DB
}

// New store
func New(_ context.Context, store db.DB) (*Store, error) {
	s := &Store{
		client: store.GetConn().(*sql.DB),
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

// Get - get
func (lite *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := lite.client.Prepare(q)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}}
	}
	defer stmt.Close() //nolint:errcheck // ignore

	var response v1.Link
	err = stmt.QueryRowContext(ctx, args...).Scan(&response.Url, &response.Hash, &response.Describe)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}}
	}

	return &response, nil
}

// List - list
func (lite *Store) List(ctx context.Context, _ *v1.FilterLink) (*v1.Links, error) {
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links")
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := lite.client.QueryContext(ctx, q, args...)
	if err != nil || rows.Err() != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}
	defer func() {
		_ = rows.Close()
	}()

	response := &v1.Links{
		Link: []*v1.Link{},
	}

	for rows.Next() {
		var result v1.Link
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}}
		}

		response.Link = append(response.GetLink(), &result)
	}

	return response, nil
}

// Add - add
func (lite *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := squirrel.Insert("links").
		Columns("url", "hash", "describe").
		Values(source.GetUrl(), source.GetHash(), source.GetDescribe())

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = lite.client.ExecContext(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Link: source}
	}

	return source, nil
}

// Update - update
func (lite *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete - delete
func (lite *Store) Delete(ctx context.Context, id string) error {
	// query builder
	links := squirrel.Delete("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = lite.client.ExecContext(ctx, q, args...)
	if err != nil {
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}}
	}

	return nil
}
