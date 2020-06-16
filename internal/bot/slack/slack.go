package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/spf13/viper"

	bot_type "github.com/batazor/shortlink/internal/bot/type"
	"github.com/batazor/shortlink/internal/notify"
)

type Bot struct {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event

	webhook string
}

func (b *Bot) Init() error {
	// Set configuration
	b.setConfig()

	// Subscribe to Event
	notify.Subscribe(bot_type.METHOD_SEND_NEW_LINK, b)

	return nil
}

func (b *Bot) Notify(ctx context.Context, event int, payload interface{}) notify.Response {
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
	requestBody, err := json.Marshal(map[string]string{
		"text": message,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(b.webhook, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Don't send message to slack")
	}

	return nil
}

// setConfig - set configuration
func (b *Bot) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("BOT_SLACK_WEBHOOK", "YOUR_WEBHOOK_URL_HERE") // Your webhook URL

	b.webhook = viper.GetString("BOT_SLACK_WEBHOOK")
}
