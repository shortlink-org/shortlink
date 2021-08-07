package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
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

// Add ...
func (p *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := psql.Insert("shortlink.link_view").
		Columns("url", "hash", "describe").
		Values(source.Url, source.Hash, source.Describe)

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&source.Url, &source.Hash, &source.Describe)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return source, nil
	}
	if errScan.Error() != "" {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

// Update ...
func (p *Store) Update(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	// query builder
	links := psql.Update("shortlink.link_view").
		Set("url", source.Url).
		Set("hash", source.Hash).
		Set("describe", source.Describe)

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&source.Url, &source.Hash, &source.Describe)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return source, nil
	}
	if errScan.Error() != "" {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

// Delete ...
func (p *Store) Delete(ctx context.Context, id string) error {
	// query builder
	request := psql.Delete("shortlink.link_view").
		Where(squirrel.Eq{"hash": id})
	q, args, err := request.ToSql()
	if err != nil {
		return err
	}

	_, err = p.client.Exec(ctx, q, args...)
	if err != nil {
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}
