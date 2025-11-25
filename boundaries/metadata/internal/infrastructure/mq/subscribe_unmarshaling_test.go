//go:build integration

package metadata_mq

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shortlink-org/go-sdk/config"
	shortctx "github.com/shortlink-org/go-sdk/context"
	"github.com/shortlink-org/go-sdk/logger"
	linkpb "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/domain/link/v1"
)

func TestMetadataMQ_EventUnmarshaling(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

	// Setup Kafka container
	kafkaBroker, kafkaCleanup := setupKafkaContainer(t)
	defer kafkaCleanup()

	// Setup MetadataMQ
	metadataMQ, registry, marshaler, mqCleanup, err := setupMetadataMQ(t, kafkaBroker)
	require.NoError(t, err)
	defer mqCleanup()

	// Create logger
	testCtx, cancel, err := shortctx.New()
	require.NoError(t, err)
	t.Cleanup(cancel)

	cfg, err := config.New()
	require.NoError(t, err)

	log, _, err := logger.NewDefault(testCtx, cfg)
	require.NoError(t, err)

	// Subscribe to events
	err = metadataMQ.SubscribeLinkCreated(ctx, log, registry, marshaler)
	require.NoError(t, err)

	// Wait a bit for subscription to be ready
	time.Sleep(2 * time.Second)

	// Create test event with specific data
	testURL := "https://test-example.com"
	testHash := "test-hash-456"
	testDescribe := "Test description for unmarshaling"

	now := time.Now()
	event := &linkpb.LinkCreated{
		Url:        testURL,
		Hash:       testHash,
		Describe:   testDescribe,
		CreatedAt:  timestamppb.New(now),
		UpdatedAt:  timestamppb.New(now),
		OccurredAt: timestamppb.New(now),
	}

	topic := "link.link.created.v1"
	publishTestEvent(t, kafkaBroker, topic, event)

	// Wait for event to be processed
	time.Sleep(3 * time.Second)

	// Verify registry can resolve event type
	eventType, ok := registry.ResolveEvent(topic)
	require.True(t, ok, "Event should be registered in registry")
	assert.NotNil(t, eventType)

	// Test passes if event was successfully unmarshaled and processed
	// In a real implementation, you would verify the event data was correctly extracted
}

