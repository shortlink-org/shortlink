package batch

import (
	"time"
)

// Option is the type for batch options.
type Option[T any] func(*Batch[T])

// WithInterval sets the interval for the batch.
func WithInterval[T any](interval time.Duration) Option[T] {
	return func(b *Batch[T]) {
		b.interval = interval
	}
}

// WithSize sets the size for the batch.
func WithSize[T any](size int) Option[T] {
	return func(b *Batch[T]) {
		b.size = size
	}
}
