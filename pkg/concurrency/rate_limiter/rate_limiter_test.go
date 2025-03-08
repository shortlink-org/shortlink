package rate_limiter_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/shortlink-org/shortlink/pkg/concurrency/rate_limiter"
)

func TestRateLimiter(t *testing.T) {
	sum := int64(0)

	// Use a context with a timeout to prevent test from running indefinitely
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	rl, err := rate_limiter.New(ctx, 100, 5*time.Millisecond)
	require.NoError(t, err)

	var wg errgroup.Group

	for range 10000 {
		wg.Go(func() error {
			errWait := rl.Wait()
			if errWait != nil {
				return errWait
			}

			atomic.AddInt64(&sum, 1)

			return nil
		})
	}

	err = wg.Wait()
	require.NoError(t, err)
	require.Equal(t, int64(10000), atomic.LoadInt64(&sum))
}
