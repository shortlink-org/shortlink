package observer

import (
	"context"
	"sync"
	"testing"

	"github.com/batazor/shortlink/internal/logger"
)

type mockSubscriber struct {
	isClose    bool
	testNotify *func(string)
}

func (s *mockSubscriber) Notify(msg interface{}) {
	(*s.testNotify)(msg.(string))
}
func (s *mockSubscriber) Close() {
	s.isClose = true
}

func TestPublisher(t *testing.T) {
	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}
	ctx = logger.WithLogger(ctx, log)

	pub := NewPublisher(ctx)

	var testFunNotify func(string)

	t.Run("AddSubscriber", func(t *testing.T) {
		cntSub := 50
		wg := sync.WaitGroup{}
		pub.addSubHandler = func(s Subscriber) {
			wg.Done()
		}

		for i := 0; i < cntSub; i++ {
			wg.Add(1)
			go func() {
				sub := mockSubscriber{
					isClose:    false,
					testNotify: &testFunNotify,
				}

				pub.addSubCh <- &sub
			}()
		}

		wg.Wait()

		if cntSub != len(pub.subscribers) {
			t.Errorf("expected cnt sub:%d, got:%d", cntSub, len(pub.subscribers))
		}
	})

	t.Run("PublishMessage", func(t *testing.T) {
		msg := "Test Msg"

		testFunNotify = func(s string) {
			if msg != s {
				t.Errorf("expected:%s got:%s", msg, s)
			}
		}

		pub.PublishMessage() <- msg
	})

	t.Run("RemoveSubscribe", func(t *testing.T) {
		cntSub := 40
		wg := sync.WaitGroup{}
		pub.removeSubHandler = func(s Subscriber) {
			wg.Done()
		}

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				pub.removeSubCh <- pub.subscribers[0]
			}()
		}

		wg.Wait()

		if cntSub != len(pub.subscribers) {
			t.Errorf("expected cnt sub:%d, got:%d", cntSub, len(pub.subscribers))
		}
	})

	pub.stop <- struct{}{}

}
