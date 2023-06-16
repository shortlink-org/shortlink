package notify

import (
	"context"
	"sync"
)

type Publisher[T any] interface {
	Subscribe(event *int, subscriber Subscriber[T])
	UnSubscribe(subscriber Subscriber[T])
}

type Subscriber[T any] interface {
	Notify(ctx context.Context, event uint32, payload T) Response[T]
}

type Notify[T any] struct { // nolint:decorder
	mu            sync.RWMutex
	subscriberMap map[uint32][]Subscriber[T]
}

type Response[T any] struct {
	Payload T
	Error   error
	Name    string
}

type Callback struct { // nolint:decorder
	CB             chan<- interface{}
	ResponseFilter string
}
