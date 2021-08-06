/*
MQ Endpoint
*/

package api_mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
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

	switch v1.LinkEvent(event) {
	case v1.LinkEvent_LINK_EVENT_ADD:
		// TODO: send []byte
		msg := payload.(*v1.Link) // nolint errcheck
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
	case v1.LinkEvent_LINK_EVENT_GET:
		panic("implement me")
	case v1.LinkEvent_LINK_EVENT_LIST:
		panic("implement me")
	case v1.LinkEvent_LINK_EVENT_UPDATE:
		panic("implement me")
	case v1.LinkEvent_LINK_EVENT_DELETE:
		panic("implement me")
	}

	return notify.Response{}
}
