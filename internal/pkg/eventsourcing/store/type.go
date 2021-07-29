package event_store

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/batazor/shortlink/internal/pkg/db"
	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	"github.com/batazor/shortlink/internal/pkg/notify"
)

// EventStore saves the events from an aggregate
type EventStore interface {
	Init(ctx context.Context, db *db.Store) error

	//save(events []*eventsourcing.Event, version int, safe bool) error
	Save(ctx context.Context, events []*eventsourcing.Event) error
	SafeSave(ctx context.Context, events []*eventsourcing.Event) error
	Load(ctx context.Context, aggregateID string) ([]*eventsourcing.Event, error)
}

// Store abstract type
type Repository struct { // nolint unused
	typeStore string

	// Base interface
	EventStore

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}

// Stream represents ordered set of events. Simply speaking stream is an log of the events
// that happened for the specific aggregate/entity.
type Stream struct {
	Id      uuid.UUID // needs to be unique
	Types   uuid.UUID // information about the stream type. It' mostly used to make debugging easier or some optimizations
	Version int       // auto-incremented value used mostly for optimistic concurrency check.
}

/*
	Event table is the main table for Event Sourcing storage.
	It contains the information about events that occurred in the system.
	Each event is stored in separate row.

	They're stored as key/value pair (id + event data) plus additional data like stream id, version, type, creation timestamp.
*/
type Event struct {
	Id        uuid.UUID // unique event identifier
	Payload   string    // event data serialized as JSON
	StreamId  uuid.UUID // id of the stream that event occurred
	Type      uuid.UUID // information about the event type. It' mostly used to make debugging easier or some optimizations.
	Version   int32     // version of the stream at which event occurred used for keeping sequence of the event and for optimistic concurrency check
	CreatedAt time.Time // Timestamp at which event was created. Used to get the state of the stream at exact time.
}
