package rabbit

import (
	"context"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
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

func (mq *MQ) Init(ctx context.Context, log logger.Logger) error {
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

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if errClose := mq.close(); errClose != nil {
			log.Error("RabbitMQ close error", field.Fields{
				"error": errClose.Error(),
			})
		}
	}()

	return nil
}

// close - close connection
func (mq *MQ) close() error {
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
