package es_postgres

import (
	"context"
	"embed"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"

	"github.com/shortlink-org/go-sdk/db"
	"github.com/shortlink-org/go-sdk/db/drivers/postgres/migrate"
	eventsourcing "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing/domain/eventsourcing/v1"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS

	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func New(ctx context.Context, store db.DB) (*EventStore, error) {
	conn, ok := store.GetConn().(*pgxpool.Pool)
	if !ok {
		return nil, db.ErrGetConnection
	}

	// Migration ---------------------------------------------------------------------------------------------------
	err := migrate.Migration(ctx, store, migrations, "eventsourcing")
	if err != nil {
		return nil, err
	}

	return &EventStore{
		db: conn,
	}, nil
}

func (e *EventStore) save(ctx context.Context, events []*eventsourcing.Event, _ bool) error {
	if len(events) == 0 {
		return nil
	}

	for _, event := range events {
		// TODO: use transaction
		// Either insert a new aggregate or append to an existing.
		if event.GetVersion() == 1 { //nolint:nestif // TODO: refactor
			err := e.addAggregate(ctx, event)
			if err != nil {
				return err
			}

			err = e.addEvent(ctx, event)
			if err != nil {
				return err
			}
		} else {
			err := e.updateAggregate(ctx, event)
			if err != nil {
				return err
			}

			err = e.addEvent(ctx, event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *EventStore) Save(ctx context.Context, events []*eventsourcing.Event) error {
	// start tracing
	newCtx, span := otel.Tracer("store").Start(ctx, "Save")
	defer span.End()

	return e.save(newCtx, events, false)
}

func (e *EventStore) SafeSave(ctx context.Context, events []*eventsourcing.Event) error {
	// start tracing
	newCtx, span := otel.Tracer("store").Start(ctx, "SafeSave")
	defer span.End()

	return e.save(newCtx, events, true)
}

func (e *EventStore) Load(ctx context.Context, aggregateID string) (*eventsourcing.Snapshot, []*eventsourcing.Event, error) {
	// start tracing
	_, span := otel.Tracer("store").Start(ctx, "Load")
	defer span.End()

	// get snapshot if exist
	querySnaphot := psql.Select("aggregate_id", "aggregate_type", "aggregate_version", "payload").
		From("snapshots").
		Where(squirrel.Eq{
			"aggregate_id": aggregateID,
		})
	q, args, err := querySnaphot.ToSql()
	if err != nil {
		return nil, nil, err
	}

	var snapshot eventsourcing.Snapshot

	row := e.db.QueryRow(ctx, q, args...)

	err = row.Scan(&snapshot.AggregateId, &snapshot.AggregateType, &snapshot.AggregateVersion, &snapshot.Payload)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, nil, err
	}

	// get new events
	queryEvents := psql.Select("aggregate_type", "id", "type", "payload", "version").
		From("events").
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

	rows, err := e.db.Query(ctx, q, args...)
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
