package rabbit

import (
	"context"
	"fmt"
	"sync"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

type RabbitMQ struct {
	mu sync.Mutex

	config *Config

	log  logger.Logger
	conn *Connection
	ch   *Channel
}

func New(log logger.Logger) *RabbitMQ {
	return &RabbitMQ{
		log:    log,
		config: loadConfig(), // Set configuration
	}
}

func (mq *RabbitMQ) Init(_ context.Context) error {
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

func (mq *RabbitMQ) Close() error {
	if err := mq.conn.Close(); err != nil {
		return err
	}

	if err := mq.ch.Close(); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) Check(ctx context.Context) error {
	if r.conn.IsClosed() {
		return fmt.Errorf("connection is closed")
	}

	return nil
}
