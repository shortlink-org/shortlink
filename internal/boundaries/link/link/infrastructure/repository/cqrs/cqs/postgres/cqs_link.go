package postgres

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

func (s *Store) LinkAdd(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := psql.Insert("link.link_view").
		Columns("url", "hash", "describe").
		Values(source.GetUrl(), source.GetHash(), source.GetDescribe())

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
		return nil, &v1.NotFoundError{Link: source}
	}

	return source, nil
}

// LinkUpdate - update link
func (s *Store) LinkUpdate(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	// query builder
	links := psql.Update("link.link_view").
		Set("url", source.GetUrl()).
		Set("hash", source.GetHash()).
		Set("describe", source.GetDescribe()).
		Where(squirrel.Eq{"hash": source.GetHash()})

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
		return nil, &v1.NotFoundError{Link: source}
	}

	return source, nil
}

// LinkDelete - delete link
func (s *Store) LinkDelete(ctx context.Context, id string) error {
	// query builder
	request := psql.Delete("link.link_view").
		Where(squirrel.Eq{"hash": id})
	q, args, err := request.ToSql()
	if err != nil {
		return err
	}

	_, err = s.client.Exec(ctx, q, args...)
	if err != nil {
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}}
	}

	return nil
}
