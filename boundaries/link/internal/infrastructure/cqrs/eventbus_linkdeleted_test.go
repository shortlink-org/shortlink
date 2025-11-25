//go:build integration

package cqrs

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

func TestEventBus_PublishLinkDeleted(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

	// Setup Kafka container
	kafkaBroker, kafkaCleanup := setupKafkaContainer(t)
	defer kafkaCleanup()

	// Setup EventBus
	eventBus, eventBusCleanup, err := setupEventBus(t, kafkaBroker)
	require.NoError(t, err)
	defer eventBusCleanup()

	// Create test event
	event := &linkpb.LinkDeleted{
		Hash:       "test-hash-123",
		OccurredAt: timestamppb.New(time.Now()),
	}

	// Publish event
	err = eventBus.Publish(ctx, event)
	require.NoError(t, err)

	// Wait a bit for message to be published
	time.Sleep(2 * time.Second)

	// Setup consumer to verify message
	topic := "link.link.deleted.v1"
	messages, consumerCleanup := setupKafkaConsumer(t, kafkaBroker, topic)
	defer consumerCleanup()

	// Wait for message with timeout
	select {
	case msg := <-messages:
		require.NotNil(t, msg)
		assert.Equal(t, topic, msg.Topic)
		assert.NotEmpty(t, msg.Value)

	case <-time.After(10 * time.Second):
		t.Fatal("Timeout waiting for message")
	}
}

