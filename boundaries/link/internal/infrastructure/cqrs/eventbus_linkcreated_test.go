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

func TestEventBus_PublishLinkCreated(t *testing.T) {
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
	now := time.Now()
	event := &linkpb.LinkCreated{
		Url:        "https://example.com",
		Hash:       "test-hash-123",
		Describe:   "Test link description",
		CreatedAt:  timestamppb.New(now),
		UpdatedAt:  timestamppb.New(now),
		OccurredAt: timestamppb.New(now),
	}

	// Publish event
	err = eventBus.Publish(ctx, event)
	require.NoError(t, err)

	// Wait a bit for message to be published
	time.Sleep(2 * time.Second)

	// Setup consumer to verify message
	topic := "link.link.created.v1"
	messages, consumerCleanup := setupKafkaConsumer(t, kafkaBroker, topic)
	defer consumerCleanup()

	// Wait for message with timeout
	select {
	case msg := <-messages:
		require.NotNil(t, msg)
		assert.Equal(t, topic, msg.Topic)
		assert.NotEmpty(t, msg.Value)

		// Verify event can be unmarshaled
		registry, err := NewEventRegistry()
		require.NoError(t, err)

		eventType, ok := registry.ResolveEvent(topic)
		require.True(t, ok, "Event should be registered in registry")

		// Create new instance of event type
		eventValue := reflect.New(eventType.Elem()).Interface()
		unmarshaledEvent, ok := eventValue.(proto.Message)
		require.True(t, ok)

		// Unmarshal using ProtoMarshaler
		namer := NewShortlinkNamer()
		marshaler := NewProtoMarshaler(namer)

		watermillMsg := message.NewMessage(watermill_std.NewUUID(), msg.Value)
		watermillMsg.Metadata = make(map[string]string)
		for _, header := range msg.Headers {
			watermillMsg.Metadata[string(header.Key)] = string(header.Value)
		}

		err = marshaler.Unmarshal(watermillMsg, unmarshaledEvent)
		require.NoError(t, err)

		// Verify event data
		linkCreated, ok := unmarshaledEvent.(*linkpb.LinkCreated)
		require.True(t, ok)
		assert.Equal(t, event.Url, linkCreated.Url)
		assert.Equal(t, event.Hash, linkCreated.Hash)
		assert.Equal(t, event.Describe, linkCreated.Describe)

	case <-time.After(10 * time.Second):
		t.Fatal("Timeout waiting for message")
	}
}

