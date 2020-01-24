package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/mq/query"
)

type Config struct { // nolint unused
	URI string
}

type Kafka struct { // nolint unused
	*Config
	client *kafka.Conn
	writer *kafka.Writer
	reader *kafka.Reader
}

func (mq *Kafka) Init(ctx context.Context) error { // nolint unparam
	var err error

	// Set configuration
	mq.setConfig()

	// to produce messages
	topic := "shortlink"
	partition := 0

	if mq.client, err = kafka.DialLeader(context.Background(), "tcp", mq.Config.URI, topic, partition); err != nil {
		return err
	}

	mq.writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{mq.Config.URI},
		Topic:   topic,
	})

	mq.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:       []string{mq.Config.URI},
		GroupID:       "",
		Topic:         topic,
		Partition:     partition,
		Dialer:        nil,
		QueueCapacity: 0,
		MinBytes:      10e3, // 10KB
		MaxBytes:      10e6, // 10MB
	})

	return nil
}

func (mq *Kafka) Close() error {
	var err error

	if mq.client != nil {
		err = mq.client.Close()
	}

	if mq.writer != nil {
		err = mq.writer.Close()
	}

	return err
}

func (k *Kafka) Publish(message query.Message) error {
	_, err := k.client.WriteMessages(
		kafka.Message{
			Key:   message.Key,
			Value: message.Payload,
		},
	)

	return err
}

func (mq *Kafka) Subscribe(message query.Response) error {
	for {
		msg, err := mq.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		message.Chan <- msg.Value
	}
}

func (mq *Kafka) UnSubscribe() error {
	panic("implement me!")
}

// setConfig - set configuration
func (mq *Kafka) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_KAFKA_URI", "localhost:9092")
	mq.Config = &Config{
		URI: viper.GetString("MQ_KAFKA_URI"),
	}
}
