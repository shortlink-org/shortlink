/*
Bot Service
*/
package application

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	mq "github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/notify/di"
	bot_type "github.com/batazor/shortlink/internal/services/notify/type"
)

func New(mq mq.MQ, log logger.Logger) (*Bot, error) {
	return &Bot{
		mq:  nil,
		log: nil,
	}, nil
}

func (b *Bot) Use(ctx context.Context) { // nolint unused
	// Subscribe to Event
	notify.Subscribe(bot_type.METHOD_NEW_LINK, b)

	// Init slack bot
	_, _, err := di.InitializeFullBotService()
	if err != nil {
		panic(err)
	}

	// TODO: refactoring this code
	getEventNewLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	g := errgroup.Group{}

	g.Go(func() error {
		if b.mq != nil {
			if errSubscribe := b.mq.Subscribe(link.MQ_EVENT_LINK_CREATED, getEventNewLink); errSubscribe != nil {
				return errSubscribe
			}
		}

		return nil
	})

	g.Go(func() error {
		for {
			msg := <-getEventNewLink.Chan

			// Convert: []byte to link.Link
			myLink := &link.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				b.log.ErrorWithContext(msg.Context, fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				continue
			}

			b.log.InfoWithContext(msg.Context, "Get new LINK", field.Fields{"url": myLink.Url})
			notify.Publish(msg.Context, bot_type.METHOD_NEW_LINK, myLink, nil)
		}
	})

	if err := g.Wait(); err != nil {
		b.log.Error(err.Error())
	}
}

// Notify ...
func (b *Bot) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case bot_type.METHOD_NEW_LINK:
		if addLink, ok := payload.(*link.Link); ok {
			b.Send(ctx, addLink)
		}
	}

	return notify.Response{}
}

func (b *Bot) Send(ctx context.Context, in *link.Link) {
	payload := fmt.Sprintf("LINK: %s", in.Url)

	notify.Publish(ctx, bot_type.METHOD_SEND_NEW_LINK, payload, nil)
}
