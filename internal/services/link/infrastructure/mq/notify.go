/*
MQ Endpoint
*/

package api_mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	api_domain "github.com/batazor/shortlink/internal/services/api/domain"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	metadata_domain "github.com/batazor/shortlink/internal/services/metadata/domain"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Event struct {
	mq  mq.MQ
	log logger.Logger

	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event
}

func New(mq mq.MQ, log logger.Logger) (*Event, error) {
	event := &Event{
		mq:  mq,
		log: log,
	}

	// Subscribe
	event.SubscribeCQRSGetMetadata(func(ctx context.Context, in *metadata_domain.Meta) error {
		go notify.Publish(ctx, metadata_domain.METHOD_ADD, in, nil)
		return nil
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
	case api_domain.METHOD_ADD:
		return e.add(ctx, payload)
	case api_domain.METHOD_GET:
		panic("implement me")
	case api_domain.METHOD_LIST:
		panic("implement me")
	case api_domain.METHOD_UPDATE:
		panic("implement me")
	case api_domain.METHOD_DELETE:
		panic("implement me")
	}

	return notify.Response{}
}

func (e *Event) add(ctx context.Context, payload interface{}) notify.Response {
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

	err = e.mq.Publish(ctx, "shortlink.link.event", query.Message{
		Key:     nil,
		Payload: data,
	})
	return notify.Response{
		Name:    "RESPONSE_MQ_ADD",
		Payload: nil,
		Error:   err,
	}
}
