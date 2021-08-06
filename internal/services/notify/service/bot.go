/*
Bot Service
*/
package service

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/notify/di"
	bot_type "github.com/batazor/shortlink/internal/services/notify/type"
)

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
		if b.MQ != nil {
			if errSubscribe := b.MQ.Subscribe("shortlink", getEventNewLink); errSubscribe != nil {
				return errSubscribe
			}
		}

		return nil
	})

	g.Go(func() error {
		for {
			msg := <-getEventNewLink.Chan

			// Convert: []byte to link.Link
			myLink := &v1.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				b.Log.ErrorWithContext(msg.Context, fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				continue
			}

			b.Log.InfoWithContext(msg.Context, "Get new LINK", field.Fields{"url": myLink.Url})
			notify.Publish(msg.Context, bot_type.METHOD_NEW_LINK, myLink, nil)
		}
	})

	if err := g.Wait(); err != nil {
		b.Log.Error(err.Error())
	}
}

// Notify ...
func (b *Bot) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case bot_type.METHOD_NEW_LINK:
		if addLink, ok := payload.(*v1.Link); ok {
			b.Send(ctx, addLink)
		}
	}

	return notify.Response{}
}

func (b *Bot) Send(ctx context.Context, link *v1.Link) {
	payload := fmt.Sprintf("LINK: %s", link.Url)

	notify.Publish(ctx, bot_type.METHOD_SEND_NEW_LINK, payload, nil)
}
