package es_postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
)

type Aggregates interface {
	addAggregate(ctx context.Context, event *eventsourcing.Event) error
	updateAggregate(ctx context.Context, event *eventsourcing.Event) error
}

func (s *Store) addAggregate(ctx context.Context, event *eventsourcing.Event) error {
	// start tracing
	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("addAggregate"))
	defer span.Finish()

	entities := psql.Insert("billing.aggregates").
		Columns("id", "type", "version").
		Values(event.AggregateId, event.AggregateType, event.Version)

	q, args, err := entities.ToSql()
	if err != nil {
		return err
	}

	row := s.db.QueryRow(ctx, q, args...)
	errScan := row.Scan()
	if errors.Is(errScan, pgx.ErrNoRows) {
		return nil
	}
	if errScan.Error() != "" {
		span.SetTag("error", true)
		span.SetTag("message", errScan.Error())
		return errScan
	}

	return nil
}

func (s *Store) updateAggregate(ctx context.Context, event *eventsourcing.Event) error {
	// start tracing
	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("updateAggregate"))
	defer span.Finish()

	entities := psql.Update("billing.aggregates").
		Set("version", event.Version).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{
			"version": event.Version - 1,
			"id":      event.AggregateId,
		})

	q, args, err := entities.ToSql()
	if err != nil {
		span.SetTag("error", true)
		span.SetTag("message", err.Error())
		return err
	}

	row, err := s.db.Exec(ctx, q, args...)
	if err != nil {
		span.SetTag("error", true)
		span.SetTag("message", err.Error())
		return err
	}

	if row.RowsAffected() != 1 {
		return fmt.Errorf(`Incorrect updated billing.aggregates. Updated: %d/1`, row.RowsAffected())
	}

	return nil
}
