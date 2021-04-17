/*
Bot Service
*/
package service

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain/link"
	"github.com/batazor/shortlink/internal/services/bot/di"
	bot_type "github.com/batazor/shortlink/internal/services/bot/type"
)

func (b *Bot) Use(_ context.Context) { // nolint unused
	// Subscribe to Event
	notify.Subscribe(bot_type.METHOD_NEW_LINK, b)

	// Init slack bot
	_, _, err := di.InitializeFullBotService()
	if err != nil {
		panic(err)
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

func (b *Bot) Send(ctx context.Context, link *link.Link) {
	payload := fmt.Sprintf("LINK: %s", link.Url)

	notify.Publish(ctx, bot_type.METHOD_SEND_NEW_LINK, payload, nil)
}
