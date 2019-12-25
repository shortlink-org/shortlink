package notify

var (
	subsribers = Notify{
		subsribers: map[int][]Subscriber{},
	}
)

func Subscribe(event int, subscriber Subscriber) { // nolint unused
	subsribers.mux.Lock()
	defer subsribers.mux.Unlock()

	subsribers.subsribers[event] = append(subsribers.subsribers[event], subscriber)
}

func UnSubscribe(event int, subscriber Subscriber) { // nolint unused
	subsribers.mux.Lock()
	defer subsribers.mux.Unlock()

	for _, v := range subsribers.subsribers[event] {
		if subscriber == v {
			delete(subsribers.subsribers, event)
			break
		}
	}
}

func Publish(event int, payload interface{}, cb chan<- interface{}) { // nolint unused
	var responses []*Response
	subsribers.mux.Lock()
	defer subsribers.mux.Unlock()

	if len(subsribers.subsribers[event]) == 0 {
		cb <- nil
	}

	for key := range subsribers.subsribers[event] {
		response := subsribers.subsribers[event][key].Notify(event, payload)
		if response.Error != nil {
			// TODO: Need to add undo operations
			cb <- *response
			return
		}

		responses = append(responses, response)
	}

	// TODO: Send only first success response for simple implementation
	if len(responses) > 0 {
		cb <- *responses[0]
	}
}

func Clean() { // nolint unused
	subsribers = Notify{
		subsribers: map[int][]Subscriber{},
	}
}
