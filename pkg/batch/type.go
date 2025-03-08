package batch

import (
	"sync"
	"time"
)

// Batch is a structure for batch processing
type Batch[T any] struct {
	done     chan struct{}
	errChan  chan error
	callback func([]*Item[T]) error
	items    []*Item[T]
	wg       sync.WaitGroup
	interval time.Duration
	size     int
	mu       sync.Mutex
}

// Item represents an item that can be pushed to the batch.
type Item[T any] struct {
	CallbackChannel chan T
	Item            T
}
