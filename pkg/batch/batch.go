/*
Package for work in batch mode
*/
package batch

import (
	"context"
	"time"

	"github.com/spf13/viper"
)

// New creates a new batch with a specified callback function.
func New[T any](ctx context.Context, cb func([]*Item[T]) error, opts ...Option[T]) (*Batch[T], error) {
	viper.SetDefault("BATCH_INTERVAL", "100ms")
	viper.SetDefault("BATCH_SIZE", 100)

	b := &Batch[T]{
		callback: cb,
		interval: viper.GetDuration("BATCH_INTERVAL"),
		size:     viper.GetInt("BATCH_SIZE"),
		ctx:      ctx,
	}

	// Apply options
	for _, opt := range opts {
		opt(b)
	}

	go b.run()

	return b, nil
}

// Push adds an item to the batch.
func (b *Batch[T]) Push(item T) chan T {
	newItem := &Item[T]{
		CallbackChannel: make(chan T, 1),
		Item:            item,
	}

	select {
	case <-b.ctx.Done():
		close(newItem.CallbackChannel)
		return newItem.CallbackChannel
	default:
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
func (b *Batch[T]) run() {
	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()

	for {
		select {
		case <-b.ctx.Done():
			b.flushItems()
			b.wg.Wait()
			b.closePendingChannels()

			return
		case <-ticker.C:
			b.flushItems()
		}
	}
}

func (b *Batch[T]) closePendingChannels() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, item := range b.items {
		close(item.CallbackChannel)
	}
}

// flushItems flushes all items to the callback function.
func (b *Batch[T]) flushItems() {
	b.mu.Lock()
	items := b.items
	b.items = nil
	b.mu.Unlock()

	if len(items) == 0 {
		return
	}

	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		if err := b.callback(items); err != nil {
			// Handle error if necessary
		}
	}()
}
