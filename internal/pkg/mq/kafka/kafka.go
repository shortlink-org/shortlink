package kafka

import (
	"context"
	"time"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"github.com/heptiolabs/healthcheck"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

type Config struct {
	ConsumerGroup string
	URI           []string
}

type Kafka struct {
	*Config
	client   sarama.Client
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup
}

func (mq *Kafka) Init(ctx context.Context, log logger.Logger) error {
	var err error

	// Set configuration
	config, err := mq.setConfig()
	if err != nil {
		return err
	}

	if mq.client, err = sarama.NewClient(mq.Config.URI, config); err != nil {
		return err
	}

	// Create new producer
	if mq.producer, err = sarama.NewSyncProducerFromClient(mq.client); err != nil {
		return err
	}

	// OpenTelemetry
	mq.producer = otelsarama.WrapSyncProducer(config, mq.producer)

	// Create new consumer
	if mq.consumer, err = sarama.NewConsumerGroupFromClient(mq.ConsumerGroup, mq.client); err != nil {
		return err
	}

	// run cron for check connection
	healthcheck.AsyncWithContext(ctx, func() error {
		if len(mq.client.Brokers()) > 0 {
			return nil
		}

		return sarama.ErrOutOfBrokers
	}, 5*time.Second) //nolint:gomnd // 5s

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if errClose := mq.close(); errClose != nil {
			log.Error("Kafka close error", field.Fields{
				"error": errClose.Error(),
			})
		}
	}()

	return nil
}

// close - Close all connections
func (mq *Kafka) close() error {
	var err error

	if mq.client != nil {
		err = mq.client.Close()
	}

	if mq.producer != nil {
		err = mq.producer.Close()
	}

	if mq.consumer != nil {
		err = mq.consumer.Close()
	}

	return err
}

func (mq *Kafka) Publish(_ context.Context, target string, routingKey, payload []byte) error {
	_, _, err := mq.producer.SendMessage(&sarama.ProducerMessage{
		Topic:     target,
		Key:       sarama.StringEncoder(routingKey),
		Value:     sarama.ByteEncoder(payload),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
	})

	return err
}

func (mq *Kafka) Subscribe(ctx context.Context, target string, message query.Response) error {
	consumer := &Consumer{
		ch: message,
	}

	// OpenTelemetry
	handler := otelsarama.WrapConsumerGroupHandler(consumer)

	err := mq.consumer.Consume(ctx, []string{target}, handler)
	if err != nil {
		return err
	}

	return nil
}

func (mq *Kafka) UnSubscribe(_ string) error {
	panic("implement me!")
}

// setConfig - Construct a new Sarama configuration.
func (mq *Kafka) setConfig() (*sarama.Config, error) {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_KAFKA_URI", "localhost:9092")                                                         // Kafka URI
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP", viper.GetString("SERVICE_NAME"))                               // Kafka consumer group
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP_PARTITION_ASSIGNMENT_STRATEGY", sarama.RangeBalanceStrategyName) // Consumer group partition assignment strategy (range, roundrobin, sticky)
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP_OFFSET", sarama.OffsetNewest)                                    // Kafka consumer consumes initial offset from oldest
	viper.SetDefault("MQ_KAFKA_PRODUCER_RETRY_MAX", 5)                                                         // Kafka producer retry max

	mq.Config = &Config{
		URI: []string{
			viper.GetString("MQ_KAFKA_URI"),
		},
		ConsumerGroup: viper.GetString("MQ_KAFKA_CONSUMER_GROUP"),
	}

	// sarama config
	config := sarama.NewConfig()
	config.ClientID = viper.GetString("SERVICE_NAME")

	strategy := viper.GetString("MQ_KAFKA_CONSUMER_GROUP_PARTITION_ASSIGNMENT_STRATEGY")
	switch strategy {
	case sarama.StickyBalanceStrategyName:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case sarama.RoundRobinBalanceStrategyName:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case sarama.RangeBalanceStrategyName:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		return nil, sarama.ErrConsumerCoordinatorNotAvailable
	}

	config.Consumer.Offsets.Initial = viper.GetInt64("MQ_KAFKA_CONSUMER_GROUP_OFFSET")

	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = viper.GetInt("MQ_KAFKA_PRODUCER_RETRY_MAX")
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionSnappy
	config.Version = sarama.MaxVersion

	config.Consumer.Return.Errors = true

	return config, nil
}
