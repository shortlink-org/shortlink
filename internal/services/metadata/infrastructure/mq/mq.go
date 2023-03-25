/*
MQ Endpoint
*/

package metadata_mq

import (
	"context"

	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	mq "github.com/shortlink-org/shortlink/internal/pkg/mq/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/v1/query"
	metadata "github.com/shortlink-org/shortlink/internal/services/metadata/domain/metadata/v1"

	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

type Event struct {
	mq *mq.DataBus

	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]
}

func New(mq *mq.DataBus) (*Event, error) {
	return &Event{
		mq: mq,
	}, nil
}

// Notify ...
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

func (e *Event) add(ctx context.Context, payload interface{}) notify.Response[any] {
	// TODO: send []byte
	msg := payload.(*metadata.Meta) // nolint:errcheck
	data, err := proto.Marshal(msg)
	if err != nil {
		return notify.Response[any]{
			Name:    "RESPONSE_MQ_ADD",
			Payload: nil,
			Error:   err,
		}
	}

	err = e.mq.Publish(ctx, metadata.MQ_EVENT_CQRS_NEW, query.Message{
		Key:     nil,
		Payload: data,
	})

	return notify.Response[any]{
		Name:    "RESPONSE_MQ_ADD",
		Payload: nil,
		Error:   err,
	}
}
