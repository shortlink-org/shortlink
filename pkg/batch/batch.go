/*
Package for work in batch mode
*/
package batch

import (
	"context"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// New creates a new batch with a specified callback function.
func New[T any](ctx context.Context, callback func([]*Item[T]) error, opts ...Option[T]) (*Batch[T], error) {
	viper.SetDefault("BATCH_INTERVAL", "100ms")
	viper.SetDefault("BATCH_SIZE", 100)

	batch := &Batch[T]{
		ctx: ctx,
		mu:  sync.Mutex{},
		wg:  sync.WaitGroup{},

		callback: callback,
		items:    []*Item[T]{},
		interval: viper.GetDuration("BATCH_INTERVAL"),
		size:     viper.GetInt("BATCH_SIZE"),
	}

	// Apply options
	for _, opt := range opts {
		opt(batch)
	}

	go batch.run()

	return batch, nil
}

// Push adds an item to the batch.
func (batch *Batch[T]) Push(item T) chan T {
	newItem := &Item[T]{
		CallbackChannel: make(chan T, 1),
		Item:            item,
	}

	select {
	case <-batch.ctx.Done():
		close(newItem.CallbackChannel)
		return newItem.CallbackChannel
	default:
	}

	batch.mu.Lock()
	batch.items = append(batch.items, newItem)
	shouldFlush := len(batch.items) >= batch.size
	batch.mu.Unlock()

	// If the batch is full, flush it
	if shouldFlush {
		go batch.flushItems()
	}

	return newItem.CallbackChannel
}

// run starts a loop flushing at the specified interval.
func (batch *Batch[T]) run() {
	ticker := time.NewTicker(batch.interval)
	defer ticker.Stop()

	for {
		select {
		case <-batch.ctx.Done():
			batch.flushItems()
			batch.wg.Wait()
			batch.closePendingChannels()

			return
		case <-ticker.C:
			batch.flushItems()
		}
	}
}

func (batch *Batch[T]) closePendingChannels() {
	batch.mu.Lock()
	defer batch.mu.Unlock()

	for _, item := range batch.items {
		close(item.CallbackChannel)
	}
}

// flushItems flushes all items to the callback function.
func (batch *Batch[T]) flushItems() {
	batch.mu.Lock()
	items := batch.items
	batch.items = nil
	batch.mu.Unlock()

	if len(items) == 0 {
		return
	}

	batch.wg.Add(1)

	go func() {
		defer batch.wg.Done()

		if err := batch.callback(items); err != nil {
			// Handle error if necessary
		}
	}()
}
