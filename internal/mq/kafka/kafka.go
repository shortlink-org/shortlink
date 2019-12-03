package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Config struct{} // nolint unused

type Kafka struct { // nolint unused
	*Config
	client *kafka.Conn
}

func (mq *Kafka) Init(ctx context.Context) error {
	var err error

	// to produce messages
	topic := "shortlink"
	partition := 0

	if mq.client, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition); err != nil {
		return err
	}

	if err = mq.client.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return err
	}

	return nil
}

func (mq *Kafka) Close() error {
	return nil
}

func (mq *Kafka) Send(message []byte) error {
	_, err := mq.client.WriteMessages(
		kafka.Message{
			Key:   []byte("TEST"),
			Value: message,
		},
	)

	return err
}

func (mq *Kafka) Subscribe(message chan []byte) error {
	msg, err := mq.client.ReadMessage(10e3) // fetch 10KB max
	if err != nil {
		return err
	}

	for {
		message <- msg.Value
	}
}
