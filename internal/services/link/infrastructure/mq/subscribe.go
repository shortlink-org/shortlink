package api_mq

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
	metadata_domain "github.com/batazor/shortlink/internal/services/metadata/domain"
)

func (e *Event) SubscribeCQRSGetMetadata(handler func(ctx context.Context, in *metadata_domain.Meta) error) {
	// TODO: refactoring this code
	getCQRSGetMetadata := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if err := e.mq.Subscribe("shortlink.metadata.cqrs", getCQRSGetMetadata); err != nil {
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

			err := handler(msg.Context, myLink)
			if err != nil {
				e.log.ErrorWithContext(msg.Context, err.Error())
			}
			msg.Context.Done()
		}
	}()
}
