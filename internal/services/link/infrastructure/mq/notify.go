/*
MQ Endpoint
*/

package api_mq

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/logger"
	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	link_application "github.com/batazor/shortlink/internal/services/link/application/link"
	metadata "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"

	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Event struct {
	mq  mq.MQ
	log logger.Logger

	// Application
	service *link_application.Service
}

func New(mq mq.MQ, log logger.Logger, service *link_application.Service) (*Event, error) {
	event := &Event{
		mq:  mq,
		log: log,

		// Application
		service: service,
	}

	// Subscribe
	event.SubscribeCQRSGetMetadata(func(ctx context.Context, in *metadata.Meta) error {
		go notify.Publish(ctx, metadata.METHOD_ADD, in, nil)
		return nil
	})

	// Subscribe a new link
	err := event.SubscribeNewLink()
	if err != nil {
		return nil, err
	}

	return event, nil
}
