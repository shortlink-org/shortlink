package es_postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
)

type Aggregates interface {
	addAggregate(ctx context.Context, event *eventsourcing.Event) error
	updateAggregate(ctx context.Context, event *eventsourcing.Event) error
}

func (s *Store) addAggregate(ctx context.Context, event *eventsourcing.Event) error {
	// start tracing
	_, span := otel.Tracer("aggregate").Start(ctx, "addAggregate")
	defer span.End()

	entities := psql.Insert("billing.aggregates").
		Columns("id", "type", "version").
		Values(event.AggregateId, event.AggregateType, event.Version)

	q, args, err := entities.ToSql()
	if err != nil {
		return err
	}

	row := s.db.QueryRow(ctx, q, args...)
	err = row.Scan()
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err.Error() != "" {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

func (s *Store) updateAggregate(ctx context.Context, event *eventsourcing.Event) error {
	// start tracing
	_, span := otel.Tracer("aggregate").Start(ctx, "updateAggregate")
	defer span.End()

	entities := psql.Update("billing.aggregates").
		Set("version", event.Version).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{
			"version": event.Version - 1,
			"id":      event.AggregateId,
		})

	q, args, err := entities.ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	row, err := s.db.Exec(ctx, q, args...)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	if row.RowsAffected() != 1 {
		return fmt.Errorf(`Incorrect updated billing.aggregates. Updated: %d/1`, row.RowsAffected())
	}

	return nil
}
