package notify

import "sync"

type Publisher interface { // nolint unused
	Subscribe(event int, subscriber Subscriber)
	UnSubscribe(subscriber Subscriber)
}

type Subscriber interface {
	Notify(event int, payload interface{}) *Response
}

type Notify struct {
	subsribers map[int][]Subscriber
	mx         sync.RWMutex
}

type Response struct {
	Name    string
	Payload interface{}
	Error   error
}
