package rate_limiter

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrRateLimiterCanceled = errors.New("rate limiter context canceled")

type RateLimiter struct {
	mu   sync.Mutex
	done chan struct{}

	limiter chan struct{}
	ticker  *time.Ticker
	limit   int64
}

func New(ctx context.Context, limit int64, interval time.Duration) (*RateLimiter, error) {
	ticker := time.NewTicker(interval)
	done := make(chan struct{})

	rl := &RateLimiter{
		limiter: make(chan struct{}, limit),
		ticker:  ticker,
		limit:   limit,
		done:    done,
	}

	go rl.refill()

	// Graceful shutdown: when the context is canceled, signal via the done channel.
	go func() {
		<-ctx.Done()
		close(done)
		ticker.Stop()
		close(rl.limiter)
	}()

	return rl, nil
}

func (r *RateLimiter) Wait() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-r.limiter:
		return nil
	case <-r.done:
		return ErrRateLimiterCanceled
	}
}

// refill refills tokens periodically
func (r *RateLimiter) refill() {
	for {
		select {
		case <-r.ticker.C:
			// Refill up to 'limit' tokens.
			for range r.limit {
				select {
				case r.limiter <- struct{}{}:
				default:
				}
			}
		case <-r.done:
			r.ticker.Stop()
			return
		}
	}
}
