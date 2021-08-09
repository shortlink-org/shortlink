/*
MQ Endpoint
*/

package api_mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Event struct {
	MQ mq.MQ

	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event
}

// Notify ...
func (e *Event) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	// Skip if MQ disabled
	if !viper.GetBool("MQ_ENABLED") {
		return notify.Response{}
	}

	switch event {
	case v1.METHOD_ADD:
		return e.add(ctx, payload)
	case v1.METHOD_GET:
		panic("implement me")
	case v1.METHOD_LIST:
		panic("implement me")
	case v1.METHOD_UPDATE:
		panic("implement me")
	case v1.METHOD_DELETE:
		panic("implement me")
	}

	return notify.Response{}
}

func (e *Event) add(ctx context.Context, payload interface{}) notify.Response {
	// TODO: send []byte
	msg := payload.(*v1.Meta) // nolint errcheck
	data, err := proto.Marshal(msg)
	if err != nil {
		return notify.Response{
			Name:    "RESPONSE_MQ_ADD",
			Payload: nil,
			Error:   err,
		}
	}

	err = e.MQ.Publish(ctx, "shortlink.metadata.cqrs", query.Message{
		Key:     nil,
		Payload: data,
	})
	return notify.Response{
		Name:    "RESPONSE_MQ_ADD",
		Payload: nil,
		Error:   err,
	}
}
