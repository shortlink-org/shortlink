package notify

type Publisher interface { // nolint unused
	Subscribe(event int, subscriber Subscriber)
	UnSubscribe(subscriber Subscriber)
}

type Subscriber interface {
	Notify(event int, payload interface{}) *Response
}

type Notify struct {
	subsribers map[int][]Subscriber
}

type Response struct {
	Payload interface{}
	Error   error
}
