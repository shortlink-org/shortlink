package mq

import (
	"context"
	"io"

	"github.com/batazor/shortlink/internal/mq/query"
	"github.com/batazor/shortlink/internal/notify"
)

type MQ interface { // nolint unused
	// setting
	Init(ctx context.Context) error
	io.Closer // Closer is the interface that wraps the basic Close method.

	// system event
	notify.Subscriber // Observer interface for subscribe on system event

	// Pub/Sub a pattern
	Publish(message query.Message) error
	Subscribe(message query.Response) error
	UnSubscribe() error
}

// MQ abstract type
type DataBus struct { // nolint unused
	Databus MQ
	typeMQ  string
}
