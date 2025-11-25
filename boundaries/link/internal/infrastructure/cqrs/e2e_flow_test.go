//go:build integration

package cqrs

import (
	"context"
	"reflect"
	"testing"
	"time"

	watermill_std "github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// TestE2E_LinkCreatedToMetadataService tests the full event flow:
// 1. Link Service publishes LinkCreated event
// 2. Event is published to Kafka topic "link.link.created.v1"
// 3. Metadata Service can consume and process the event
func TestE2E_LinkCreatedToMetadataService(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping E2E integration test")
	}

	ctx := context.Background()

	// Setup Kafka container
	kafkaBroker, kafkaCleanup := setupKafkaContainer(t)
	defer kafkaCleanup()

	// Setup Link Service EventBus (publisher)
	linkEventBus, linkCleanup, err := setupEventBus(t, kafkaBroker)
	require.NoError(t, err)
	defer linkCleanup()

	// Setup Metadata Service subscriber (simulated)
	// In real scenario, this would be a separate service
	metadataRegistry, err := NewEventRegistry()
	require.NoError(t, err)

	// Create test event
	now := time.Now()
	event := &linkpb.LinkCreated{
		Url:        "https://example.com/e2e-test",
		Hash:       "e2e-test-hash-789",
		Describe:   "E2E test link description",
		CreatedAt:  timestamppb.New(now),
		UpdatedAt:  timestamppb.New(now),
		OccurredAt: timestamppb.New(now),
	}

	// Step 1: Link Service publishes event
	err = linkEventBus.Publish(ctx, event)
	require.NoError(t, err, "Link Service should publish event successfully")

	// Wait for message to be published
	time.Sleep(2 * time.Second)

	// Step 2: Verify event is in Kafka with correct topic
	topic := "link.link.created.v1"
	messages, consumerCleanup := setupKafkaConsumer(t, kafkaBroker, topic)
	defer consumerCleanup()

	// Step 3: Consume and verify event
	select {
	case msg := <-messages:
		require.NotNil(t, msg, "Message should be received from Kafka")
		assert.Equal(t, topic, msg.Topic, "Topic should match canonical name")
		assert.NotEmpty(t, msg.Value, "Message payload should not be empty")

		// Step 4: Verify event can be unmarshaled using Metadata Service registry
		eventType, ok := metadataRegistry.ResolveEvent(topic)
		require.True(t, ok, "Event should be registered in Metadata Service registry")
		assert.NotNil(t, eventType)

		// Step 5: Unmarshal event (simulating Metadata Service processing)
		eventValue := reflect.New(eventType.Elem()).Interface()
		unmarshaledEvent, ok := eventValue.(proto.Message)
		require.True(t, ok)

		namer := NewShortlinkNamer()
		marshaler := NewProtoMarshaler(namer)

		watermillMsg := message.NewMessage(watermill_std.NewUUID(), msg.Value)
		watermillMsg.Metadata = make(map[string]string)
		for _, header := range msg.Headers {
			watermillMsg.Metadata[string(header.Key)] = string(header.Value)
		}

		err = marshaler.Unmarshal(watermillMsg, unmarshaledEvent)
		require.NoError(t, err, "Metadata Service should unmarshal event successfully")

		// Step 6: Verify event data
		linkCreated, ok := unmarshaledEvent.(*linkpb.LinkCreated)
		require.True(t, ok)
		assert.Equal(t, event.Url, linkCreated.Url, "URL should match")
		assert.Equal(t, event.Hash, linkCreated.Hash, "Hash should match")
		assert.Equal(t, event.Describe, linkCreated.Describe, "Description should match")

		t.Logf("E2E test passed: Link Service published event, Metadata Service can consume and unmarshal it")

	case <-time.After(10 * time.Second):
		t.Fatal("Timeout waiting for message - E2E flow failed")
	}
}

