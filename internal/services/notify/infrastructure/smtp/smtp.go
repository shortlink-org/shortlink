package smtp

import (
	"context"
	"net/smtp"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/notify/domain/events"
)

type Bot struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]

	from string
	pass string
	to   string
	host string
	addr string
}

func (b *Bot) Init() error {
	// Set configuration
	b.setConfig()

	// Subscribe to Event
	notify.Subscribe(events.METHOD_SEND_NEW_LINK, b)

	return nil
}

func (b *Bot) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	switch event {
	case events.METHOD_SEND_NEW_LINK:
		{
			if err := b.send(payload.(string)); err != nil {
				return notify.Response[any]{
					Error: err,
				}
			}

			return notify.Response[any]{}
		}
	default:
		return notify.Response[any]{}
	}
}

func (b *Bot) send(message string) error {
	msg := `To: "Some User" <vitya.login@yandex.ru>
From: "Other User" <vitya.login@yandex.ru>
Subject: Add new link!!

` + message

	auth := smtp.PlainAuth("", b.from, b.pass, b.host)

	err := smtp.SendMail(
		b.addr,
		auth,
		b.from,
		[]string{b.to},
		[]byte(msg),
	)
	if err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (b *Bot) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("BOT_SMTP_FROM", "example@site.com")
	viper.SetDefault("BOT_SMTP_PASS", "YOUR_PASSWORD")
	viper.SetDefault("BOT_SMTP_TO", "EMAIL_USER")
	viper.SetDefault("BOT_SMTP_HOST", "smtp.gmail.com")
	viper.SetDefault("BOT_SMTP_ADDR", "smtp.gmail.com:587")

	b.from = viper.GetString("BOT_SMTP_FROM")
	b.pass = viper.GetString("BOT_SMTP_PASS")
	b.to = viper.GetString("BOT_SMTP_TO")
	b.host = viper.GetString("BOT_SMTP_HOST")
	b.addr = viper.GetString("BOT_SMTP_ADDR")
}
