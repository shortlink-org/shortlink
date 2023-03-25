//go:build unit
// +build unit

package notify

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

var (
	METHOD_ADD    = NewEventID()
	METHOD_GET    = NewEventID()
	METHOD_LIST   = NewEventID()
	METHOD_UPDATE = NewEventID()
	METHOD_DELETE = NewEventID()
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

type mockSubscriber[T any] struct {
	Payload T
}

func (mockSubscriber[T]) Notify(ctx context.Context, event uint32, payload T) Response[T] {
	return Response[T]{
		Name:    "RESPONSE_STORE_ADD",
		Payload: payload,
		Error:   nil,
	}
}

func TestSubscribe(t *testing.T) {
	sub := mockSubscriber[any]{}

	// Subscribe to Event
	Subscribe(0, sub)
	Subscribe(1, sub)
	assert.Equal(t, len(subscribers.subscriberMap), 2)
}

func TestUnsubscribe(t *testing.T) {
	sub := mockSubscriber[any]{}

	// Subscribe to Event
	Subscribe(0, sub)
	Subscribe(1, sub)
	Subscribe(2, sub)
	assert.Equal(t, len(subscribers.subscriberMap), 3)

	// Unsubscribe from Event
	UnSubscribe(0, sub)
	assert.Equal(t, len(subscribers.subscriberMap), 2)
}

// TODO: Need kafka-service or add mock
func TestPublish(t *testing.T) {
	ctx := context.Background()
	sub := mockSubscriber[any]{}

	// Subscribe to Event
	Subscribe(METHOD_ADD, sub)
	Subscribe(METHOD_GET, sub)
	Subscribe(METHOD_LIST, sub)
	Subscribe(METHOD_UPDATE, sub)
	Subscribe(METHOD_DELETE, sub)
	assert.Equal(t, len(subscribers.subscriberMap), 5)

	responseCh := make(chan interface{})

	// Publish
	go Publish(ctx, METHOD_ADD, "hello world", &Callback{responseCh, "RESPONSE_STORE_ADD"})

	select {
	case c := <-responseCh:
		{
			switch r := c.(type) {
			case string:
				assert.Equal(t, "hello world", r)
			}
		}
	case <-time.After(1 * time.Millisecond):
		t.Error("timeout")
		return
	}
}

func TestClean(t *testing.T) {
	sub := mockSubscriber[any]{}

	// Subscribe to Event
	Subscribe(METHOD_ADD, sub)
	Subscribe(METHOD_GET, sub)
	Subscribe(METHOD_LIST, sub)
	Subscribe(METHOD_UPDATE, sub)
	Subscribe(METHOD_DELETE, sub)
	assert.Equal(t, len(subscribers.subscriberMap), 5)

	// Unsubscribe from all Event
	Clean()
	assert.Equal(t, len(subscribers.subscriberMap), 0)
}
