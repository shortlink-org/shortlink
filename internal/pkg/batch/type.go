package batch

import (
	"context"
	"time"

	"github.com/sasha-s/go-deadlock"
)

// Batch is a structure for batch processing
type Batch struct {
	callback func([]*Item) any
	items    []*Item
	interval time.Duration
	mu       deadlock.Mutex

	ctx        context.Context
	done       chan struct{}
	cancelFunc context.CancelFunc
}

// Item represents an item that can be pushed to the batch.
type Item struct {
	CallbackChannel chan any
	Item            any
}
