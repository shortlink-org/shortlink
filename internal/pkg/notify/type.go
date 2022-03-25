package notify

import (
	"context"
	"sync"
)

type Publisher interface { // nolint:unused
	Subscribe(event *int, subscriber Subscriber)
	UnSubscribe(subscriber Subscriber)
}

type Subscriber interface { // nolint:decorder
	Notify(ctx context.Context, event uint32, payload interface{}) Response
}

type Notify struct { // nolint:decorder
	subsribers map[uint32][]Subscriber
	sync.RWMutex
}

type Response struct { // nolint:decorder
	Name    string
	Payload interface{}
	Error   error
}

type Callback struct { // nolint:decorder
	CB             chan<- interface{}
	ResponseFilter string
}
