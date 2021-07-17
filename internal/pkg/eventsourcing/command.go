package eventsourcing

import (
	"github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
)

// Command contains the methods to retrieve basic info about it
type Command interface {
	GetType() string
	GetAggregateID() string
	GetAggregateType() string
	GetVersion() int
}

// BaseCommand contains the basic info
// that all commands should have
type BaseCommand struct {
	v1.BaseCommand
}

// GetType returns the command type
func (b *BaseCommand) GetType() string {
	return b.Type
}

// GetAggregateID returns the command aggregate ID
func (b *BaseCommand) GetAggregateID() string {
	return b.AggregateId
}

// GetAggregateType returns the command aggregate type
func (b *BaseCommand) GetAggregateType() string {
	return b.AggregateType
}

// GetVersion of the command
func (b *BaseCommand) GetVersion() int32 {
	return b.Version
}
