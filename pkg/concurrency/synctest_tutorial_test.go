package concurrency_test

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

// TestSynctest1_BasicTimeControl demonstrates basic time control in synctest
func TestSynctest1_BasicTimeControl(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		
		// This sleep executes instantly in synctest, but fake time still advances
		time.Sleep(1 * time.Second)
		
		elapsed := time.Since(start)
		require.Equal(t, 1*time.Second, elapsed)
	})
}

// TestSynctest2_ChannelCommunication shows channel operations with time
func TestSynctest2_ChannelCommunication(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan string)
		var received string

		go func() {
			time.Sleep(100 * time.Millisecond)
			ch <- "hello"
		}()

		received = <-ch
		require.Equal(t, "hello", received)
	})
}

// TestSynctest3_Timer demonstrates timer operations
func TestSynctest3_Timer(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		fired := make(chan bool, 1)
		
		timer := time.NewTimer(50 * time.Millisecond)
		defer timer.Stop()

		go func() {
			<-timer.C
			fired <- true
		}()

		result := <-fired
		require.True(t, result)
	})
}

// TestSynctest4_ContextTimeout shows context timeout behavior
func TestSynctest4_ContextTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		done := make(chan bool)
		go func() {
			select {
			case <-time.After(200 * time.Millisecond):
				done <- false // Operation completed
			case <-ctx.Done():
				done <- true // Context timed out
			}
		}()

		timedOut := <-done
		require.True(t, timedOut, "context should have timed out")
	})
}

// TestSynctest5_MultipleGoroutines demonstrates coordinating multiple goroutines
func TestSynctest5_MultipleGoroutines(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var counter int64
		var wg sync.WaitGroup
		const numGoroutines = 10

		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				
				// Each goroutine waits a different amount
				time.Sleep(time.Duration(id*10) * time.Millisecond)
				atomic.AddInt64(&counter, 1)
			}(i)
		}

		wg.Wait()
		require.Equal(t, int64(numGoroutines), atomic.LoadInt64(&counter))
	})
}

// TestSynctest6_Ticker demonstrates ticker operations
func TestSynctest6_Ticker(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ticks int64
		ticker := time.NewTicker(25 * time.Millisecond)
		
		done := make(chan struct{})
		go func() {
			defer close(done)
			for i := 0; i < 4; i++ {
				<-ticker.C
				atomic.AddInt64(&ticks, 1)
			}
			ticker.Stop()
		}()

		<-done
		require.Equal(t, int64(4), atomic.LoadInt64(&ticks))
	})
}

// TestSynctest7_Select demonstrates select statement with timeouts
func TestSynctest7_Select(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan string)
		
		go func() {
			time.Sleep(30 * time.Millisecond)
			ch <- "data"
		}()

		select {
		case data := <-ch:
			require.Equal(t, "data", data)
		case <-time.After(100 * time.Millisecond):
			t.Fatal("should not timeout")
		}
	})
}

// TestSynctest8_ResourcePool demonstrates bounded resource testing
func TestSynctest8_ResourcePool(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const poolSize = 3
		pool := make(chan struct{}, poolSize)
		
		// Fill the pool
		for i := 0; i < poolSize; i++ {
			pool <- struct{}{}
		}

		var completed int64
		var wg sync.WaitGroup

		// More workers than pool capacity
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				
				// Acquire resource
				<-pool
				
				// Use resource
				time.Sleep(time.Duration(id*5) * time.Millisecond)
				atomic.AddInt64(&completed, 1)
				
				// Release resource
				pool <- struct{}{}
			}(i)
		}

		wg.Wait()
		require.Equal(t, int64(10), atomic.LoadInt64(&completed))
	})
}

// TestSynctest9_ProducerConsumer demonstrates producer-consumer pattern
func TestSynctest9_ProducerConsumer(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		jobs := make(chan int, 5)
		results := make(chan int, 5)

		// Producer
		go func() {
			for i := 1; i <= 5; i++ {
				time.Sleep(10 * time.Millisecond)
				jobs <- i
			}
			close(jobs)
		}()

		// Consumer
		go func() {
			for job := range jobs {
				// Process job
				time.Sleep(5 * time.Millisecond)
				results <- job * 2
			}
			close(results)
		}()

		// Collect results
		var allResults []int
		for result := range results {
			allResults = append(allResults, result)
		}

		require.Equal(t, []int{2, 4, 6, 8, 10}, allResults)
	})
}

// TestSynctest10_RateLimitingPattern demonstrates a simple rate limiting pattern
func TestSynctest10_RateLimitingPattern(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Simple rate limiter: 2 tokens, refill every 100ms
		tokens := make(chan struct{}, 2)
		tokens <- struct{}{}
		tokens <- struct{}{}

		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Token refiller
		go func() {
			for {
				select {
				case <-ticker.C:
					select {
					case tokens <- struct{}{}:
					default:
					}
				case <-ctx.Done():
					return
				}
			}
		}()

		// Use first two tokens immediately
		<-tokens
		<-tokens

		// Third request should wait for refill
		done := make(chan struct{})
		go func() {
			<-tokens
			close(done)
		}()

		// Should complete after ticker fires
		<-done
		
		cancel() // Clean up
		synctest.Wait()
	})
}

// TestSynctest11_TimeComparison shows the performance benefit
func TestSynctest11_TimeComparison(t *testing.T) {
	// Traditional test (uncomment to see real timing)
	/*
	t.Run("Traditional", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < 10; i++ {
			time.Sleep(10 * time.Millisecond)
		}
		elapsed := time.Since(start)
		t.Logf("Traditional took: %v", elapsed) // ~100ms
	})
	*/

	// Synctest equivalent
	t.Run("Synctest", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			fakeStart := time.Now()
			
			for i := 0; i < 10; i++ {
				time.Sleep(10 * time.Millisecond)
			}
			
			fakeElapsed := time.Since(fakeStart)
			
			// Fake time should be exactly 100ms
			t.Logf("Fake time: %v (executed instantly)", fakeElapsed)
			require.Equal(t, 100*time.Millisecond, fakeElapsed)
		})
	})
}