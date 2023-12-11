package slack

import (
	"bytes"
	"context"
	"net/http"

	"github.com/segmentio/encoding/json"
	"github.com/spf13/viper"

	link "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/boundaries/notification/notify/domain/events"
	"github.com/shortlink-org/shortlink/internal/boundaries/notification/notify/infrastructure"
	http_client "github.com/shortlink-org/shortlink/internal/pkg/http/client"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

type Bot struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]

	webhook string
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
			if err := b.send(ctx, payload.(string)); err != nil { //nolint:forcetypeassert // simple type assertion
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

func (b *Bot) send(ctx context.Context, message string) error {
	requestBody, err := json.Marshal(map[string]string{
		"text": message,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, b.webhook, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	client := http_client.New()

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return infrastructure.ErrSendMessage
	}

	defer resp.Body.Close()

	return nil
}

// setConfig - set configuration
func (b *Bot) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("BOT_SLACK_WEBHOOK", "YOUR_WEBHOOK_URL_HERE") // Your webhook URL

	b.webhook = viper.GetString("BOT_SLACK_WEBHOOK")
}
