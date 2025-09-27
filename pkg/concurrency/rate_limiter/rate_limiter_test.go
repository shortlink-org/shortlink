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
	t.Attr("type", "unit")
	t.Attr("package", "rate_limiter")
	t.Attr("component", "concurrency")

		t.Attr("type", "unit")
		t.Attr("package", "rate_limiter")
		t.Attr("component", "concurrency")
	
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

// TestRateLimiterWithSynctest validates rate limiter token consumption and refill behavior.
// Tests that tokens are properly consumed and that subsequent requests block until
// token refill occurs, ensuring correct rate limiting enforcement.
func TestRateLimiterWithSynctest(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "rate_limiter")
	t.Attr("component", "concurrency")

		t.Attr("type", "unit")
		t.Attr("package", "rate_limiter")
		t.Attr("component", "concurrency")
	
	synctest.Test(t, func(t *testing.T) {, cancel := context.WithCancel(t.Context())
		defer cancel()
		
		// Initialize rate limiter with 2 initial tokens, refilling every 100ms
		rl, err := rate_limiter.New(ctx, 2, 100*time.Millisecond)
		require.NoError(t, err)
		
		// Consume both available tokens - should succeed immediately
		require.NoError(t, rl.Wait())
		require.NoError(t, rl.Wait())
		
		// Third request should block waiting for token refill
		done := make(chan error, 1)
		go func() {
			done <- rl.Wait()
		}()
		
		// Verify the third request is blocked waiting for refill
		synctest.Wait()
		
		// Confirm request is still blocking before refill occurs
		select {
		case <-done:
			t.Fatal("request should be blocked waiting for token refill")
		default:
		}
		
		// Allow refill interval to pass and verify request completes
		// synctest automatically advances time when all goroutines are blocked
		err = <-done
		require.NoError(t, err)
		
		// Cancel to clean up background goroutines
		cancel()
		synctest.Wait()
	})
}

// TestRateLimiterCancellation verifies proper error handling during context cancellation.
// Ensures that blocked rate limiter operations return the appropriate cancellation error
// and that resources are cleaned up correctly when the context is cancelled.
func TestRateLimiterCancellation(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "rate_limiter")
	t.Attr("component", "concurrency")

		t.Attr("type", "unit")
		t.Attr("package", "rate_limiter")
		t.Attr("component", "concurrency")
	
	synctest.Test(t, func(t *testing.T) {, cancel := context.WithCancel(t.Context())
		
		// Create rate limiter with single token and long refill interval
		rl, err := rate_limiter.New(ctx, 1, 1*time.Second)
		require.NoError(t, err)
		
		// Consume the available token
		require.NoError(t, rl.Wait())
		
		// Launch goroutine that will block waiting for token refill
		done := make(chan error, 1)
		go func() {
			done <- rl.Wait()
		}()
		
		// Verify the goroutine is blocked waiting for token
		synctest.Wait()
		
		// Cancel context while request is blocked
		cancel()
		
		// Verify blocked request returns appropriate cancellation error
		err = <-done
		require.Equal(t, rate_limiter.ErrRateLimiterCanceled, err)
		
		// Allow cleanup operations to complete
		synctest.Wait()
	})
}

// TestSimpleRateLimiterWithSynctest validates basic rate limiting functionality using
// a simplified token bucket implementation. Tests token consumption and refill behavior
// in a controlled environment to ensure rate limiting mechanics work correctly.
func TestSimpleRateLimiterWithSynctest(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "rate_limiter")
	t.Attr("component", "concurrency")

		t.Attr("type", "unit")
		t.Attr("package", "rate_limiter")
		t.Attr("component", "concurrency")
	
	synctest.Test(t, func(t *testing.T) {
		// Create a simple rate limiter scenario without persistent background goroutines
		limiter := make(chan struct{}, 2)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop(), cancel := context.WithCancel(t.Context())
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
