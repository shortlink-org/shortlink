package batch

import (
	"context"
	"sync"
	"time"
)

// Batch is a structure for batch processing
type Batch struct {
	callback func([]*Item) any
	items    []*Item
	interval time.Duration
	mu       sync.Mutex

	ctx        context.Context
	done       chan struct{}
	cancelFunc context.CancelFunc
}

// Item represents an item that can be pushed to the batch.
type Item struct {
	CallbackChannel chan any
	Item            any
}
