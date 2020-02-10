package telegram

import (
	"context"
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
)

func Run(ctx context.Context, log logger.Logger) error {
	viper.SetDefault("TELEGRAM_APITOKEN", "secret")

	bot, err := tgbotapi.NewBotAPI(viper.GetString("TELEGRAM_APITOKEN"))
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Info(fmt.Sprintf("[%s] %s", update.Message.From.UserName, update.Message.Text))
	}

	return nil
}
