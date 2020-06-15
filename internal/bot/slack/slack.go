package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/spf13/viper"
)

type Bot struct {
	webhook string
}

func (b *Bot) Init() error {
	// Set configuration
	b.setConfig()

	return nil
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
	viper.SetDefault("BOT_SLACK_WEBHOOK", "YOUR_WEBHOOK_URL_HERE")

	b.webhook = viper.GetString("BOT_SLACK_WEBHOOK")
}
