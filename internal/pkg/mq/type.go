package mq

import (
	"context"
	"io"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
)

// MQ - common interface of DataBus
type MQ interface { // nolint unused
	// setting
	Init(context.Context) error
	io.Closer // Closer is the interface that wraps the basic Close method.

	// Pub/Sub a pattern
	Publish(ctx context.Context, target string, message query.Message) error
	Subscribe(target string, message query.Response) error
	UnSubscribe(target string) error
}

// DataBus abstract type
type DataBus struct { // nolint unused
	mq     MQ
	typeMQ string
}
