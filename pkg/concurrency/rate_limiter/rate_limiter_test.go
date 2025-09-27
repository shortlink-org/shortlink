package rate_limiter_test

import (
	"context"
	"sync/atomic"
	"testing"
	"testing/synctest"
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

// TestRateLimiterWithSynctest demonstrates testing rate limiter with controlled time
func TestRateLimiterWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		
		// Create rate limiter with 2 tokens, refilling every 100ms
		rl, err := rate_limiter.New(ctx, 2, 100*time.Millisecond)
		require.NoError(t, err)
		
		// First two requests should succeed immediately
		require.NoError(t, rl.Wait())
		require.NoError(t, rl.Wait())
		
		// The third request should block until the next refill
		done := make(chan error, 1)
		go func() {
			done <- rl.Wait()
		}()
		
		// Ensure the goroutine is blocked
		synctest.Wait()
		
		// No completion yet since we haven't advanced time
		select {
		case <-done:
			t.Fatal("request should be blocked")
		default:
		}
		
		// After refill time passes, the request should complete
		// In synctest, time advances automatically when all goroutines are blocked
		err = <-done
		require.NoError(t, err)
		
		// Cancel to clean up background goroutines
		cancel()
		synctest.Wait()
	})
}

// TestRateLimiterCancellation tests proper cleanup when context is cancelled
func TestRateLimiterCancellation(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		
		rl, err := rate_limiter.New(ctx, 1, 1*time.Second)
		require.NoError(t, err)
		
		// Use the token
		require.NoError(t, rl.Wait())
		
		// Start a goroutine that will block waiting for the next token
		done := make(chan error, 1)
		go func() {
			done <- rl.Wait()
		}()
		
		// Ensure the goroutine is blocked
		synctest.Wait()
		
		// Cancel the context
		cancel()
		
		// The blocked wait should return with cancellation error
		err = <-done
		require.Equal(t, rate_limiter.ErrRateLimiterCanceled, err)
		
		// Wait for cleanup
		synctest.Wait()
	})
}

// TestSimpleRateLimiterWithSynctest demonstrates basic rate limiting in controlled environment
func TestSimpleRateLimiterWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create a simple rate limiter scenario without persistent background goroutines
		limiter := make(chan struct{}, 2)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Fill initial tokens
		limiter <- struct{}{}
		limiter <- struct{}{}

		// Background refiller
		go func() {
			for {
				select {
				case <-ticker.C:
					select {
					case limiter <- struct{}{}:
					default:
					}
				case <-ctx.Done():
					return
				}
			}
		}()

		// Test immediate consumption
		<-limiter // First token
		<-limiter // Second token

		// Third request should block until refill
		done := make(chan struct{})
		go func() {
			<-limiter // This should block and wait for refill
			close(done)
		}()

		// Wait for refill (100ms passes instantly in synctest)
		synctest.Wait()

		// Should complete after time advancement
		select {
		case <-done:
			// Success
		default:
			t.Fatal("should have received token after refill")
		}

		cancel()
		synctest.Wait()
	})
}
