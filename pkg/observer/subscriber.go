package observer

import "io"

type subscriber struct {
	in    chan interface{}
	stop  chan struct{}
	store io.Writer
}

func (s subscriber) Notify(msg interface{}) {
	s.in <- msg
}

func (s subscriber) Close() {
	close(s.stop)
}

func (s subscriber) run() {
	for {
		select {
		case msg := <-s.in:
			{
				s.store.Write([]byte(msg.(string)))
			}
		case <-s.stop:
			{
				close(s.in)
				return
			}
		}
	}
}

func NewSubscriber(w io.Writer) subscriber {
	sub := subscriber{
		in:    make(chan interface{}),
		stop:  make(chan struct{}),
		store: w,
	}
	go sub.run()

	return sub
}
