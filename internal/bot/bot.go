package bot

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/bot/slack"
	"github.com/batazor/shortlink/internal/bot/smtp"
	"github.com/batazor/shortlink/internal/bot/telegram"
	bot_type "github.com/batazor/shortlink/internal/bot/type"
	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/pkg/link"
)

func (b *Bot) Use(ctx context.Context) { // nolint unused
	// Subscribe to Event
	notify.Subscribe(bot_type.METHOD_NEW_LINK, b)

	// Init slack bot
	go func() {
		slackBot := &slack.Bot{}
		err := slackBot.Init()
		if err != nil {
			return
		}
	}()

	// Init telegram bot
	go func() {
		telegramBot := &telegram.Bot{}
		err := telegramBot.Init()
		if err != nil {
			return
		}
	}()

	// Init SMTP bot
	go func() {
		smtpBot := &smtp.Bot{}
		err := smtpBot.Init()
		if err != nil {
			return
		}
	}()
}

// Notify ...
func (b *Bot) Notify(ctx context.Context, event int, payload interface{}) notify.Response { // nolint unused
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

	notify.Publish(ctx, bot_type.METHOD_SEND_NEW_LINK, payload, nil, "")
}
