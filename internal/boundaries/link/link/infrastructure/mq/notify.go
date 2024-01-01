/*
MQ Endpoint
*/

package api_mq

import (
	link_application "github.com/shortlink-org/shortlink/internal/boundaries/link/link/application/link"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
)

type Event struct {
	mq  mq.MQ
	log logger.Logger

	// Application
	service *link_application.UC
}

func New(dataBus mq.MQ, log logger.Logger, service *link_application.UC) (*Event, error) {
	event := &Event{
		mq:  dataBus,
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
