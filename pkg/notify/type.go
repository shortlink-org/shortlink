package notify

type Notify interface {
}

type Sudcriber interface {
	Notify(interface{})
	Close()
}

type Publisher interface {
	run()
	AddSubscriber() chan<- Sudcriber
	RemoveSubscriber() chan<- Sudcriber
	PublishMessage() chan<- interface{}
	Stop()
}
