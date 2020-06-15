package slack

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

type Bot struct {
	WEBHOOK string
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

	resp, err := http.Post(b.WEBHOOK, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// setConfig - set configuration
func (b *Bot) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("BOT_SLACK_WEBHOOK", "YOUR_WEBHOOK_URL_HERE")

	b.WEBHOOK = viper.GetString("BOT_SLACK_WEBHOOK")
}
