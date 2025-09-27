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

const defaultErrChanBuffer = 10

// New creates a new batch with a specified callback function.
func New[T any](ctx context.Context, callback func([]*Item[T]) error, opts ...Option[T]) (*Batch[T], <-chan error) {
	viper.SetDefault("BATCH_INTERVAL", "100ms")
	viper.SetDefault("BATCH_SIZE", 100)
	viper.SetDefault("BATCH_ERROR_BUFFER", defaultErrChanBuffer)

	batch := &Batch[T]{
		mu: sync.Mutex{},
		wg: sync.WaitGroup{},

		callback: callback,
		items:    []*Item[T]{},
		interval: viper.GetDuration("BATCH_INTERVAL"),
		size:     viper.GetInt("BATCH_SIZE"),
		// Instead of storing ctx, use a done channel.
		done: make(chan struct{}),
		// Buffered error channel to report errors from callback.
		errChan: make(chan error, viper.GetInt("BATCH_ERROR_BUFFER")),
	}

	// Apply options
	for _, opt := range opts {
		opt(batch)
	}

	// Launch a goroutine to monitor the passed context.
	go func() {
		<-ctx.Done()
		close(batch.done)
	}()

	go batch.run()

	return batch, batch.errChan
}

// Push adds an item to the batch.
func (batch *Batch[T]) Push(item T) chan T {
	newItem := &Item[T]{
		CallbackChannel: make(chan T, 1),
		Item:            item,
	}

	// Check for cancellation using the done channel.
	select {
	case <-batch.done:
		close(newItem.CallbackChannel)
		return newItem.CallbackChannel
	default:
	}

	batch.mu.Lock()
	batch.items = append(batch.items, newItem)
	shouldFlush := len(batch.items) >= batch.size
	batch.mu.Unlock()

	// If the batch is full, flush it.
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
		case <-batch.done:
			batch.flushItems()
			batch.wg.Wait()
			batch.closePendingChannels()
			close(batch.errChan)

			return
		case <-ticker.C:
			batch.flushItems()
		}
	}
}

// closePendingChannels closes all pending channels.
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

	// Check if cancellation has already occurred while still holding the lock.
	doneClosed := false

	select {
	case <-batch.done:
		doneClosed = true
	default:
	}

	batch.mu.Unlock()

	if len(items) == 0 {
		return
	}

	if doneClosed {
		for _, item := range items {
			close(item.CallbackChannel)
		}

		return
	}

	batch.wg.Go(func() {
		// Check cancellation again before proceeding.
		select {
		case <-batch.done:
			for _, item := range items {
				close(item.CallbackChannel)
			}
		default:
			if err := batch.callback(items); err != nil {
				select {
				case batch.errChan <- err:
				default:
				}
			}
		}
	})
}
