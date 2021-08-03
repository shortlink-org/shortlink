package v1

import (
	"context"
)

// AggregateHandler defines the methods to process commands
type AggregateHandler interface {
	ApplyChange(event *Event) error
	ApplyChangeHelper(aggregate AggregateHandler, event *Event, commit bool) error
	HandleCommand(ctx context.Context, command *BaseCommand) error
	Uncommitted() []*Event
	ClearUncommitted()
	IncrementVersion()
}

// Uncommitted return the events to be saved
func (b *BaseAggregate) Uncommitted() []*Event {
	return b.Changes
}

// ClearUncommited the events
func (b *BaseAggregate) ClearUncommitted() {
	b.Changes = []*Event{}
}

// IncrementVersion ads 1 to the current version
func (b *BaseAggregate) IncrementVersion() {
	b.Version++
}

// ApplyChangeHelper increments the version of an aggregate and apply the change itself
func (b *BaseAggregate) ApplyChangeHelper(aggregate AggregateHandler, event *Event, commit bool) error {
	// increments the version in event and aggregate
	b.IncrementVersion()

	// apply the event itself
	err := aggregate.ApplyChange(event)
	if err != nil {
		return err
	}

	if commit {
		event.Version = b.Version
		b.Changes = append(b.Changes, event)
	}

	return nil
}
