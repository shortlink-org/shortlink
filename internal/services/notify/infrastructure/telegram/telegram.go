package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/notify/domain/events"
)

type Bot struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]

	client *tgbotapi.BotAPI

	webhook   string
	chatId    int64
	debugMode bool
}

func (b *Bot) Init() error {
	var err error

	// Set configuration
	b.setConfig()

	b.client, err = tgbotapi.NewBotAPI(b.webhook)
	if err != nil {
		return err
	}

	// Set debug mode
	b.client.Debug = b.debugMode

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
	msg := tgbotapi.NewMessage(b.chatId, message)
	if _, err := b.client.Send(msg); err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (b *Bot) setConfig() {
	const BOT_TELEGRAM_CHAT_ID = 123

	viper.AutomaticEnv()
	viper.SetDefault("BOT_TELEGRAM_WEBHOOK", "YOUR_WEBHOOK_URL_HERE") // Your webhook URL
	viper.SetDefault("BOT_TELEGRAM_CHAT_ID", BOT_TELEGRAM_CHAT_ID)    // Your chat ID
	viper.SetDefault("BOT_TELEGRAM_DEBUG_MODE", false)                // Debug mode

	b.webhook = viper.GetString("BOT_TELEGRAM_WEBHOOK")
	b.chatId = viper.GetInt64("BOT_TELEGRAM_CHAT_ID")
	b.debugMode = viper.GetBool("BOT_TELEGRAM_DEBUG_MODE")
}
