package nats

import (
	"net/url"
	"sync"

	"github.com/nats-io/nats.go"
)

// Config - configuration
type Config struct {
	URI         *url.URL
	ChannelSize int
}

type MQ struct {
	mu sync.Mutex

	client *nats.Conn
	config *Config

	subscribes map[string]chan *nats.Msg
}
