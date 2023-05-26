/*
Package for work in batch mode
*/
package batch

import (
	"context"
	"time"
)

// New creates a new batch Config with a specified callback function.
func New(ctx context.Context, cb func([]*Item) interface{}, opts ...Option) (*Batch, error) {
	b := &Batch{
		callback: cb,
		interval: time.Millisecond * 100, // nolint:gomnd
	}

	// Apply options
	for _, opt := range opts {
		opt(b)
	}

	go b.run(ctx)
	return b, nil
}

// Push adds an item to the batch.
func (b *Batch) Push(item interface{}) chan interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()
	newItem := &Item{
		CallbackChannel: make(chan interface{}),
		Item:            item,
	}
	b.items = append(b.items, newItem)
	return newItem.CallbackChannel
}

// run starts a loop flushing at the specified interval.
func (b *Batch) run(ctx context.Context) {
	ticker := time.NewTicker(b.interval)

	for {
		select {
		case <-ctx.Done():
			b.mu.Lock()
			for _, item := range b.items {
				item.CallbackChannel <- struct{}{}
			}
			b.mu.Unlock()

		case <-ticker.C:
			b.mu.Lock()
			if len(b.items) > 0 {
				b.callback(b.items)
				b.items = []*Item{}
			}
			b.mu.Unlock()
		}
	}
}
