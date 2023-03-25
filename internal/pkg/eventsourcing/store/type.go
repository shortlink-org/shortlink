package event_store

import (
	"context"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// EventStore saves the events from an aggregate
type EventStore interface {
	Init(ctx context.Context, db *db.Store) error

	Save(ctx context.Context, events []*eventsourcing.Event) error
	SafeSave(ctx context.Context, events []*eventsourcing.Event) error
	Load(ctx context.Context, aggregateID string) (*eventsourcing.Snapshot, []*eventsourcing.Event, error)

	GetAggregateWithoutSnapshot(ctx context.Context) ([]*eventsourcing.BaseAggregate, error)
	SaveSnapshot(ctx context.Context, snapshot *eventsourcing.Snapshot) error
}

// Store abstract type
type Repository struct {
	typeStore string

	// Base interface
	EventStore

	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]
}
