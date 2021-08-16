/*
MQ Endpoint
*/

package api_mq

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/logger"
	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	link_application "github.com/batazor/shortlink/internal/services/link/application/link"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	metadata "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"

	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Event struct {
	mq  mq.MQ
	log logger.Logger

	// Application
	service *link_application.Service
}

func New(mq mq.MQ, log logger.Logger) (*Event, error) {
	event := &Event{
		mq:  mq,
		log: log,
	}

	// Subscribe
	event.SubscribeCQRSGetMetadata(func(ctx context.Context, in *metadata.Meta) error {
		go notify.Publish(ctx, metadata.METHOD_ADD, in, nil)
		return nil
	})

	// Subscribe a new link
	event.SubscribeNewLink(func(ctx context.Context, in *link.Link) error {
		_, err := event.service.Add(ctx, in)
		return err
	})

	return event, nil
}
