package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/heptiolabs/healthcheck"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/mq/v1/query"
)

type Config struct {
	ConsumerGroup string
	URI           []string
}

type Kafka struct { // nolint:decorder
	*Config
	client   sarama.Client
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup
}

func (mq *Kafka) Init(ctx context.Context) error {
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

	// Create new consumer
	if mq.consumer, err = sarama.NewConsumerGroupFromClient(mq.ConsumerGroup, mq.client); err != nil {
		return err
	}

	// run cron for check connection
	healthcheck.AsyncWithContext(ctx, func() error {
		if len(mq.client.Brokers()) > 0 {
			return nil
		}

		return fmt.Errorf("kafka connection error")
	}, 5*time.Second)

	return nil
}

func (mq *Kafka) Close() error {
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

func (k *Kafka) Publish(ctx context.Context, target string, message query.Message) error {
	_, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic:     target,
		Key:       sarama.StringEncoder(message.Key),
		Value:     sarama.ByteEncoder(message.Payload),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
	})

	return err
}

func (mq *Kafka) Subscribe(target string, message query.Response) error {
	consumer := Consumer{
		ch: message,
	}

	if err := mq.consumer.Consume(context.Background(), []string{target}, &consumer); err != nil {
		return err
	}

	return nil
}

func (mq *Kafka) UnSubscribe(target string) error {
	panic("implement me!")
}

// setConfig - Construct a new Sarama configuration.
func (mq *Kafka) setConfig() (*sarama.Config, error) {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_KAFKA_URI", "localhost:9092")                                                         // Kafka URI
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP", viper.GetString("SERVICE_NAME"))                               // Kafka consumer group
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP_PARTITION_ASSIGNMENT_STRATEGY", sarama.RangeBalanceStrategyName) // Consumer group partition assignment strategy (range, roundrobin, sticky)
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP_OFFSET", sarama.OffsetNewest)                                    // Kafka consumer consume initial offset from oldest

	mq.Config = &Config{
		URI: []string{
			viper.GetString("MQ_KAFKA_URI"),
		},
		ConsumerGroup: viper.GetString("MQ_KAFKA_CONSUMER_GROUP"),
	}

	// sarama config
	config := sarama.NewConfig()
	config.ClientID = viper.GetString("SERVICE_NAME")

	switch viper.GetString("MQ_KAFKA_CONSUMER_GROUP_PARTITION_ASSIGNMENT_STRATEGY") {
	case sarama.StickyBalanceStrategyName:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case sarama.RoundRobinBalanceStrategyName:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case sarama.RangeBalanceStrategyName:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		return nil, fmt.Errorf("Unrecognized consumer group partition assignor: %s", viper.GetString("MQ_KAFKA_CONSUMER_GROUP_PARTITION_ASSIGNMENT_STRATEGY"))
	}

	config.Consumer.Offsets.Initial = viper.GetInt64("MQ_KAFKA_CONSUMER_GROUP_OFFSET")

	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionSnappy
	config.Version = sarama.MaxVersion

	config.Consumer.Return.Errors = true

	return config, nil
}
