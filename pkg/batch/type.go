package batch

import (
	"context"
	"sync"
	"time"
)

// Batch is a structure for batch processing
type Batch[T any] struct {
	mu sync.Mutex

	callback func([]*Item[T]) error
	items    []*Item[T]

	interval time.Duration
	size     int

	wg  sync.WaitGroup
	ctx context.Context
}

// Item represents an item that can be pushed to the batch.
type Item[T any] struct {
	CallbackChannel chan T
	Item            T
}
