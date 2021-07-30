package es_postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
)

type Store struct {
	db *pgxpool.Pool

	Aggregates
	Events
}

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

func (s *Store) Init(ctx context.Context, db *db.Store) error {
	s.db = db.Store.GetConn().(*pgxpool.Pool)
	return nil
}

func (s *Store) save(ctx context.Context, events []*eventsourcing.Event, safe bool) error { // nolint govet
	if len(events) == 0 {
		return nil
	}

	for _, event := range events {
		// Either insert a new aggregate or append to an existing.
		if event.Version == 1 {
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
	return s.save(ctx, events, false)
}

func (s *Store) SafeSave(ctx context.Context, events []*eventsourcing.Event) error {
	return s.save(ctx, events, true)
}

func (s *Store) Load(ctx context.Context, aggregateID string) ([]*eventsourcing.Event, error) {
	var events []*eventsourcing.Event

	query := psql.Select("aggregate_type", "id", "type", "payload", "version").
		From("billing.events").
		Where(squirrel.Eq{
			"aggregate_id": aggregateID,
		}).
		OrderBy("created_at")

	q, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for rows.Next() {
		event := eventsourcing.Event{
			AggregateId: aggregateID,
		}
		err = rows.Scan(&event.AggregateType, &event.Id, &event.Type, &event.Payload, &event.Version)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}
