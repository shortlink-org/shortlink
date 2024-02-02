package batch

import (
	"time"

	"github.com/sasha-s/go-deadlock"
)

// Batch is a structure for batch processing
type Batch struct {
	mu deadlock.Mutex

	callback func([]*Item) any
	items    []*Item

	interval time.Duration
	size     int
}

// Item represents an item that can be pushed to the batch.
type Item struct {
	CallbackChannel chan any
	Item            any
}
