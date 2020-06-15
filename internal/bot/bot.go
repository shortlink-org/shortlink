package bot

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/bot/slack"
	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/pkg/link"
)

type Bot struct {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event

	slack slack.Bot
}

func (b *Bot) Use(ctx context.Context) { // nolint unused
	// Subscribe to Event
	notify.Subscribe(METHOD_NEW_LINK, b)

	// Init bot
	err := b.slack.Init()
	if err != nil {
		return
	}
}

// Notify ...
func (b *Bot) Notify(ctx context.Context, event int, payload interface{}) notify.Response { // nolint unused
	switch event {
	case METHOD_NEW_LINK:
		if addLink, ok := payload.(*link.Link); ok {
			b.Send(ctx, addLink)
		}
	}

	return notify.Response{}
}

func (b *Bot) Send(ctx context.Context, link *link.Link) {
	payload := fmt.Sprintf("LINK: %s", link.Url)

	err := b.slack.Send(payload)
	if err != nil {
		return
	}
}
