package observer

type Subscriber interface {
	Notify(interface{})
	Close()
}

type Publisher interface {
	run()
	AddSubscriber() chan<- Subscriber
	RemoveSubscriber() chan<- Subscriber
	PublishMessage() chan<- interface{}
	Stop()
}
