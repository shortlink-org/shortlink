package nats

import (
	"context"
	"net/url"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/mq/query"
)

func New() *MQ {
	return &MQ{
		subscribes: make(map[string]chan *nats.Msg),
	}
}

func (mq *MQ) Init(ctx context.Context, log logger.Logger) error {
	// Set configuration
	err := mq.setConfig()
	if err != nil {
		return err
	}

	// Connect to a server
	mq.client, err = nats.Connect(mq.config.URI.String())
	if err != nil {
		return err
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if errClose := mq.close(); errClose != nil {
			log.Error("NATS close", field.Fields{
				"error": errClose.Error(),
			})
		}
	}()

	return err
}

// close - drain connection
func (mq *MQ) close() error {
	err := mq.client.Drain()
	if err != nil {
		return err
	}

	return nil
}

// Publish - publish a message
func (mq *MQ) Publish(_ context.Context, _ string, routingKey, payload []byte) error {
	err := mq.client.Publish(string(routingKey), payload)
	if err != nil {
		return err
	}

	return nil
}

// Subscribe - subscribe to message
func (mq *MQ) Subscribe(ctx context.Context, target string, message query.Response) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if _, exists := mq.subscribes[target]; exists {
		return nil
	}

	ch := make(chan *nats.Msg, mq.config.ChannelSize)
	mq.subscribes[target] = ch

	_, err := mq.client.ChanSubscribe(target, ch)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-ch:
				// we can get a nil message if we close the connection
				if msg == nil {
					continue
				}

				message.Chan <- query.ResponseMessage{
					Body: msg.Data,
				}
			}
		}
	}()

	return nil
}

// UnSubscribe - unsubscribe from message
func (mq *MQ) UnSubscribe(name string) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if ch, exists := mq.subscribes[name]; exists {
		close(ch)
		delete(mq.subscribes, name)
	}

	return nil
}

// setConfig - set configuration
func (mq *MQ) setConfig() error {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_NATS_URI", "nats://localhost:4222") // NATS_URI
	//nolint:revive,mnd // ignore magics numbers
	viper.SetDefault("MQ_NATS_CHANNEL_SIZE", 64) // NATS_CHANNEL_SIZE

	// parse uri
	uri, err := url.Parse(viper.GetString("MQ_NATS_URI"))
	if err != nil {
		return err
	}

	// set config
	mq.config = &Config{
		URI:         uri,
		ChannelSize: viper.GetInt("MQ_NATS_CHANNEL_SIZE"),
	}

	return nil
}
