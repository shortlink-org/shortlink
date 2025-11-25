//go:build integration

package metadata_mq

import (
	"context"
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	tckafka "github.com/testcontainers/testcontainers-go/modules/kafka"
	"google.golang.org/protobuf/proto"

	"github.com/shortlink-org/go-sdk/cqrs/bus"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"
	"github.com/shortlink-org/go-sdk/config"
	shortctx "github.com/shortlink-org/go-sdk/context"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/watermill"
	watermill_kafka "github.com/shortlink-org/go-sdk/watermill/backends/kafka"
	cqrs_registry "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/cqrs"
	metadata_uc "github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/metadata"
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

// setupMetadataMQ creates MetadataMQ with real Kafka subscriber
func setupMetadataMQ(t *testing.T, kafkaBroker string) (*Event, *bus.TypeRegistry, cqrsmessage.Marshaler, func(), error) {
	ctx := context.Background()

	// Set environment variables for config
	t.Setenv("MQ_ENABLED", "true")
	t.Setenv("MQ_TYPE", "kafka")
	t.Setenv("WATERMILL_KAFKA_BROKERS", kafkaBroker)
	t.Setenv("WATERMILL_KAFKA_CLIENT_ID", "test-metadata-client")
	t.Setenv("WATERMILL_KAFKA_CONSUMER_GROUP", "test-metadata-group")

	// Create context for logger
	testCtx, cancel, err := shortctx.New()
	require.NoError(t, err)
	t.Cleanup(cancel)

	// Create config
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
	registry, err := cqrs_registry.NewEventRegistry()
	require.NoError(t, err)

	namer := cqrs_registry.NewShortlinkNamer()
	marshaler := cqrs_registry.NewProtoMarshaler(namer)

	// Create mock metadata use case
	metadataUC := &metadata_uc.UC{} // Mock or real implementation

	// Create MetadataMQ
	metadataMQ, err := New(watermillClient.Subscriber, metadataUC)
	require.NoError(t, err)

	cleanup := func() {
		// Cleanup is handled by testcontainers
	}

	return metadataMQ, registry, marshaler, cleanup, nil
}

// publishTestEvent publishes a test event to Kafka
func publishTestEvent(t *testing.T, kafkaBroker string, topic string, event proto.Message) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Version = sarama.V2_8_0_0

	producer, err := sarama.NewSyncProducer([]string{kafkaBroker}, config)
	require.NoError(t, err)
	defer producer.Close()

	// Marshal event to protobuf
	data, err := proto.Marshal(event)
	require.NoError(t, err)

	// Publish message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = producer.SendMessage(msg)
	require.NoError(t, err)
}

