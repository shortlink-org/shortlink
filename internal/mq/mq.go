package mq

import (
	"context"

	"github.com/batazor/shortlink/internal/notify"
)

type MQ interface { // nolint unused
	notify.Subscriber // Observer interface

	Init(ctx context.Context) error
	Close() error

	Send(message []byte) error
	Subscribe(message chan []byte) error
}
