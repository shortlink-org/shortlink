package notify

var (
	subsribers = Notify{
		// TODO: use mutex?
		subsribers: map[int][]Subscriber{},
	}
)

func Subscribe(event int, subscriber Subscriber) { // nolint unused
	subsribers.subsribers[event] = append(subsribers.subsribers[event], subscriber)
}

func UnSubscribe(event int, subscriber Subscriber) { // nolint unused
	for i, v := range subsribers.subsribers[event] {
		if subscriber == v {
			subsribers.subsribers[event] = append(subsribers.subsribers[event][:i], subsribers.subsribers[event][i+1:]...)
			break
		}
	}
}

func Publish(event int, payload interface{}, cb chan<- interface{}) { // nolint unused
	var responses []*Response

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
