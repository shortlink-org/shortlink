// Copy from: https://github.com/sirius1024/go-amqp-reconnect

package rabbit

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/batazor/shortlink/internal/pkg/logger"
)

// Connection amqp.Connection wrapper
type Connection struct {
	*amqp.Connection
	delay int

	log logger.Logger
}

// Channel wrap amqp.Connection.Channel, get a auto reconnect channel
func (c *Connection) Channel() (*Channel, error) {
	ch, err := c.Connection.Channel()
	if err != nil {
		return nil, err
	}

	channel := &Channel{
		Channel: ch,

		delay: c.delay,
		log:   c.log,
	}

	go func() {
		for {
			reason, ok := <-channel.Channel.NotifyClose(make(chan *amqp.Error))
			// exit this goroutine if closed by developer
			if !ok || channel.IsClosed() {
				c.log.Error("channel closed")
				err = channel.Close() // close again, ensure closed flag set when connection closed
				if err != nil {
					c.log.Error(err.Error())
				}
				break
			}
			c.log.Error(fmt.Sprintf("channel closed, reason: %v", reason))

			// reconnect if not closed by developer
			for {
				// wait 1s for connection reconnect
				time.Sleep(time.Duration(c.delay) * time.Second)

				ch, err := c.Connection.Channel()
				if err == nil {
					c.log.Info("channel recreate success")
					channel.Channel = ch
					break
				}

				c.log.Error(fmt.Sprintf("channel recreate failed, err: %v", err))
			}
		}
	}()

	return channel, nil
}

// Dial wrap amqp.Dial, dial and get a reconnect connection
func (mq *RabbitMQ) Dial() error {
	conn, err := amqp.Dial(mq.URI)
	if err != nil {
		return err
	}

	mq.conn = &Connection{
		Connection: conn,
		delay:      mq.reconnectTime,
		log:        mq.log,
	}

	go func() {
		for {
			reason, ok := <-mq.conn.NotifyClose(make(chan *amqp.Error))
			// exit this goroutine if closed by developer
			if !ok {
				mq.log.Error("connection closed")
				break
			}
			mq.log.Error(fmt.Sprintf("connection closed, reason: %v", reason))

			// reconnect if not closed by developer
			for {
				// wait 1s for reconnect
				time.Sleep(time.Duration(mq.reconnectTime) * time.Second)

				conn, err := amqp.Dial(mq.URI)
				if err == nil {
					mq.conn.Connection = conn
					mq.log.Info("reconnect success")
					break
				}

				mq.log.Error(fmt.Sprintf("reconnect failed, err: %v", err))
			}
		}
	}()

	return nil
}
