package mq

import (
	"context"
	"io"

	"github.com/batazor/shortlink/internal/notify"
)

type MQ interface { // nolint unused
	// setting
	Init(ctx context.Context) error
	io.Closer // Closer is the interface that wraps the basic Close method.

	// system event
	notify.Subscriber // Observer interface for subscribe on system event

	// mq methods
	Send(message []byte) error
	Subscribe(message chan []byte) error
}
