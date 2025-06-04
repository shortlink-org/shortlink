package sqlite

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
)

// Store implementation of db interface
type Store struct {
	client *sql.DB
}

// New store
func New(_ context.Context, store db.DB) (*Store, error) {
	conn, ok := store.GetConn().(*sql.DB)
	if !ok {
		return nil, db.ErrGetConnection
	}

	s := &Store{
		client: conn,
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
	links := squirrel.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := lite.client.Prepare(q)
	if err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}
	defer stmt.Close() //nolint:errcheck // ignore

	var (
		link     string
		hash     string
		describe string
	)

	err = stmt.QueryRowContext(ctx, args...).Scan(&link, &hash, &describe)
	if err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	response, err := v1.NewLinkBuilder().SetURL(link).Build()
	if err != nil {
		return nil, err
	}

	return response, nil
}

// List - list
func (lite *Store) List(ctx context.Context, _ *types.FilterLink) (*v1.Links, error) {
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

	response := v1.NewLinks()

	for rows.Next() {
		var (
			link string
			hash string
			desc string
		)

		err = rows.Scan(&link, &hash, &desc)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}}
		}

		result, err := v1.NewLinkBuilder().SetURL(link).SetDescribe(desc).Build()
		if err != nil {
			return nil, err
		}

		response.Push(result)
	}

	return response, nil
}

// Add - add
func (lite *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
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
	links := squirrel.Delete("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = lite.client.ExecContext(ctx, q, args...)
	if err != nil {
		return &v1.NotFoundByHashError{Hash: id}
	}

	return nil
}
