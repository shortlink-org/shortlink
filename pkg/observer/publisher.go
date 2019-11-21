package observer

import (
	"context"
	"fmt"
	"github.com/batazor/shortlink/internal/logger"
)

type publisher struct {
	subscribers []Subscriber
	addSubCh    chan Subscriber
	removeSubCh chan Subscriber
	inMsg       chan interface{}
	stop        chan struct{}

	addSubHandler    func(Subscriber)
	removeSubHandler func(Subscriber)

	log logger.Logger
}

func (p *publisher) AddSubscriber() chan<- Subscriber {
	return p.addSubCh
}

func (p *publisher) RemoveSubscribe() chan<- Subscriber {
	return p.removeSubCh
}

func (p *publisher) PublishMessage() chan<- interface{} {
	return p.inMsg
}

func (p *publisher) Stop() {
	close(p.stop)
}

func (p *publisher) onAddSubscriber(sub Subscriber) {
	if p.addSubHandler != nil {
		defer func() {
			if r := recover(); r != nil {
				p.log.Fatal(fmt.Sprintf("%v", r))
			}
		}()

		p.addSubHandler(sub)
	}
}

func (p *publisher) onRemoveSubscriber(sub Subscriber) {
	if p.removeSubHandler != nil {
		defer func() {
			if r := recover(); r != nil {
				p.log.Fatal(fmt.Sprintf("%v", r))
			}
		}()

		p.onRemoveSubscriber(sub)
	}
}

func (p *publisher) run() {
	for {
		select {
		case sub := <-p.addSubCh:
			{
				p.subscribers = append(p.subscribers, sub)
				p.onAddSubscriber(sub)
			}
		case sub := <-p.removeSubCh:
			{
				for i, s := range p.subscribers {
					if sub == s {
						p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
						s.Close()
						p.onRemoveSubscriber(sub)
						break
					}
				}
			}
		case msg := <-p.inMsg:
			{
				for _, sub := range p.subscribers {
					sub.Notify(msg)
				}
			}
		case <-p.stop:
			{
				for _, sub := range p.subscribers {
					sub.Close()
				}

				close(p.addSubCh)
				close(p.removeSubCh)
				close(p.inMsg)

				return
			}
		}
	}
}

func NewPublisher(ctx context.Context) *publisher {
	log := logger.GetLogger(ctx)

	em := publisher{
		addSubCh:    make(chan Subscriber),
		removeSubCh: make(chan Subscriber),
		inMsg:       make(chan interface{}),
		stop:        make(chan struct{}),

		log: log,
	}
	go em.run()

	return &em
}
