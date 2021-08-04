package es_postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
)

type Events interface {
	addEvent(ctx context.Context, event *eventsourcing.Event) error
}

func (s *Store) addEvent(ctx context.Context, event *eventsourcing.Event) error {
	// start tracing
	span, _ := opentracing.StartSpanFromContext(ctx, "addEvent")
	span.SetTag("event id", event.Id)
	defer span.Finish()

	entities := psql.Insert("billing.events").
		Columns("aggregate_id", "aggregate_type", "type", "payload", "version").
		Values(event.AggregateId, event.AggregateType, event.Type, event.Payload, event.Version)

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
