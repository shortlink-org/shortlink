package kafka

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"github.com/heptiolabs/healthcheck"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/mq/query"
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

	// Use a sync.Map to keep track of the ConsumerGroup sessions
	sessions sync.Map
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

	// Check connection
	if len(mq.client.Brokers()) == 0 {
		return sarama.ErrOutOfBrokers
	}

	// run cron for check connection
	healthcheck.AsyncWithContext(ctx, func() error {
		if len(mq.client.Brokers()) > 0 {
			return nil
		}

		return sarama.ErrOutOfBrokers
	}, 5*time.Second) //nolint:mnd // 5s

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
	var errs error

	if mq.client != nil {
		err := mq.client.Close()
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	if mq.producer != nil {
		err := mq.producer.Close()
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	if mq.consumer != nil {
		err := mq.consumer.Close()
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	mq.sessions.Range(func(key, value any) bool {
		sess, ok := value.(*sarama.ConsumerGroupSession)
		if !ok {
			return true
		}

		(*sess).Context().Done()
		mq.sessions.Delete(key)

		return true
	})

	return errs
}

func (mq *Kafka) Publish(_ context.Context, target string, routingKey []byte, payload []byte) error {
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

// Subscribe - subscribe to message
func (mq *Kafka) Subscribe(ctx context.Context, target string, message query.Response) error {
	// Set up a new Sarama consumer group
	consumer := &Consumer{
		ch:    message,
		ready: make(chan bool),
	}

	// OpenTelemetry
	handler := otelsarama.WrapConsumerGroupHandler(consumer)

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			err := mq.consumer.Consume(ctx, []string{target}, handler)
			if err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}

				panic(err)
			}

			// check if context was canceled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}

			consumer.ready = make(chan bool)
		}
	}()

	// Wait until the consumer has been set up
	<-consumer.ready

	// Keep track of sessions to be able to close them
	mq.sessions.Store(target, consumer.session)

	return nil
}

func (mq *Kafka) UnSubscribe(target string) error {
	if session, ok := mq.sessions.Load(target); ok {
		sess, okType := session.(*sarama.ConsumerGroupSession)
		if !okType {
			return nil
		}

		(*sess).Context().Done()
		mq.sessions.Delete(target)
	}

	return nil
}

// setConfig - Construct a new Sarama configuration.
//
// Reference:
// - https://developers.redhat.com/articles/2022/05/03/fine-tune-kafka-performance-kafka-optimization-theorem#the_kafka_optimization_theorem
func (mq *Kafka) setConfig() (*sarama.Config, error) {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_KAFKA_URI", "localhost:9092")                                                         // Kafka URI
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP", viper.GetString("SERVICE_NAME"))                               // Kafka consumer group
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP_PARTITION_ASSIGNMENT_STRATEGY", sarama.RangeBalanceStrategyName) // Consumer group partition assignment strategy (range, roundrobin, sticky)
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP_OFFSET", sarama.OffsetNewest)                                    // Kafka consumer consumes initial offset from oldest
	viper.SetDefault("MQ_KAFKA_PRODUCER_RETRY_MAX", 3)                                                         // Kafka producer retry max
	viper.SetDefault("MQ_KAFKA_SARAMA_VERSION", "MAX")                                                         // Kafka sarama version: MAX, DEFAULT

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
	config.Consumer.Return.Errors = true
	config.Producer.Compression = sarama.CompressionSnappy

	// set sarama version for support redpanda
	switch viper.GetString("MQ_KAFKA_SARAMA_VERSION") {
	case "MAX":
		config.Version = sarama.MaxVersion
	case "DEFAULT":
		config.Version = sarama.DefaultVersion
	}

	// idempotent producer
	config.Producer.Idempotent = true
	if config.Producer.Idempotent {
		if config.Producer.Retry.Max == 0 {
			return nil, sarama.ErrInvalidConfig
		}
		if config.Producer.RequiredAcks != sarama.WaitForAll {
			return nil, sarama.ErrInvalidConfig
		}

		config.Net.MaxOpenRequests = 1
	}

	return config, nil
}
