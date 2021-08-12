package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
)

// LinkAdd ...
func (s *Store) LinkAdd(ctx context.Context, source *v1.Link) (*v1.Link, error) {
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

	row := s.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&source.Url, &source.Hash, &source.Describe)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return source, nil
	}
	if errScan.Error() != "" {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

// LinkUpdate ...
func (s *Store) LinkUpdate(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	// query builder
	links := psql.Update("shortlink.link_view").
		Set("url", source.Url).
		Set("hash", source.Hash).
		Set("describe", source.Describe).
		Where(squirrel.Eq{"hash": source.Hash})

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := s.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&source.Url, &source.Hash, &source.Describe)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return source, nil
	}
	if errScan.Error() != "" {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

// LinkDelete ...
func (s *Store) LinkDelete(ctx context.Context, id string) error {
	// query builder
	request := psql.Delete("shortlink.link_view").
		Where(squirrel.Eq{"hash": id})
	q, args, err := request.ToSql()
	if err != nil {
		return err
	}

	_, err = s.client.Exec(ctx, q, args...)
	if err != nil {
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}
