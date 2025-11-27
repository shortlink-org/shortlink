//go:build integration

package cqrs

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	rpclinkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/rpc/link/v1"
	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// TestE2E_CanonicalTopicNames verifies that all events use canonical topic names
func TestE2E_CanonicalTopicNames(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping E2E integration test")
	}

	ctx := context.Background()

	// Setup Kafka container
	kafkaBroker, kafkaCleanup := setupKafkaContainer(t)
	defer kafkaCleanup()

	// Setup EventBus
	eventBus, eventBusCleanup, err := setupEventBus(t, kafkaBroker)
	require.NoError(t, err)
	defer eventBusCleanup()

	now := time.Now()

	// Test all three event types
	testCases := []struct {
		name  string
		event proto.Message
		topic string
	}{
		{
			name: "LinkCreated",
			event: &linkpb.LinkCreated{
				Url:        "https://example.com/created",
				Hash:       "hash-created",
				Describe:   "Created link",
				CreatedAt:  timestamppb.New(now),
				UpdatedAt:  timestamppb.New(now),
				OccurredAt: timestamppb.New(now),
			},
			topic: "link.link.created.v1",
		},
		{
			name: "LinkUpdated",
			event: &linkpb.LinkUpdated{
				Url:        "https://example.com/updated",
				Hash:       "hash-updated",
				Describe:   "Updated link",
				CreatedAt:  timestamppb.New(now),
				UpdatedAt:  timestamppb.New(now),
				OccurredAt: timestamppb.New(now),
			},
			topic: "link.link.updated.v1",
		},
		{
			name: "LinkDeleted",
			event: &linkpb.LinkDeleted{
				Hash:       "hash-deleted",
				OccurredAt: timestamppb.New(now),
			},
			topic: "link.link.deleted.v1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Publish event
			err := eventBus.Publish(ctx, tc.event)
			require.NoError(t, err, "Should publish %s event", tc.name)

			// Wait for message
			time.Sleep(1 * time.Second)

			// Verify topic name
			messages, consumerCleanup := setupKafkaConsumer(t, kafkaBroker, tc.topic)
			defer consumerCleanup()

			select {
			case msg := <-messages:
				assert.Equal(t, tc.topic, msg.Topic, "Topic should match canonical name: %s", tc.topic)
				assert.NotEmpty(t, msg.Value, "Message should have payload")

			case <-time.After(5 * time.Second):
				t.Fatalf("Timeout waiting for %s event on topic %s", tc.name, tc.topic)
			}
		})
	}
}

