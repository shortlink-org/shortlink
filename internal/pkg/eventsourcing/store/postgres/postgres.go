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

func (s *Store) Init(ctx context.Context, store db.DB) error {
	var ok bool

	s.db, ok = store.GetConn().(*pgxpool.Pool)
	if !ok {
		return db.ErrGetConnection
	}

	return nil
}

func (s *Store) save(ctx context.Context, events []*eventsourcing.Event, _ bool) error {
	if len(events) == 0 {
		return nil
	}

	for _, event := range events {
		// TODO: use transaction
		// Either insert a new aggregate or append to an existing.
		if event.GetVersion() == 1 { //nolint:nestif // TODO: refactor
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
				squirrel.Gt{"version": snapshot.GetAggregateVersion()},
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

	defer rows.Close()

	var events []*eventsourcing.Event //nolint:prealloc // false positive

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

	if rows.Err() != nil {
		return nil, nil, rows.Err()
	}

	return &snapshot, events, nil
}
