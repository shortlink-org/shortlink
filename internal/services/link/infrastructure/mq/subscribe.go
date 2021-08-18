package api_mq

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	metadata_domain "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"
)

func (e *Event) SubscribeNewLink() error {
	// TODO: refactoring this code
	getNewLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if err := e.mq.Subscribe(link.MQ_EVENT_LINK_NEW, getNewLink); err != nil {
			e.log.Error(err.Error())
		}
	}()

	go func() {
		for {
			msg := <-getNewLink.Chan

			// Convert: []byte to link.Link
			myLink := &link.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				e.log.ErrorWithContext(msg.Context, fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				msg.Context.Done()
				continue
			}

			if _, err := e.service.Add(msg.Context, myLink); err != nil {
				e.log.ErrorWithContext(msg.Context, err.Error())
			}
			msg.Context.Done()
		}
	}()

	return nil
}

func (e *Event) SubscribeCQRSNewLink() error {
	// TODO: refactoring this code
	getCreatedLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if err := e.mq.Subscribe(link.MQ_EVENT_LINK_CREATED, getCreatedLink); err != nil {
			e.log.Error(err.Error())
		}
	}()

	go func() {
		for {
			msg := <-getCreatedLink.Chan

			// Convert: []byte to link.Link
			myLink := &link.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				e.log.ErrorWithContext(msg.Context, fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				msg.Context.Done()
				continue
			}

			go notify.Publish(msg.Context, link.METHOD_ADD, myLink, nil)
			msg.Context.Done()
		}
	}()

	return nil
}

func (e *Event) SubscribeCQRSGetMetadata() {
	// TODO: refactoring this code
	getCQRSGetMetadata := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if err := e.mq.Subscribe(metadata_domain.MQ_CQRS_EVENT, getCQRSGetMetadata); err != nil {
			e.log.Error(err.Error())
		}
	}()

	go func() {
		for {
			msg := <-getCQRSGetMetadata.Chan

			// Convert: []byte to link.Link
			myLink := &metadata_domain.Meta{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				e.log.ErrorWithContext(msg.Context, fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				msg.Context.Done()
				continue
			}

			go notify.Publish(msg.Context, metadata_domain.METHOD_ADD, myLink, nil)
			msg.Context.Done()
		}
	}()
}
