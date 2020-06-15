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
	Notify(ctx context.Context, event int, payload interface{}) Response
}

type Notify struct {
	subsribers map[int][]Subscriber
	sync.RWMutex
}

type Response struct {
	Name    string
	Payload interface{}
	Error   error
}
