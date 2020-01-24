package nats

import (
	"context"

	"github.com/nats-io/nats.go"

	"github.com/batazor/shortlink/internal/mq/query"
)

type Config struct{} // nolint unused

type NATS struct { // nolint unused
	*Config
	client *nats.Conn
}

func (mq *NATS) Init(ctx context.Context) error { // nolint unparam
	var err error

	// Connect to a server
	mq.client, err = nats.Connect(nats.DefaultURL)

	return err
}

func (mq *NATS) Close() error { // nolint unparam
	mq.client.Close()
	return nil
}

func (mq *NATS) Publish(message query.Message) error {
	err := mq.client.Publish(string(message.Key), message.Payload)
	return err
}

func (mq *NATS) Subscribe(message query.Response) error {
	_, err := mq.client.Subscribe(string(message.Key), func(m *nats.Msg) {
		message.Chan <- m.Data
	})
	if err != nil {
		return err
	}

	ch := make(chan *nats.Msg, 64)
	_, err = mq.client.ChanSubscribe(string(message.Key), ch)
	if err != nil {
		return err
	}

	for {
		msg := <-ch
		message.Chan <- msg.Data
	}
}

func (mq *NATS) UnSubscribe() error {
	panic("implement me!")
}
