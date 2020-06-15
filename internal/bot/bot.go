package bot

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/bot/slack"
	"github.com/batazor/shortlink/internal/bot/smtp"
	"github.com/batazor/shortlink/internal/bot/telegram"
	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/pkg/link"
)

func (b *Bot) Use(ctx context.Context) { // nolint unused
	// Subscribe to Event
	notify.Subscribe(METHOD_NEW_LINK, b)

	// Init slack bot
	slackBot := &slack.Bot{}
	err := slackBot.Init()
	if err != nil {
		return
	}
	b.services = append(b.services, slackBot)

	// Init telegram bot
	telegramBot := &telegram.Bot{}
	err = telegramBot.Init()
	if err != nil {
		return
	}
	b.services = append(b.services, telegramBot)

	// Init SMTP bot
	smtpBot := &smtp.Bot{}
	err = smtpBot.Init()
	if err != nil {
		return
	}
	b.services = append(b.services, smtpBot)
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

	// Send message
	for _, service := range b.services {
		err := service.Send(payload)
		if err != nil {
			continue
		}
	}
}
