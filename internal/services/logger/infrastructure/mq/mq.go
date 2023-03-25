/*
MQ Endpoint
*/

package logger_mq

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	mq "github.com/shortlink-org/shortlink/internal/pkg/mq/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/v1/query"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	logger_application "github.com/shortlink-org/shortlink/internal/services/logger/application"
)

type Event struct {
	mq  *mq.DataBus
	log logger.Logger

	service *logger_application.Service
}

func New(mq *mq.DataBus, log logger.Logger, service *logger_application.Service) (*Event, error) {
	if mq == nil {
		return nil, fmt.Errorf("MQ is nil")
	}

	event := &Event{
		mq:  mq,
		log: log,

		service: service,
	}

	// Subscribe
	event.Subscribe()

	return event, nil
}

func (e *Event) Subscribe() {
	getEventNewLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if err := e.mq.Subscribe(v1.MQ_EVENT_LINK_CREATED, getEventNewLink); err != nil {
			e.log.Error(err.Error())
		}
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan

			// Convert: []byte to link.Link
			myLink := &v1.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				e.log.Error(fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				msg.Context.Done()

				continue
			}

			e.service.Log(msg.Context, myLink)
			msg.Context.Done()
		}
	}()
}
