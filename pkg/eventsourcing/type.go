package eventsourcing

import (
	"context"

	eventsourcing "github.com/shortlink-org/shortlink/pkg/eventsourcing/domain/eventsourcing/v1"
)

// EventSourcing - interface for event store
type EventSourcing interface {
	Save(ctx context.Context, events []*eventsourcing.Event) error
	SafeSave(ctx context.Context, events []*eventsourcing.Event) error
	Load(ctx context.Context, aggregateID string) (*eventsourcing.Snapshot, []*eventsourcing.Event, error)

	GetAggregateWithoutSnapshot(ctx context.Context) ([]*eventsourcing.BaseAggregate, error)
	SaveSnapshot(ctx context.Context, snapshot *eventsourcing.Snapshot) error
}

// eventSourcing is the implementation of the event store
type eventSourcing struct {
	repository EventSourcing

	typeStore string
}
