package es_postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
)

type Events interface {
	addEvent(ctx context.Context, event *eventsourcing.Event) error
}

func (e *Store) addEvent(ctx context.Context, event *eventsourcing.Event) error {
	entities := psql.Insert("billing.events").
		Columns("aggregate_id", "aggregate_type", "type", "payload", "version").
		Values(event.AggregateId, event.AggregateType, event.Type, event.Payload, event.Version)

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
