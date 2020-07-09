package notify

import (
	"context"
	"sync"
)

type Publisher interface { // nolint unused
	Subscribe(event *int, subscriber Subscriber)
	UnSubscribe(subscriber Subscriber)
}

type Subscriber interface {
	Notify(ctx context.Context, event uint32, payload interface{}) Response
}

type Notify struct {
	subsribers map[uint32][]Subscriber
	sync.RWMutex
}

type Response struct {
	Name    string
	Payload interface{}
	Error   error
}
