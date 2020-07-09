package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"

	bot_type "github.com/batazor/shortlink/internal/bot/type"
	"github.com/batazor/shortlink/internal/notify"
)

type Bot struct {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event

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
	notify.Subscribe(bot_type.METHOD_SEND_NEW_LINK, b)

	return nil
}

func (b *Bot) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case bot_type.METHOD_SEND_NEW_LINK:
		{
			if err := b.Send(payload.(string)); err != nil {
				return notify.Response{
					Error: err,
				}
			}

			return notify.Response{}
		}
	default:
		return notify.Response{}
	}
}

func (b *Bot) Send(message string) error {
	msg := tgbotapi.NewMessage(b.chatId, message)
	if _, err := b.client.Send(msg); err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (b *Bot) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("BOT_TELEGRAM_WEBHOOK", "YOUR_WEBHOOK_URL_HERE") // Your webhook URL
	viper.SetDefault("BOT_TELEGRAM_CHAT_ID", 123)                     // Your chat ID
	viper.SetDefault("BOT_TELEGRAM_DEBUG_MODE", false)                // Debug mode

	b.webhook = viper.GetString("BOT_TELEGRAM_WEBHOOK")
	b.chatId = viper.GetInt64("BOT_TELEGRAM_CHAT_ID")
	b.debugMode = viper.GetBool("BOT_TELEGRAM_DEBUG_MODE")
}
