/*
Package for work in batch mode
*/
package batch

import (
	"context"
	"time"

	"github.com/spf13/viper"
)

// New creates a new batch Config with a specified callback function.
func New(ctx context.Context, cb func([]*Item) any, opts ...Option) (*Batch, error) {
	viper.SetDefault("BATCH_INTERVAL", "100ms")
	viper.SetDefault("BATCH_SIZE", 100)

	b := &Batch{
		callback: cb,

		interval: viper.GetDuration("BATCH_INTERVAL"),
		size:     viper.GetInt("BATCH_SIZE"),
	}

	// Apply options
	for _, opt := range opts {
		opt(b)
	}

	go b.run(ctx)

	return b, nil
}

// Push adds an item to the batch.
func (b *Batch) Push(item any) chan any {
	newItem := &Item{
		CallbackChannel: make(chan any),
		Item:            item,
	}

	b.mu.Lock()
	b.items = append(b.items, newItem)
	shouldFlush := len(b.items) >= b.size
	b.mu.Unlock()

	// If the batch is full, flush it
	if shouldFlush {
		go b.flushItems()
	}

	return newItem.CallbackChannel
}

// run starts a loop flushing at the specified interval.
func (b *Batch) run(ctx context.Context) {
	ticker := time.NewTicker(b.interval)
	defer ticker.Stop() // Stop the ticker

	for {
		select {
		case <-ctx.Done():
			// Clear all items
			b.clearItems()

			return
		case <-ticker.C:
			b.flushItems()
		}
	}
}

func (b *Batch) clearItems() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, item := range b.items {
		close(item.CallbackChannel)
	}
}

// flushItems - flushes all items to the callback function.
func (b *Batch) flushItems() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.callback(b.items)
	b.items = []*Item{}
}
