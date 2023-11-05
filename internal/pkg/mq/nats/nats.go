package nats

import (
	"context"

	"github.com/nats-io/nats.go"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

type Config struct{}

type NATS struct {
	*Config
	client *nats.Conn
}

func New() *NATS {
	return &NATS{}
}

func (mq *NATS) Init(ctx context.Context, log logger.Logger) error {
	var err error

	// Connect to a server
	mq.client, err = nats.Connect(nats.DefaultURL)

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if errClose := mq.close(); errClose != nil {
			log.Error("NATS close", field.Fields{
				"error": errClose.Error(),
			})
		}
	}()

	return err
}

// close - close connection
func (mq *NATS) close() error {
	mq.client.Close()
	return nil
}

func (mq *NATS) Publish(_ context.Context, _ string, routingKey, payload []byte) error {
	err := mq.client.Publish(string(routingKey), payload)
	return err
}

func (mq *NATS) Subscribe(_ context.Context, _ string, message query.Response) error {
	_, err := mq.client.Subscribe(string(message.Key), func(m *nats.Msg) {
		message.Chan <- query.ResponseMessage{
			Body: m.Data,
		}
	})
	if err != nil {
		return err
	}

	ch := make(chan *nats.Msg, 64) //nolint:gomnd,revive // TODO: move to config
	_, err = mq.client.ChanSubscribe(string(message.Key), ch)

	if err != nil {
		return err
	}

	for {
		msg := <-ch
		message.Chan <- query.ResponseMessage{
			Body: msg.Data,
		}
	}
}

func (mq *NATS) UnSubscribe(_ string) error {
	panic("implement me!")
}
