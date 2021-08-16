/*
MQ Endpoint
*/

package metadata_mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
	metadata "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"

	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Event struct {
	mq mq.MQ

	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event
}

func New(mq mq.MQ) (*Event, error) {
	return &Event{
		mq: mq,
	}, nil
}

// Notify ...
func (e *Event) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	// Skip if MQ disabled
	if !viper.GetBool("MQ_ENABLED") {
		return notify.Response{}
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

	return notify.Response{}
}

func (e *Event) add(ctx context.Context, payload interface{}) notify.Response {
	// TODO: send []byte
	msg := payload.(*metadata.Meta) // nolint errcheck
	data, err := proto.Marshal(msg)
	if err != nil {
		return notify.Response{
			Name:    "RESPONSE_MQ_ADD",
			Payload: nil,
			Error:   err,
		}
	}

	err = e.mq.Publish(ctx, metadata.MQ_CQRS_EVENT, query.Message{
		Key:     nil,
		Payload: data,
	})
	return notify.Response{
		Name:    "RESPONSE_MQ_ADD",
		Payload: nil,
		Error:   err,
	}
}
