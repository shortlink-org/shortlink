package nats

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"

	"github.com/batazor/shortlink/internal/mq/query"
	"github.com/batazor/shortlink/internal/notify"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

type Config struct{} // nolint unused

type NATS struct { // nolint unused
	*Config
	client *nats.Conn
}

func (mq *NATS) Init(ctx context.Context) error { // nolint unparam
	var err error

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, mq)

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

func (mq *NATS) Notify(event int, payload interface{}) *notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		msg := payload.(link.Link) // nolint errcheck
		data, err := proto.Marshal(&msg)
		if err != nil {
			return &notify.Response{
				Payload: nil,
				Error:   err,
			}
		}

		err = mq.Publish(query.Message{
			Key:     nil,
			Payload: data,
		})
		return &notify.Response{
			Payload: nil,
			Error:   err,
		}
	case api_type.METHOD_GET:
		panic("implement me")
	case api_type.METHOD_LIST:
		panic("implement me")
	case api_type.METHOD_UPDATE:
		panic("implement me")
	case api_type.METHOD_DELETE:
		panic("implement me")
	}

	return nil
}
