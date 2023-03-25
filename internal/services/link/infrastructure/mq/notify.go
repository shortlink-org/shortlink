/*
MQ Endpoint
*/

package api_mq

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	mq "github.com/shortlink-org/shortlink/internal/pkg/mq/v1"
	link_application "github.com/shortlink-org/shortlink/internal/services/link/application/link"
)

type Event struct {
	mq  mq.DataBus
	log logger.Logger

	// Application
	service *link_application.Service
}

func New(mq *mq.DataBus, log logger.Logger, service *link_application.Service) (*Event, error) {
	event := &Event{
		mq:  mq,
		log: log,

		// Application
		service: service,
	}

	// Subscribe on metadata
	// event.SubscribeCQRSGetMetadata()

	// Subscribe on metadata
	// err := event.SubscribeCQRSNewLink()
	// if err != nil {
	// 	return nil, err
	// }

	// Subscribe a new link
	// err = event.SubscribeNewLink()
	// if err != nil {
	// 	return nil, err
	// }

	return event, nil
}
