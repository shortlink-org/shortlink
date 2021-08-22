/*
MQ Endpoint
*/

package logger_mq

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/logger"
	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	logger_application "github.com/batazor/shortlink/internal/services/logger/application"
)

type Event struct {
	mq  mq.MQ
	log logger.Logger

	service *logger_application.Service
}

func New(mq mq.MQ, log logger.Logger, service *logger_application.Service) (*Event, error) {
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
