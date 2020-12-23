/*
Notify system
*/
package notify

import (
	"context"

	"go.uber.org/atomic"
)

var (
	subsribers = Notify{
		subsribers: map[uint32][]Subscriber{},
	}
	eventCounter atomic.Uint32
)

func NewEventID() uint32 {
	eventCounter.Inc()
	return eventCounter.Load()
}

func Subscribe(event uint32, subscriber Subscriber) { // nolint unused
	subsribers.Lock()
	subsribers.subsribers[event] = append(subsribers.subsribers[event], subscriber)
	subsribers.Unlock()
}

func UnSubscribe(event uint32, subscriber Subscriber) { // nolint unused
	subsribers.Lock()
	defer subsribers.Unlock()

	for _, v := range subsribers.subsribers[event] {
		if subscriber == v {
			delete(subsribers.subsribers, event)
			break
		}
	}
}

// Publish - add new event
func Publish(ctx context.Context, event uint32, payload interface{}, cb *Callback) {
	responses := map[string]Response{}
	subsribers.RLock()
	defer subsribers.RUnlock()

	if len(subsribers.subsribers[event]) == 0 && cb != nil {
		cb.CB <- nil
	}

	// send event to all subscribes
	for key := range subsribers.subsribers[event] {
		response := subsribers.subsribers[event][key].Notify(ctx, event, payload)

		// TODO: How to handle errors?
		//if response.Error != nil && cb != nil {
		//	cb.CB <- response
		//	return
		//}

		if response.Name != "" {
			responses[response.Name] = response
		}
	}

	// TODO: Send only first success response for simple implementation
	if cb != nil && responses[cb.ResponseFilter].Name != "" {
		cb.CB <- responses[cb.ResponseFilter]
	}
}

func Clean() { // nolint unused
	subsribers = Notify{
		subsribers: map[uint32][]Subscriber{},
	}
}
