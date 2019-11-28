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

func (mq *Kafka) Init() error {
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
		kafka.Message{Value: message},
	)

	return err
}

func (mq *Kafka) Subscribe(message chan []byte) {
	batch := mq.client.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	b := make([]byte, 10e3)                 // 10KB max per message

	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}

		message <- b
	}
}
