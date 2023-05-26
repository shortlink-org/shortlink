package batch

import (
	"time"
)

// Option is the type for batch options.
type Option func(*Batch)

// WithInterval sets the interval for the batch.
func WithInterval(interval time.Duration) Option {
	return func(b *Batch) {
		b.interval = interval
	}
}
