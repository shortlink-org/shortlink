package batch

import (
	"context"
	"sync"
	"time"
)

// Config
type Config struct {
	ctx      context.Context
	callback func([]*Item) interface{}
	items    []*Item
	interval time.Duration
	mu       sync.Mutex
}

// Item represents an item that can be pushed to the batch.
type Item struct {
	CallbackChannel chan interface{}
	Item            interface{}
}
