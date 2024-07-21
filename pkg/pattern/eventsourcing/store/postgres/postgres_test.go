//go:build unit || (database && postgres)

package es_postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
	"go.uber.org/goleak"

	db "github.com/shortlink-org/shortlink/pkg/db/postgres"
	eventsourcing "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing/domain/eventsourcing/v1"
)

func TestMain(m *testing.M) {
	// TODO: research how correct close store
	// pgxpool: https://github.com/jackc/pgx/pull/1642
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"),
		goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("sync.runtime_Semacquire"))

	os.Exit(m.Run())
}

var linkUniqId atomic.Int64

func TestPostgres(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	st := &db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("ghcr.io/dbsystel/postgresql-partman", "16", []string{
		"POSTGRESQL_USERNAME=postgres",
		"POSTGRESQL_PASSWORD=shortlink",
		"POSTGRESQL_DATABASE=eventsourcing",
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/eventsourcing?sslmode=disable", resource.GetPort("5432/tcp")))

		errInit := st.Init(ctx)
		if errInit != nil {
			return errInit
		}

		return nil
	}); err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		require.NoError(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	// new event sourcing store
	eventSourcing, err := New(ctx, st)
	if err != nil {
		t.Fatalf("Could not create store: %s", err)
	}

	t.Run("TestEventSourcingSaveAndLoad", func(t *testing.T) {
		eventUID := mustNewV7(t).String()

		// Create a dummy event
		event := &eventsourcing.Event{
			AggregateId:   eventUID,
			AggregateType: "test-type",
			Type:          "test-event",
			Payload:       `{"test": "test-payload"}`,
			Version:       1,
		}

		// Save the event
		err := eventSourcing.Save(ctx, []*eventsourcing.Event{event})
		require.NoError(t, err, "Could not save event")

		// Load the event
		_, events, err := eventSourcing.Load(ctx, eventUID)
		require.NoError(t, err, "Could not load events")

		// Check if the event was saved correctly
		require.Equal(t, 1, len(events), "Expected 1 event")

		loadedEvent := events[0]
		require.Equal(t, event.GetAggregateId(), loadedEvent.GetAggregateId(), "AggregateId does not match")
		require.Equal(t, event.GetAggregateType(), loadedEvent.GetAggregateType(), "AggregateType does not match")
		require.Equal(t, event.GetType(), loadedEvent.GetType(), "Type does not match")
		require.Equal(t, event.GetPayload(), loadedEvent.GetPayload(), "Payload does not match")
		require.Equal(t, event.GetVersion(), loadedEvent.GetVersion(), "Version does not match")
	})

	t.Run("TestEventSourcingSaveMultipleEvents", func(t *testing.T) {
		// Create multiple dummy events with unique aggregate IDs
		events := []*eventsourcing.Event{
			{
				AggregateId:   mustNewV7(t).String(),
				AggregateType: "test-type",
				Type:          "test-event-1",
				Payload:       `{"test": "test-payload-1"}`,
				Version:       1,
			},
			{
				AggregateId:   mustNewV7(t).String(),
				AggregateType: "test-type",
				Type:          "test-event-2",
				Payload:       `{"test": "test-payload-2"}`,
				Version:       1,
			},
		}

		// Save the events
		for _, event := range events {
			err := eventSourcing.Save(ctx, []*eventsourcing.Event{event})
			require.NoError(t, err, "Could not save event")
		}

		// Load the events and check if they were saved correctly
		for _, event := range events {
			_, loadedEvents, err := eventSourcing.Load(ctx, event.GetAggregateId())
			require.NoError(t, err, "Could not load events")
			require.Equal(t, 1, len(loadedEvents), "Expected 1 event")

			loadedEvent := loadedEvents[0]
			require.Equal(t, event.GetAggregateId(), loadedEvent.GetAggregateId(), "AggregateId does not match")
			require.Equal(t, event.GetAggregateType(), loadedEvent.GetAggregateType(), "AggregateType does not match")
			require.Equal(t, event.GetType(), loadedEvent.GetType(), "Type does not match")
			require.Equal(t, event.GetPayload(), loadedEvent.GetPayload(), "Payload does not match")
			require.Equal(t, event.GetVersion(), loadedEvent.GetVersion(), "Version does not match")
		}
	})

	t.Run("TestEventSourcingSaveEventExistingAggregate", func(t *testing.T) {
		eventUID := mustNewV7(t).String()

		// Create a dummy event with an existing aggregate ID but a different version
		event := &eventsourcing.Event{
			AggregateId:   eventUID,
			AggregateType: "test-type",
			Type:          "test-event-3",
			Payload:       `{"test": "test-payload-3"}`,
			Version:       1,
		}

		// Save the event
		err := eventSourcing.Save(ctx, []*eventsourcing.Event{event})
		require.NoError(t, err, "Could not save event")

		// Create a new event with the same aggregate ID but a different version
		event2 := &eventsourcing.Event{
			AggregateId:   eventUID,
			AggregateType: "test-type",
			Type:          "test-event-4",
			Payload:       `{"test": "test-payload-4"}`,
			Version:       2,
		}

		// Save the new event
		err = eventSourcing.Save(ctx, []*eventsourcing.Event{event2})
		require.NoError(t, err, "Could not save event")

		// Load the events
		_, events, err := eventSourcing.Load(ctx, eventUID)
		require.NoError(t, err, "Could not load events")

		// Check if the events were saved correctly
		require.Equal(t, 2, len(events), "Expected 2 events")
	})

	t.Run("TestEventSourcingLoadNoEvents", func(t *testing.T) {
		// Try to load events for a non-existent aggregate ID
		_, events, err := eventSourcing.Load(ctx, mustNewV7(t).String())
		require.NoError(t, err, "Error occurred while loading events")

		// Check if no events were loaded
		require.Equal(t, 0, len(events), "Expected no events")
	})

	t.Run("TestEventSourcingUpdateAggregate", func(t *testing.T) {
		eventUID := mustNewV7(t).String()

		// Create a dummy event
		event := &eventsourcing.Event{
			AggregateId:   eventUID,
			AggregateType: "test-type",
			Type:          "test-event",
			Payload:       `{"test": "test-payload"}`,
			Version:       1,
		}

		// Save the event
		err := eventSourcing.Save(ctx, []*eventsourcing.Event{event})
		require.NoError(t, err, "Could not save event")

		// Update the event
		event.Version = 2
		event.Payload = `{"test": "updated-payload"}`
		err = eventSourcing.Save(ctx, []*eventsourcing.Event{event})
		require.NoError(t, err, "Could not update event")

		// Load the event
		_, events, err := eventSourcing.Load(ctx, eventUID)
		require.NoError(t, err, "Could not load events")

		// Check if the event was updated correctly
		require.Equal(t, 2, len(events), "Expected 1 event")
		require.Equal(t, event.GetAggregateId(), events[1].GetAggregateId(), "AggregateId does not match")
		require.Equal(t, event.GetAggregateType(), events[1].GetAggregateType(), "AggregateType does not match")
		require.Equal(t, event.GetType(), events[1].GetType(), "Type does not match")
		require.Equal(t, event.GetPayload(), events[1].GetPayload(), "Payload does not match")
		require.Equal(t, event.GetVersion(), events[1].GetVersion(), "Version does not match")
	})
}

func mustNewV7(t *testing.T) uuid.UUID {
	id, err := uuid.NewV7()
	require.NoError(t, err)
	return id
}
