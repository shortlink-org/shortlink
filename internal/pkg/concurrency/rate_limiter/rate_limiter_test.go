package rate_limiter

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestRateLimiter(t *testing.T) {
	sum := int64(0)

	// Use a context with a timeout to prevent test from running indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rl, err := New(ctx, 100, 5*time.Millisecond)
	require.NoError(t, err)

	defer rl.Close()

	var wg errgroup.Group

	for i := 0; i < 10000; i++ {
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
