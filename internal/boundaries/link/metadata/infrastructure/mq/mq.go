/*
MQ Endpoint
*/

package metadata_mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	link "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	metadata "github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/domain/metadata/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

type Event struct {
	mq mq.MQ

	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]
}

func New(dataBus mq.MQ) (*Event, error) {
	return &Event{
		mq: dataBus,
	}, nil
}

// Notify - notify
func (e *Event) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	// Skip if MQ disabled
	if !viper.GetBool("MQ_ENABLED") {
		return notify.Response[any]{}
	}

	switch event {
	case metadata.METHOD_ADD:
		return e.add(ctx, payload)
	case metadata.METHOD_GET:
		panic("implement me")
	case metadata.METHOD_LIST:
		panic("implement me")
	case metadata.METHOD_UPDATE:
		panic("implement me")
	case metadata.METHOD_DELETE:
		panic("implement me")
	}

	return notify.Response[any]{}
}

func (e *Event) add(ctx context.Context, payload any) notify.Response[any] {
	// TODO: send []byte
	msg, ok := payload.(*metadata.Meta)
	if !ok {
		return notify.Response[any]{
			Name:    "RESPONSE_MQ_ADD",
			Payload: nil,
			Error:   ErrInvalidPayload,
		}
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		return notify.Response[any]{
			Name:    "RESPONSE_MQ_ADD",
			Payload: nil,
			Error:   err,
		}
	}

	err = e.mq.Publish(ctx, metadata.MQ_EVENT_CQRS_NEW, nil, data)

	return notify.Response[any]{
		Name:    "RESPONSE_MQ_ADD",
		Payload: nil,
		Error:   err,
	}
}
