/*
MQ Endpoint
*/

package api_mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/mq"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
)

type Event struct {
	MQ mq.MQ

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}

// Notify ...
func (e *Event) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	// Skip if MQ disabled
	if !viper.GetBool("MQ_ENABLED") {
		return notify.Response{}
	}

	switch link.LinkEvent(event) {
	case link.LinkEvent_ADD:
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
	case link.LinkEvent_GET:
		panic("implement me")
	case link.LinkEvent_LIST:
		panic("implement me")
	case link.LinkEvent_UPDATE:
		panic("implement me")
	case link.LinkEvent_DELETE:
		panic("implement me")
	}

	return notify.Response{}
}
