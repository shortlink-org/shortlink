package es_postgres

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
)

type Store struct {
	db *pgxpool.Pool

	Aggregates
	Events
}

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func (s *Store) Init(ctx context.Context, db *db.Store) error {
	var ok bool

	s.db, ok = db.Store.GetConn().(*pgxpool.Pool)
	if !ok {
		return errors.New("Error get connection to PostgreSQL")
	}

	return nil
}

func (s *Store) save(ctx context.Context, events []*eventsourcing.Event, safe bool) error { // nolint:govet
	if len(events) == 0 {
		return nil
	}

	for _, event := range events {
		// TODO: use transaction
		// Either insert a new aggregate or append to an existing.
		if event.Version == 1 { // nolint:nestif
			err := s.addAggregate(ctx, event)
			if err != nil {
				return err
			}

			err = s.addEvent(ctx, event)
			if err != nil {
				return err
			}
		} else {
			err := s.updateAggregate(ctx, event)
			if err != nil {
				return err
			}

			err = s.addEvent(ctx, event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Store) Save(ctx context.Context, events []*eventsourcing.Event) error {
	// start tracing
	newCtx, span := otel.Tracer("store").Start(ctx, "Save")
	defer span.End()

	return s.save(newCtx, events, false)
}

func (s *Store) SafeSave(ctx context.Context, events []*eventsourcing.Event) error {
	// start tracing
	newCtx, span := otel.Tracer("store").Start(ctx, "SafeSave")
	defer span.End()

	return s.save(newCtx, events, true)
}

func (s *Store) Load(ctx context.Context, aggregateID string) (*eventsourcing.Snapshot, []*eventsourcing.Event, error) {
	// start tracing
	_, span := otel.Tracer("store").Start(ctx, "Load")
	defer span.End()

	// get snapshot if exist
	querySnaphot := psql.Select("aggregate_id", "aggregate_type", "aggregate_version", "payload").
		From("billing.snapshots").
		Where(squirrel.Eq{
			"aggregate_id": aggregateID,
		})
	q, args, err := querySnaphot.ToSql()
	if err != nil {
		return nil, nil, err
	}

	var snapshot eventsourcing.Snapshot
	row := s.db.QueryRow(ctx, q, args...)
	err = row.Scan(&snapshot.AggregateId, &snapshot.AggregateType, &snapshot.AggregateVersion, &snapshot.Payload)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, nil, err
	}

	// get new events
	queryEvents := psql.Select("aggregate_type", "id", "type", "payload", "version").
		From("billing.events").
		Where(
			squirrel.And{
				squirrel.Eq{"aggregate_id": aggregateID},
				squirrel.Gt{"version": snapshot.AggregateVersion},
			},
		).
		OrderBy("created_at")

	q, args, err = queryEvents.ToSql()
	if err != nil {
		return nil, nil, err
	}

	rows, err := s.db.Query(ctx, q, args...)
	if err != nil {
		return nil, nil, err
	}
	if rows.Err() != nil {
		return nil, nil, rows.Err()
	}

	var events []*eventsourcing.Event // nolint:prealloc
	for rows.Next() {
		event := eventsourcing.Event{
			AggregateId: aggregateID,
		}
		err = rows.Scan(&event.AggregateType, &event.Id, &event.Type, &event.Payload, &event.Version)
		if err != nil {
			return nil, nil, err
		}

		events = append(events, &event)
	}

	return &snapshot, events, nil
}
