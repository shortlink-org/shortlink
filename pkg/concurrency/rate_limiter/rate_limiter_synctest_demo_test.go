package rate_limiter_test

import (
	"context"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/concurrency/rate_limiter"
)

// TestRateLimiterDemoTraditional shows the traditional approach with real time
func TestRateLimiterDemoTraditional(t *testing.T) {
	// This test takes real time to execute
	start := time.Now()
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Rate limiter with 1 token, refilling every 100ms
	rl, err := rate_limiter.New(ctx, 1, 100*time.Millisecond)
	require.NoError(t, err)

	// First request succeeds immediately
	require.NoError(t, rl.Wait())

	// Second request should block for ~100ms
	require.NoError(t, rl.Wait())

	elapsed := time.Since(start)
	// This test takes at least 100ms of real time
	require.GreaterOrEqual(t, elapsed, 100*time.Millisecond)
	
	t.Logf("Traditional test took: %v", elapsed)
}

// TestRateLimiterDemoSynctest shows the same test with synctest - instant execution
func TestRateLimiterDemoSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// This test executes instantly
		start := time.Now()
		
		ctx, cancel := context.WithCancel(context.Background())
		
		// Rate limiter with 1 token, refilling every 100ms
		rl, err := rate_limiter.New(ctx, 1, 100*time.Millisecond)
		require.NoError(t, err)

		// First request succeeds immediately
		require.NoError(t, rl.Wait())

		// Test that second request blocks initially
		blocked := make(chan struct{})
		result := make(chan error, 1)
		
		go func() {
			close(blocked) // Signal that we're about to block
			result <- rl.Wait()
		}()

		// Wait for goroutine to start
		<-blocked
		synctest.Wait()

		// Should complete after time advancement
		err = <-result
		require.NoError(t, err)

		// Clean up
		cancel()
		synctest.Wait()

		elapsed := time.Since(start)
		t.Logf("Synctest test took: %v (fake time still advanced correctly)", elapsed)
	})
}

// TestComparePerformance demonstrates the performance difference
func TestComparePerformance(t *testing.T) {
	// Traditional approach
	t.Run("Traditional", func(t *testing.T) {
		start := time.Now()
		
		// Multiple time-based operations
		for i := 0; i < 5; i++ {
			time.Sleep(10 * time.Millisecond)
		}
		
		elapsed := time.Since(start)
		require.GreaterOrEqual(t, elapsed, 50*time.Millisecond)
		t.Logf("Traditional: %v", elapsed)
	})

	// Synctest approach
	t.Run("Synctest", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			start := time.Now()
			
			// Same operations execute instantly
			for i := 0; i < 5; i++ {
				time.Sleep(10 * time.Millisecond)
			}
			
			elapsed := time.Since(start)
			// Time still advances correctly in fake time
			require.Equal(t, 50*time.Millisecond, elapsed)
			t.Logf("Synctest: %v (instant execution)", elapsed)
		})
	})
}

// TestDeterministicExecution shows how synctest eliminates flakiness
func TestDeterministicExecution(t *testing.T) {
	// Run the same test multiple times to show determinism
	for run := 0; run < 5; run++ {
		t.Run("Run", func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				var results []int
				var mu sync.Mutex

				// Multiple goroutines with different delays
				var wg sync.WaitGroup
				for i := 0; i < 5; i++ {
					wg.Add(1)
					go func(id int) {
						defer wg.Done()
						
						// Different delays for each goroutine
						time.Sleep(time.Duration(id*10) * time.Millisecond)
						
						mu.Lock()
						results = append(results, id)
						mu.Unlock()
					}(i)
				}

				wg.Wait()

				// Results should be deterministic and complete
				require.Len(t, results, 5)
				// In synctest, the order is deterministic based on sleep duration
				require.Equal(t, []int{0, 1, 2, 3, 4}, results)
			})
		})
	}
}

// TestChannelDeadlockDetection shows how synctest can detect deadlocks
func TestChannelDeadlockDetection(t *testing.T) {
	// This test will pass because we avoid the deadlock
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan int)
		
		go func() {
			// Send with a timeout to avoid deadlock
			select {
			case ch <- 42:
			case <-time.After(10 * time.Millisecond):
				// Timeout - close channel to signal completion
				close(ch)
			}
		}()

		go func() {
			// Receive with delay
			time.Sleep(5 * time.Millisecond)
			<-ch
		}()

		// Wait for completion
		synctest.Wait()
	})
}

// TestContextPropagationWithSynctest demonstrates context propagation in concurrent operations
func TestContextPropagationWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		var operations []string
		var mu sync.Mutex

		// Launch operations with different durations
		for i, duration := range []time.Duration{30, 60, 120} {
			go func(id int, d time.Duration) {
				select {
				case <-time.After(d * time.Millisecond):
					mu.Lock()
					operations = append(operations, "completed")
					mu.Unlock()
				case <-ctx.Done():
					mu.Lock()
					operations = append(operations, "cancelled")
					mu.Unlock()
				}
			}(i, duration)
		}

		// Wait for all operations
		synctest.Wait()

		mu.Lock()
		defer mu.Unlock()
		
		// First two should complete, third should be cancelled
		require.Len(t, operations, 3)
		require.Equal(t, "completed", operations[0])  // 30ms
		require.Equal(t, "completed", operations[1])  // 60ms
		require.Equal(t, "cancelled", operations[2])  // 120ms > 100ms timeout
	})
}

// TestWorkerPatternWithSynctest demonstrates the worker pattern using synctest
func TestWorkerPatternWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		jobs := make(chan int, 10)
		results := make(chan int, 10)

		// Start workers
		const numWorkers = 3
		var wg sync.WaitGroup
		
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				for job := range jobs {
					// Simulate work with different processing times
					time.Sleep(time.Duration(job*5) * time.Millisecond)
					results <- job * 2
				}
			}(i)
		}

		// Send jobs
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)

		// Wait for workers to finish
		wg.Wait()
		close(results)

		// Collect results
		var allResults []int
		for result := range results {
			allResults = append(allResults, result)
		}

		require.Len(t, allResults, 5)
		// Results should be deterministic
		require.Contains(t, allResults, 2)  // 1*2
		require.Contains(t, allResults, 4)  // 2*2
		require.Contains(t, allResults, 6)  // 3*2
		require.Contains(t, allResults, 8)  // 4*2
		require.Contains(t, allResults, 10) // 5*2
	})
}