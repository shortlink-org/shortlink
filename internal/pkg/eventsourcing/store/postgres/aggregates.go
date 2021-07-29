package es_postgres

import (
	"context"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
)

type Aggregates interface {
	addAggregate(ctx context.Context, event *eventsourcing.Event) error
	updateAggregate(ctx context.Context, event *eventsourcing.Event) error
}

func (e *Store) addAggregate(ctx context.Context, event *eventsourcing.Event) error {
	entities := psql.Insert("billing.aggregates").
		Columns("id", "type", "version").
		Values(event.AggregateId, event.AggregateType, event.Version)

	q, args, err := entities.ToSql()
	if err != nil {
		return err
	}

	row := e.db.QueryRow(ctx, q, args...)
	errScan := row.Scan()
	if errors.Is(errScan, pgx.ErrNoRows) {
		return nil
	}
	if errScan.Error() != "" {
		return errScan
	}

	return nil
}

func (e *Store) updateAggregate(ctx context.Context, event *eventsourcing.Event) error {
	entities := psql.Update("billing.aggregates").
		Set("version", event.Version).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{
			"version": event.Version,
			"id":      event.AggregateId,
		})

	q, args, err := entities.ToSql()
	if err != nil {
		return err
	}

	row := e.db.QueryRow(ctx, q, args...)
	errScan := row.Scan()
	if errors.Is(errScan, pgx.ErrNoRows) {
		return nil
	}
	if errScan.Error() != "" {
		return errScan
	}

	return nil
}
