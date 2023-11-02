package rabbit

import (
	"context"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

type MQ struct {
	mu sync.Mutex

	config *Config

	log  logger.Logger
	conn *Connection
	ch   *Channel
}

func New(log logger.Logger) *MQ {
	return &MQ{
		log:    log,
		config: loadConfig(), // Set configuration
	}
}

func (mq *MQ) Init(_ context.Context) error {
	// connect to RabbitMQ server
	err := mq.Dial()
	if err != nil {
		return err
	}

	// create a channel
	mq.ch, err = mq.conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

func (mq *MQ) Close() error {
	if err := mq.conn.Close(); err != nil {
		return err
	}

	if err := mq.ch.Close(); err != nil {
		return err
	}

	return nil
}

func (mq *MQ) Check(_ context.Context) error {
	if mq.conn.IsClosed() {
		return amqp.ErrClosed
	}

	return nil
}
