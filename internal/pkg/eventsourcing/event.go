package eventsourcing

// Event stores the data for every event
type Event struct {
	ID            string
	AggregateID   string
	AggregateType string
	Version       int
	Type          string
	Payload       interface{}
}
