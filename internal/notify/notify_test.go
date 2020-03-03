package notify

import (
	"testing"
	"time"

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

// TODO: Need kafka-service or add mock
func TestPublish(t *testing.T) {
	sub := subscriber{}

	// Subscribe to Event
	Subscribe(api_type.METHOD_ADD, sub)
	Subscribe(api_type.METHOD_GET, sub)
	Subscribe(api_type.METHOD_LIST, sub)
	Subscribe(api_type.METHOD_UPDATE, sub)
	Subscribe(api_type.METHOD_DELETE, sub)
	assert.Equal(t, len(subsribers.subsribers), 5)

	responseCh := make(chan interface{})

	// Publish
	go Publish(api_type.METHOD_ADD, "hello world", responseCh, "RESPONSE_STORE_ADD")

	select {
	case c := <-responseCh:
		{
			switch r := c.(type) {
			case string:
				assert.Equal(t, "hello world", r)
			}
		}
	case <-time.After(1 * time.Millisecond):
		return
	}
}

func TestClean(t *testing.T) {
	sub := subscriber{}

	// Subscribe to Event
	Subscribe(api_type.METHOD_ADD, sub)
	Subscribe(api_type.METHOD_GET, sub)
	Subscribe(api_type.METHOD_LIST, sub)
	Subscribe(api_type.METHOD_UPDATE, sub)
	Subscribe(api_type.METHOD_DELETE, sub)
	assert.Equal(t, len(subsribers.subsribers), 5)

	// Unsubscribe from all Event
	Clean()
	assert.Equal(t, len(subsribers.subsribers), 0)
}
