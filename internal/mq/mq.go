package mq

import (
	"context"
	"io"

	"github.com/batazor/shortlink/internal/notify"
)

type MQ interface { // nolint unused
	notify.Subscriber // Observer interface

	// Closer is the interface that wraps the basic Close method.
	io.Closer

	Init(ctx context.Context) error

	Send(message []byte) error
	Subscribe(message chan []byte) error
}
