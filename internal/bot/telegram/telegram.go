package telegram

import (
	"github.com/go-telegram-bot-api"
)

func Run() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err) // You should add better error handling than this!
	}
}
