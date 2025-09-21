package rabbit

import (
	"context"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/shortlink-org/shortlink/pkg/logger"
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

// Init initializes the RabbitMQ connection and sets up the channel.
// It also sets up a graceful shutdown mechanism to close the connection and channel
// when the context is done.
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

// close gracefully closes the connection and channel.
func (mq *MQ) close() error {
	if err := mq.conn.Close(); err != nil {
		return err
	}

	//nolint:revive // ignore
	if err := mq.ch.Close(); err != nil {
		return err
	}

	return nil
}

// Check verifies the connection status.
func (mq *MQ) Check(_ context.Context) error {
	if mq.conn.IsClosed() {
		return amqp.ErrClosed
	}

	return nil
}
