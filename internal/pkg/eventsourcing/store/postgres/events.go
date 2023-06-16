package es_postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
)

type Events interface {
	addEvent(ctx context.Context, event *eventsourcing.Event) error
}

func (s *Store) addEvent(ctx context.Context, event *eventsourcing.Event) error {
	// start tracing
	_, span := otel.Tracer("aggregate").Start(ctx, "addEvent")
	span.SetAttributes(attribute.String("event_id", event.Id))
	defer span.End()

	entities := psql.Insert("billing.events").
		Columns("aggregate_id", "aggregate_type", "type", "payload", "version").
		Values(event.AggregateId, event.AggregateType, event.Type, event.Payload, event.Version)

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
