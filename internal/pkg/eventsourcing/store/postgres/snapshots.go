package es_postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
)

// GetAggregateWithoutSnapshot - get aggregates without a snapshot
func (s *Store) GetAggregateWithoutSnapshot(ctx context.Context) ([]*eventsourcing.BaseAggregate, error) {
	query := psql.Select("aggregates.id", "aggregates.type", "aggregates.version").
		From("billing.aggregates AS aggregates").
		LeftJoin("billing.snapshots AS snapshots ON aggregates.id = snapshots.aggregate_id").
		Where("aggregates.version > snapshots.aggregate_version OR snapshots.aggregate_version IS NULL")

	q, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(ctx, q, args...)
	if err != nil || rows.Err() != nil {
		return nil, err
	}

	var aggregates []*eventsourcing.BaseAggregate // nolint:prealloc
	for rows.Next() {
		var (
			id            sql.NullString
			typeAggregate sql.NullString
			version       sql.NullInt32
		)
		err = rows.Scan(&id, &typeAggregate, &version)
		if err != nil {
			return nil, err
		}
		if rows.Err() != nil {
			return nil, rows.Err()
		}

		aggregates = append(aggregates, &eventsourcing.BaseAggregate{
			Id:      id.String,
			Type:    typeAggregate.String,
			Version: version.Int32,
		})
	}

	return aggregates, nil
}

func (s *Store) SaveSnapshot(ctx context.Context, snapshot *eventsourcing.Snapshot) error {
	// TODO: use worker pool

	// start tracing
	_, span := otel.Tracer("snapshot").Start(ctx, "SaveSnapshot")
	span.SetAttributes(attribute.String("aggregate id", snapshot.AggregateId))
	defer span.End()

	query := psql.Insert("billing.snapshots").
		Columns("aggregate_id", "aggregate_type", "aggregate_version", "payload").
		Values(snapshot.AggregateId, snapshot.AggregateType, snapshot.AggregateVersion, snapshot.Payload).
		Suffix("ON CONFLICT (aggregate_id) DO UPDATE SET aggregate_version = ?, payload = ?, updated_at = ?", snapshot.AggregateVersion, snapshot.Payload, time.Now())

	q, args, err := query.ToSql()
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
