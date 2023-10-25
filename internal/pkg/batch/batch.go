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

	ctx, cancelFunc := context.WithCancel(ctx)

	b := &Batch{
		callback: cb,
		interval: viper.GetDuration("BATCH_INTERVAL"),

		ctx:        ctx,
		done:       make(chan struct{}),
		cancelFunc: cancelFunc,
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
	b.mu.Lock()
	defer b.mu.Unlock()
	newItem := &Item{
		CallbackChannel: make(chan any),
		Item:            item,
	}
	b.items = append(b.items, newItem)

	return newItem.CallbackChannel
}

// run starts a loop flushing at the specified interval.
func (b *Batch) run(ctx context.Context) {
	ticker := time.NewTicker(b.interval)
	defer ticker.Stop() // Stop the ticker

	for {
		select {
		case <-ctx.Done():
			b.clearItems()
			close(b.done)

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
		item.CallbackChannel <- struct{}{}
	}
}

func (b *Batch) flushItems() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.items) > 0 {
		// Check if the context is done before invoking the callback
		select {
		case <-b.ctx.Done():
			return
		default:
			b.callback(b.items)
			b.items = []*Item{}
		}
	}
}

// Stop cancels the Batch's context, effectively stopping the run goroutine.
func (b *Batch) Stop() {
	b.cancelFunc()
}
