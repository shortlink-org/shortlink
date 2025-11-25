//go:build integration

package cqrs

import (
	"context"
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	tckafka "github.com/testcontainers/testcontainers-go/modules/kafka"

	"github.com/shortlink-org/go-sdk/cqrs/bus"
	"github.com/shortlink-org/go-sdk/config"
	shortctx "github.com/shortlink-org/go-sdk/context"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/watermill"
	watermill_kafka "github.com/shortlink-org/go-sdk/watermill/backends/kafka"
)

// setupKafkaContainer creates a Kafka test container
func setupKafkaContainer(t *testing.T) (string, func()) {
	ctx := context.Background()

	kafkaContainer, err := tckafka.RunContainer(ctx,
		testcontainers.WithImage("confluentinc/cp-kafka:7.4.0"),
		tckafka.WithClusterID("test-cluster"),
	)
	require.NoError(t, err)

	brokers, err := kafkaContainer.Brokers(ctx)
	require.NoError(t, err)

	cleanup := func() {
		require.NoError(t, kafkaContainer.Terminate(ctx))
	}

	return brokers[0], cleanup
}

// setupEventBus creates EventBus with real Kafka publisher using go-sdk/watermill
func setupEventBus(t *testing.T, kafkaBroker string) (*bus.EventBus, func(), error) {
	ctx := context.Background()

	// Set environment variables for config
	t.Setenv("MQ_ENABLED", "true")
	t.Setenv("MQ_TYPE", "kafka")
	t.Setenv("WATERMILL_KAFKA_BROKERS", kafkaBroker)
	t.Setenv("WATERMILL_KAFKA_CLIENT_ID", "test-link-client")
	t.Setenv("WATERMILL_KAFKA_CONSUMER_GROUP", "test-link-group")

	// Create context for logger
	testCtx, cancel, err := shortctx.New()
	require.NoError(t, err)
	t.Cleanup(cancel)

	// Create config (reads from environment)
	cfg, err := config.New()
	require.NoError(t, err)

	// Create logger
	log, _, err := logger.NewDefault(testCtx, cfg)
	require.NoError(t, err)

	// Create Kafka backend
	backend, err := watermill_kafka.New(ctx, log, cfg)
	require.NoError(t, err)

	// Create watermill client
	watermillClient, err := watermill.New(ctx, log, cfg, backend, nil, nil, []watermill.Option{}...)
	require.NoError(t, err)

	// Create CQRS components
	namer := NewShortlinkNamer()
	marshaler := NewProtoMarshaler(namer)
	eventBus, err := NewEventBus(watermillClient.Publisher, marshaler, namer)
	require.NoError(t, err)

	cleanup := func() {
		// Cleanup is handled by testcontainers
	}

	return eventBus, cleanup, nil
}

// setupKafkaConsumer creates a Kafka consumer for testing
func setupKafkaConsumer(t *testing.T, kafkaBroker string, topic string) (<-chan *sarama.ConsumerMessage, func()) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V2_8_0_0

	consumer, err := sarama.NewConsumer([]string{kafkaBroker}, config)
	require.NoError(t, err)

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	require.NoError(t, err)

	cleanup := func() {
		require.NoError(t, partitionConsumer.Close())
		require.NoError(t, consumer.Close())
	}

	return partitionConsumer.Messages(), cleanup
}

