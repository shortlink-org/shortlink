package event_store

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	"github.com/batazor/shortlink/internal/pkg/notify"
)

// EventStore saves the events from an aggregate
type EventStore interface {
	Init(ctx context.Context, db *db.Store) error

	Save(ctx context.Context, events []*eventsourcing.Event) error
	SafeSave(ctx context.Context, events []*eventsourcing.Event) error
	Load(ctx context.Context, aggregateID string) ([]*eventsourcing.Event, error)

	GetAggregateWithoutSnapshot(ctx context.Context) ([]*eventsourcing.BaseAggregate, error)
	SaveSnapshot(ctx context.Context, snapshot *eventsourcing.Snapshot) error
}

// Store abstract type
type Repository struct { // nolint unused
	typeStore string

	// Base interface
	EventStore

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}
