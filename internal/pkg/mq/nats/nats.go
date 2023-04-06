package nats

import (
	"context"

	"github.com/nats-io/nats.go"

	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

type Config struct{}

type NATS struct { // nolint:decorder
	*Config
	client *nats.Conn
}

func New() *NATS {
	return &NATS{}
}

func (mq *NATS) Init(ctx context.Context) error {
	var err error

	// Connect to a server
	mq.client, err = nats.Connect(nats.DefaultURL)

	return err
}

func (mq *NATS) Close() error {
	mq.client.Close()
	return nil
}

func (mq *NATS) Publish(ctx context.Context, target string, message query.Message) error {
	err := mq.client.Publish(string(message.Key), message.Payload)
	return err
}

func (mq *NATS) Subscribe(target string, message query.Response) error {
	_, err := mq.client.Subscribe(string(message.Key), func(m *nats.Msg) {
		message.Chan <- query.ResponseMessage{
			Body: m.Data,
		}
	})
	if err != nil {
		return err
	}

	ch := make(chan *nats.Msg, 64) // nolint:gomnd
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

func (mq *NATS) UnSubscribe(target string) error {
	panic("implement me!")
}
