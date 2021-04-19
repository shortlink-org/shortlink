/*
MQ Endpoint
*/

package api_mq

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/mq"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain/link"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

type Event struct {
	MQ mq.MQ

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}

// Notify ...
func (e *Event) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case api_type.METHOD_ADD:
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

		err = e.MQ.Publish(ctx, "shortlink", query.Message{
			Key:     nil,
			Payload: data,
		})
		return notify.Response{
			Name:    "RESPONSE_MQ_ADD",
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

	return notify.Response{}
}
