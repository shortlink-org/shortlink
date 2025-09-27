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

// TestBasicSynctest demonstrates the fundamental synctest concepts
func TestBasicSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Test simple time-based operations
		start := time.Now()
		
		// This sleep happens instantly in synctest
		time.Sleep(1 * time.Second)
		
		elapsed := time.Since(start)
		// In synctest, exactly 1 second has passed (fake time)
		require.Equal(t, 1*time.Second, elapsed)
	})
}

// TestChannelOperationsWithSynctest shows how synctest handles channel operations
func TestChannelOperationsWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan string, 1)
		var received []string

		// Producer goroutine
		go func() {
			time.Sleep(10 * time.Millisecond)
			ch <- "message1"
			
			time.Sleep(20 * time.Millisecond)
			ch <- "message2"
			
			close(ch)
		}()

		// Consumer goroutine
		go func() {
			for msg := range ch {
				received = append(received, msg)
			}
		}()

		// Wait for all operations to complete
		synctest.Wait()

		require.Equal(t, []string{"message1", "message2"}, received)
	})
}

// TestTimerWithSynctest demonstrates testing timer-based operations
func TestTimerWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var executed bool
		
		timer := time.AfterFunc(100*time.Millisecond, func() {
			executed = true
		})
		defer timer.Stop()

		// Timer hasn't fired yet
		require.False(t, executed)

		// Wait for timer to fire (happens instantly in synctest)
		synctest.Wait()

		require.True(t, executed)
	})
}

// TestContextTimeoutWithSynctest shows context timeout testing
func TestContextTimeoutWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		done := make(chan struct{})
		go func() {
			// This operation takes longer than the timeout
			time.Sleep(100 * time.Millisecond)
			close(done)
		}()

		select {
		case <-done:
			t.Fatal("operation should have been cancelled by timeout")
		case <-ctx.Done():
			// Expected: context should timeout first
			require.Equal(t, context.DeadlineExceeded, ctx.Err())
		}
	})
}

// TestMultipleGoroutinesWithSynctest demonstrates coordinating multiple goroutines
func TestMultipleGoroutinesWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const numGoroutines = 10
		var counter int64
		var wg sync.WaitGroup

		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				
				// Each goroutine waits a different amount of time
				time.Sleep(time.Duration(id*10) * time.Millisecond)
				atomic.AddInt64(&counter, 1)
			}(i)
		}

		wg.Wait()

		require.Equal(t, int64(numGoroutines), atomic.LoadInt64(&counter))
	})
}

// TestTickerWithSynctest demonstrates testing ticker-based operations
func TestTickerWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var tickCount int64
		ticker := time.NewTicker(25 * time.Millisecond)
		defer ticker.Stop()

		done := make(chan struct{})
		go func() {
			defer close(done)
			for i := 0; i < 5; i++ {
				<-ticker.C
				atomic.AddInt64(&tickCount, 1)
			}
		}()

		// Wait for all ticks to complete
		<-done

		require.Equal(t, int64(5), atomic.LoadInt64(&tickCount))
	})
}

// TestRaceConditionWithSynctest shows deterministic race condition testing
func TestRaceConditionWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var sharedValue int64
		var mu sync.Mutex
		const numGoroutines = 50

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				
				// Add some variability with sleep
				time.Sleep(time.Duration(id) * time.Microsecond)
				
				mu.Lock()
				current := atomic.LoadInt64(&sharedValue)
				// Simulate work in critical section
				time.Sleep(1 * time.Microsecond)
				atomic.StoreInt64(&sharedValue, current+1)
				mu.Unlock()
			}(i)
		}

		wg.Wait()

		// Should be deterministic in synctest
		require.Equal(t, int64(numGoroutines), atomic.LoadInt64(&sharedValue))
	})
}

// TestSelectWithTimeoutSynctest demonstrates select statement testing
func TestSelectWithTimeoutSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan string)
		
		go func() {
			time.Sleep(30 * time.Millisecond)
			ch <- "data"
		}()

		// Test select with timeout
		select {
		case data := <-ch:
			require.Equal(t, "data", data)
		case <-time.After(100 * time.Millisecond):
			t.Fatal("should not timeout")
		}
	})
}

// TestDeadlinePropagationWithSynctest shows context deadline propagation
func TestDeadlinePropagationWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		var operationCompleted bool
		
		done := make(chan error, 1)
		go func() {
			// Simulate an operation that respects context
			select {
			case <-time.After(50 * time.Millisecond):
				operationCompleted = true
				done <- nil
			case <-ctx.Done():
				done <- ctx.Err()
			}
		}()

		err := <-done
		require.NoError(t, err)
		require.True(t, operationCompleted)
	})
}

// TestConcurrentMapAccess shows safe concurrent map operations
func TestConcurrentMapAccess(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		data := make(map[string]int)
		var mu sync.RWMutex
		var wg sync.WaitGroup

		// Writers
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				time.Sleep(time.Duration(id) * time.Millisecond)
				
				mu.Lock()
				data[string(rune('a'+id))] = id
				mu.Unlock()
			}(i)
		}

		// Readers
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				time.Sleep(time.Duration(id*2) * time.Millisecond)
				
				mu.RLock()
				_ = len(data) // Read operation
				mu.RUnlock()
			}(i)
		}

		wg.Wait()

		mu.RLock()
		require.Equal(t, 10, len(data))
		mu.RUnlock()
	})
}

// TestBoundedResourcePoolWithSynctest demonstrates testing resource pools
func TestBoundedResourcePoolWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const poolSize = 3
		pool := make(chan struct{}, poolSize)
		
		// Fill the pool
		for i := 0; i < poolSize; i++ {
			pool <- struct{}{}
		}

		var activeOperations int64
		var completedOperations int64
		var wg sync.WaitGroup

		// Launch more goroutines than pool capacity
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				
				// Acquire resource
				<-pool
				atomic.AddInt64(&activeOperations, 1)
				
				// Use resource
				time.Sleep(time.Duration(id*10) * time.Millisecond)
				
				// Release resource
				atomic.AddInt64(&activeOperations, -1)
				atomic.AddInt64(&completedOperations, 1)
				pool <- struct{}{}
			}(i)
		}

		wg.Wait()

		require.Equal(t, int64(0), atomic.LoadInt64(&activeOperations))
		require.Equal(t, int64(10), atomic.LoadInt64(&completedOperations))
	})
}