package postgres

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shortlink-org/go-sdk/db/options"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/postgres/schema/crud"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/types/v1"
)

// mapPostgresError maps PostgreSQL errors to domain errors
func mapPostgresError(err error, context string) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// PostgreSQL unique violation error code
		// https://www.postgresql.org/docs/current/errcodes-appendix.html
		if pgErr.Code == "23505" { // unique_violation
			return domain.NewConflictError(fmt.Sprintf("%s: duplicate entry (constraint: %s)", context, pgErr.ConstraintName))
		}
	}

	// Return internal error if not a known PostgreSQL error
	return domain.NewInternalErrorWithErr(err)
}

// Add - an add link
func (s *Store) Add(ctx context.Context, source *domain.Link) (*domain.Link, error) {
	switch s.config.mode {
	case options.MODE_BATCH_WRITE:
		resCh := s.config.job.Push(source)

		select {
		case res, ok := <-resCh:
			if !ok || res == nil {
				return nil, domain.NewInternalError("batch write failed")
			}
			return res, nil
		case <-ctx.Done():
			return nil, domain.NewInternalErrorWithErr(ctx.Err())
		}
	case options.MODE_SINGLE_WRITE:
		data, err := s.singleWrite(ctx, source)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, nil
}

func (s *Store) singleWrite(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	// Create DTO with protobuf timestamps for proper JSON serialization
	dto := &v1.Link{
		Url:       in.GetUrl().String(),
		Hash:      in.GetHash(),
		Describe:  in.GetDescribe(),
		CreatedAt: in.GetCreatedAt().GetTimestamp(),
		UpdatedAt: in.GetUpdatedAt().GetTimestamp(),
	}

	// Use protojson.Marshal for proper protobuf serialization (timestamps, etc.)
	payload, err := protojson.Marshal(dto)
	if err != nil {
		return nil, domain.NewInternalErrorWithErr(err)
	}

	links := psql.Insert("link.links").
		Columns("url", "hash", "describe", "json").
		Values(in.GetUrl().String(), in.GetHash(), in.GetDescribe(), string(payload))

	q, args, err := links.ToSql()
	if err != nil {
		return nil, domain.NewInternalErrorWithErr(err)
	}

	_, err = s.client.Exec(ctx, q, args...)
	if err != nil {
		// Map PostgreSQL errors (e.g., unique violation) to domain errors
		return nil, mapPostgresError(err, fmt.Sprintf("create failed: url=%s", in.GetUrl().String()))
	}

	return in, nil
}

func (s *Store) batchWrite(ctx context.Context, in *domain.Links) (*domain.Links, error) {
	links := make([]crud.CreateLinksParams, 0, len(in.GetLinks()))

	// Create a new link
	list := in.GetLinks()
	for key := range list {
		dto := &v1.Link{
			Url:       list[key].GetUrl().String(),
			Hash:      list[key].GetHash(),
			Describe:  list[key].GetDescribe(),
			CreatedAt: list[key].GetCreatedAt().GetTimestamp(),
			UpdatedAt: list[key].GetUpdatedAt().GetTimestamp(),
		}

		// Marshal to JSONB as string (PostgreSQL JSONB requires string, not []byte)
		dataJson, err := protojson.Marshal(dto)
		if err != nil {
			return nil, domain.NewInternalErrorWithErr(err)
		}

		links = append(links, crud.CreateLinksParams{
			Url:      list[key].GetUrl().String(),
			Hash:     list[key].GetHash(),
			Describe: list[key].GetDescribe(),
			Json:     string(dataJson),
		})
	}

	_, err := s.query.CreateLinks(ctx, links)
	if err != nil {
		// Map PostgreSQL errors (e.g., unique violation) to domain errors
		mappedErr := mapPostgresError(err, "batch create failed")

		// If it's a conflict error, create individual errors for each link
		var conflictErr *domain.ConflictError
		if errors.As(mappedErr, &conflictErr) {
			errs := make([]error, 0, len(list))
			for key := range list {
				errs = append(errs, domain.NewConflictError(fmt.Sprintf("duplicate link: hash=%s", list[key].GetHash())))
			}
			return nil, errors.Join(errs...)
		}

		return nil, mappedErr
	}

	return in, nil
}
