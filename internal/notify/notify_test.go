package notify

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	api_type "github.com/batazor/shortlink/pkg/api/type"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

type subscriber struct{}

func (subscriber) Notify(event int, payload interface{}) *Response {
	return &Response{
		Payload: nil,
		Error:   nil,
	}
}

func TestSubscribe(t *testing.T) {
	sub := subscriber{}

	// Subscribe to Event
	Subscribe(0, sub)
	Subscribe(1, sub)
	assert.Equal(t, len(subsribers.subsribers), 2)
}

func TestUnsubscribe(t *testing.T) {
	sub := subscriber{}

	// Subscribe to Event
	Subscribe(0, sub)
	Subscribe(1, sub)
	Subscribe(2, sub)
	assert.Equal(t, len(subsribers.subsribers), 3)

	// Unsubscribe from Event
	UnSubscribe(0, sub)
	assert.Equal(t, len(subsribers.subsribers), 2)
}

func TestPublish(t *testing.T) {
	sub := subscriber{}

	// Subscribe to Event
	Subscribe(0, sub)
	Subscribe(1, sub)
	Subscribe(2, sub)
	assert.Equal(t, len(subsribers.subsribers), 3)

	responseCh := make(chan interface{})

	// Publish
	go Publish(api_type.METHOD_ADD, "hello world", responseCh, "RESPONSE_STORE_ADD")

	c := <-responseCh
	switch r := c.(type) {
	case string:
		assert.Equal(t, "hello world", r)
	}
}

func TestClean(t *testing.T) {
	sub := subscriber{}

	// Subscribe to Event
	Subscribe(0, sub)
	Subscribe(1, sub)
	Subscribe(2, sub)
	assert.Equal(t, len(subsribers.subsribers), 3)

	// Unsubscribe from all Event
	Clean()
	assert.Equal(t, len(subsribers.subsribers), 0)
}
