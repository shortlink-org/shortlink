package nats

import (
	"context"

	"github.com/nats-io/nats.go"

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

func (mq *NATS) Init(_ context.Context) error {
	var err error

	// Connect to a server
	mq.client, err = nats.Connect(nats.DefaultURL)

	return err
}

func (mq *NATS) Close() error {
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
