package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

type Bot struct {
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

	return nil
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
	viper.SetDefault("BOT_TELEGRAM_WEBHOOK", "YOUR_WEBHOOK_URL_HERE")
	viper.SetDefault("BOT_TELEGRAM_CHAT_ID", 123)
	viper.SetDefault("BOT_TELEGRAM_DEBUG_MODE", false)

	b.webhook = viper.GetString("BOT_TELEGRAM_WEBHOOK")
	b.chatId = viper.GetInt64("BOT_TELEGRAM_CHAT_ID")
	b.debugMode = viper.GetBool("BOT_TELEGRAM_DEBUG_MODE")
}
