/*
MQ Endpoint
*/

package api_mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/logger"
	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
	link_application "github.com/batazor/shortlink/internal/services/link/application/link"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	metadata "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"

	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Event struct {
	mq  mq.MQ
	log logger.Logger

	// Application
	service *link_application.Service

	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event
}

func New(mq mq.MQ, log logger.Logger) (*Event, error) {
	event := &Event{
		mq:  mq,
		log: log,
	}

	// Subscribe
	event.SubscribeCQRSGetMetadata(func(ctx context.Context, in *metadata.Meta) error {
		go notify.Publish(ctx, metadata.METHOD_ADD, in, nil)
		return nil
	})

	// Subscribe a new link
	event.SubscribeNewLink(func(ctx context.Context, in *link.Link) error {
		_, err := event.service.Add(ctx, in)
		return err
	})

	return event, nil
}

// Notify ...
func (e *Event) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	// Skip if MQ disabled
	if !viper.GetBool("MQ_ENABLED") {
		return notify.Response{}
	}

	switch event {
	case link.METHOD_ADD:
		return e.add(ctx, payload)
	case link.METHOD_GET:
		panic("implement me")
	case link.METHOD_LIST:
		panic("implement me")
	case link.METHOD_UPDATE:
		panic("implement me")
	case link.METHOD_DELETE:
		panic("implement me")
	}

	return notify.Response{}
}

func (e *Event) add(ctx context.Context, payload interface{}) notify.Response {
	// TODO: send []byte
	msg := payload.(*link.Link) // nolint errcheck
	data, err := proto.Marshal(msg)
	if err != nil {
		return notify.Response{
			Name:    "RESPONSE_MQ_ADD",
			Payload: nil,
			Error:   err,
		}
	}

	err = e.mq.Publish(ctx, link.MQ_EVENT_LINK_CREATED, query.Message{
		Key:     nil,
		Payload: data,
	})
	return notify.Response{
		Name:    "RESPONSE_MQ_ADD",
		Payload: nil,
		Error:   err,
	}
}
