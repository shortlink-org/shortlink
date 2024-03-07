package specification

type Event struct {
	minValidValue int
}

func NewEvent(minValidValue int) *Event {
	return &Event{
		minValidValue: minValidValue,
	}
}

func (e *Event) IsSatisfiedBy(event int) bool {
	return event >= e.minValidValue
}
