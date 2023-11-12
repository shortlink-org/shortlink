package rate_limiter

import (
	"context"
	"sync"
	"time"
)

type RateLimiter struct {
	mu  sync.Mutex
	ctx context.Context

	limit   int64
	limiter chan struct{}
	ticker  *time.Ticker
}

func New(ctx context.Context, limit int64, interval time.Duration) (*RateLimiter, error) {
	ticker := time.NewTicker(interval)

	rl := &RateLimiter{
		ctx:     ctx,
		limit:   limit,
		limiter: make(chan struct{}, limit),
		ticker:  ticker,
	}

	go rl.refill()

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		rl.ticker.Stop()
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
	case <-r.ctx.Done():
		return r.ctx.Err()
	}
}

// refill refills tokens periodically
func (r *RateLimiter) refill() {
	for {
		select {
		case <-r.ticker.C:
			for i := int64(0); i < r.limit; i++ {
				select {
				case r.limiter <- struct{}{}:
				default:
				}
			}
		case <-r.ctx.Done():
			r.ticker.Stop()
			return
		}
	}
}
