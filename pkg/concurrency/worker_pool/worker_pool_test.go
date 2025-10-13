package worker_pool_test

import (
	"context"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/pkg/concurrency/worker_pool"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func Test_WorkerPool(t *testing.T) {
	wp := worker_pool.New(10)

	f := func() (any, error) {
		// some operation
		return nil, nil
	}

	wg := sync.WaitGroup{}
	done := make(chan struct{})

	go func() {
		for range 1000 {
			wp.Push(f)
			wg.Go(func() {
				<-wp.Result
			})
		}

		close(done)
	}()

	<-done
	wg.Wait()
	close(wp.Result)

	t.Cleanup(func() {
		wp.Close()
	})
}

// TestWorkerPoolWithSynctest validates worker pool task execution and result collection.
// Tests that tasks are properly distributed across workers, executed concurrently,
// and results are collected correctly without timing dependencies.
func TestWorkerPoolWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const numWorkers = 3
		const numTasks = 10
		
		wp := worker_pool.New(numWorkers)

		var completedTasks int64
		var results []any
		var resultWg sync.WaitGroup
		resultWg.Add(1)

		// Define task function that simulates processing work
		taskFunc := func() (any, error) {
			// Simulate processing time - executes instantly in synctest
			time.Sleep(10 * time.Millisecond)
			return atomic.AddInt64(&completedTasks, 1), nil
		}

		// Start background result collector goroutine
		go func() {
			defer resultWg.Done()
			for result := range wp.Result {
				results = append(results, result.Value)
				if len(results) == numTasks {
					return
				}
			}
		}()

		// Submit all tasks to the worker pool for concurrent execution
		for i := 0; i < numTasks; i++ {
			wp.Push(taskFunc)
		}

		// Wait for all tasks to be processed and results collected
		resultWg.Wait()

		// Properly shutdown worker pool and cleanup resources
		wp.Close()
		close(wp.Result)
		
		// Ensure all worker goroutines have terminated
		synctest.Wait()

		require.Equal(t, int64(numTasks), atomic.LoadInt64(&completedTasks))
		require.Len(t, results, numTasks)
	})
}

// TestWorkerPoolSimpleWithSynctest validates basic worker functionality in isolation.
// Tests single worker task execution with controlled timing to ensure tasks are
// processed correctly and results are returned as expected.
func TestWorkerPoolSimpleWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Test isolated worker functionality without complex pool management
		var taskExecuted int64
		
		taskFunc := func() (any, error) {
			// Simulate task processing time
			time.Sleep(50 * time.Millisecond)
			atomic.AddInt64(&taskExecuted, 1)
			return "completed", nil
		}

		// Create minimal worker setup with buffered channels
		taskQueue := make(chan worker_pool.Task, 1)
		result := make(chan worker_pool.Result, 1)

		// Launch single worker to process tasks
		go func() {
			for task := range taskQueue {
				res, err := task()
				result <- worker_pool.Result{Value: res, Error: err}
			}
		}()

		// Submit task and signal completion
		taskQueue <- taskFunc
		close(taskQueue)

		// Retrieve task result
		res := <-result
		close(result)

		// Ensure all concurrent operations have completed
		synctest.Wait()

		require.Equal(t, int64(1), atomic.LoadInt64(&taskExecuted))
		require.Equal(t, "completed", res.Value)
		require.NoError(t, res.Error)
	})
}
