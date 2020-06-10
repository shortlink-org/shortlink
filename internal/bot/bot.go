package bot

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/pkg/link"
)

type Bot struct {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}

func (b *Bot) Use(ctx context.Context) { // nolint unused
	// Subscribe to Event
	notify.Subscribe(METHOD_NEW_LINK, b)
}

// Notify ...
func (b *Bot) Notify(ctx context.Context, event int, payload interface{}) *notify.Response { // nolint unused
	switch event {
	case METHOD_NEW_LINK:
		if addLink, ok := payload.(*link.Link); ok {
			b.Add(ctx, addLink)
		}
	}

	return nil
}

func (b *Bot) Add(ctx context.Context, link *link.Link) {
	fmt.Println("LINK: ", link.Url)
}
