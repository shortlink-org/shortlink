package mq

import (
	"context"
)

type MQ interface { // nolint unused
	Init(ctx context.Context) error
	Close() error

	Send(message []byte) error
	Subscribe(message chan []byte)
}
