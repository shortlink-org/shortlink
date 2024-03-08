// Copy from: https://github.com/sirius1024/go-amqp-reconnect

package rabbit

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

// Connection amqp.Connection wrapper
type Connection struct {
	log logger.Logger
	*amqp.Connection
	delay int
}

// Channel wraps amqp.Connection.Channel to provide an auto-reconnecting channel.
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

				newCh, errConnectToChannel := c.Connection.Channel()
				if errConnectToChannel == nil {
					c.log.Info("channel recreate success")
					channel.Channel = newCh

					break
				}

				c.log.Error(fmt.Sprintf("channel recreate failed, err: %v", errConnectToChannel))
			}
		}
	}()

	return channel, nil
}

// Dial wraps amqp.Dial to establish a connection and set up automatic reconnection
// in case the connection is lost.
func (mq *MQ) Dial() error {
	conn, err := amqp.Dial(mq.config.URI)
	if err != nil {
		return err
	}

	mq.conn = &Connection{
		Connection: conn,
		delay:      mq.config.ReconnectTime,
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
				time.Sleep(time.Duration(mq.config.ReconnectTime) * time.Second)

				conn, err := amqp.Dial(mq.config.URI)
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
