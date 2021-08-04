package es_postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
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

	var aggregates []*eventsourcing.BaseAggregate

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
	// start tracing
	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("SaveSnapshot"))
	span.SetTag("snapshot aggregate id", snapshot.AggregateId)
	defer span.Finish()

	query := psql.Insert("billing.snapshots").
		Columns("aggregate_id", "aggregate_type", "aggregate_version", "payload").
		Values(snapshot.AggregateId, snapshot.AggregateType, snapshot.AggregateVersion, snapshot.Payload).
		Suffix("ON CONFLICT (aggregate_id) DO UPDATE SET aggregate_version = ?, payload = ?, updated_at = ?", snapshot.AggregateVersion, snapshot.Payload, time.Now())

	q, args, err := query.ToSql()
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
