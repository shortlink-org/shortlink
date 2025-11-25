/*
MQ Endpoint
*/

package metadata_mq

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/shortlink-org/go-sdk/notify"
	metadata_domain "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
	metadata_uc "github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/metadata"
)

type Event struct {
	publisher  message.Publisher
	subscriber message.Subscriber
	metadataUC *metadata_uc.UC

	// Observer interface for subscribe on system event
	// notify.Subscriber[link.Link]
}

func New(publisher message.Publisher, subscriber message.Subscriber, metadataUC *metadata_uc.UC) (*Event, error) {
	return &Event{
		publisher:  publisher,
		subscriber: subscriber,
		metadataUC: metadataUC,
	}, nil
}

// Notify - notify
func (e *Event) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	// Skip if MQ disabled
	if !viper.GetBool("MQ_ENABLED") {
		return notify.Response[any]{}
	}

	//nolint:revive // not required default case because we handle all cases
	switch event {
	case metadata_domain.METHOD_ADD:
		return e.add(ctx, payload)
	case metadata_domain.METHOD_GET:
		panic("implement me")
	case metadata_domain.METHOD_LIST:
		panic("implement me")
	case metadata_domain.METHOD_UPDATE:
		panic("implement me")
	case metadata_domain.METHOD_DELETE:
		panic("implement me")
	}

	return notify.Response[any]{}
}

func (e *Event) add(ctx context.Context, payload any) notify.Response[any] {
	// TODO: send []byte
	msg, ok := payload.(*metadata_domain.Meta)
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

	msg := message.NewMessage(message.NewUUID(), data)
	err = e.publisher.Publish(metadata_domain.MQ_EVENT_CQRS_NEW, msg)

	return notify.Response[any]{
		Name:    "RESPONSE_MQ_ADD",
		Payload: nil,
		Error:   err,
	}
}
