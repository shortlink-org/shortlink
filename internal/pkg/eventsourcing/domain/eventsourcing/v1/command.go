package v1

import (
	"context"
)

// CommandHandle defines the contract to handle commands
type CommandHandle interface {
	Handle(ctx context.Context, in any) (*BaseCommand, error)
	PublishEvents(ctx context.Context, events []*Event) error
}
