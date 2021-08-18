/*
MQ Endpoint
*/

package api_mq

import (
	"github.com/batazor/shortlink/internal/pkg/logger"
	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	link_application "github.com/batazor/shortlink/internal/services/link/application/link"
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

	// Subscribe on metadata
	event.SubscribeCQRSGetMetadata()

	// Subscribe on metadata
	event.SubscribeCQRSNewLink()

	// Subscribe a new link
	err := event.SubscribeNewLink()
	if err != nil {
		return nil, err
	}

	return event, nil
}
