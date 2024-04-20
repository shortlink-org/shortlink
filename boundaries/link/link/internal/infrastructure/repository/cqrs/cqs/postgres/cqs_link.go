package postgres

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
)

func (s *Store) LinkAdd(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	links := psql.Insert("link.link_view").
		Columns("url", "hash", "describe").
		Values(source.GetUrl(), source.GetHash(), source.GetDescribe())

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := s.client.QueryRow(ctx, q, args...)

	var (
		link     string
		hash     string
		describe string
	)
	errScan := row.Scan(&link, &hash, &describe)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return source, nil
	}

	if errScan.Error() != "" {
		return nil, &v1.NotFoundError{Link: source}
	}

	resp, err := v1.NewLinkBuilder().
		SetURL(link).
		SetDescribe(describe).
		Build()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// LinkUpdate - update link
func (s *Store) LinkUpdate(ctx context.Context, source *v1.Link) (*v1.Link, error) {
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

	var (
		link     string
		hash     string
		describe string
	)
	errScan := row.Scan(&link, &hash, &describe)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return source, nil
	}
	if errScan.Error() != "" {
		return nil, &v1.NotFoundError{Link: source}
	}

	resp, err := v1.NewLinkBuilder().
		SetURL(link).
		SetDescribe(describe).
		Build()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// LinkDelete - delete link
func (s *Store) LinkDelete(ctx context.Context, id string) error {
	request := psql.Delete("link.link_view").
		Where(squirrel.Eq{"hash": id})
	q, args, err := request.ToSql()
	if err != nil {
		return err
	}

	_, err = s.client.Exec(ctx, q, args...)
	if err != nil {
		return &v1.NotFoundByHashError{Hash: id}
	}

	return nil
}
