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
			wg.Add(1)
			wp.Push(f)
		}

		close(done)
	}()

	go func() {
		for range wp.Result {
			wg.Done()
		}
	}()

	<-done
	wg.Wait()
	close(wp.Result)

	t.Cleanup(func() {
		wp.Close()
	})
}

// TestWorkerPoolWithSynctest demonstrates deterministic worker pool testing
func TestWorkerPoolWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const numWorkers = 3
		const numTasks = 10
		
		wp := worker_pool.New(numWorkers)

		var completedTasks int64
		var results []any
		var resultWg sync.WaitGroup
		resultWg.Add(1)

		// Create tasks that simulate some work
		taskFunc := func() (any, error) {
			// Simulate some processing time
			time.Sleep(10 * time.Millisecond)
			return atomic.AddInt64(&completedTasks, 1), nil
		}

		// Collect results in background
		go func() {
			defer resultWg.Done()
			for result := range wp.Result {
				results = append(results, result.Value)
				if len(results) == numTasks {
					return
				}
			}
		}()

		// Submit all tasks
		for i := 0; i < numTasks; i++ {
			wp.Push(taskFunc)
		}

		// Wait for all tasks to complete
		resultWg.Wait()

		// Close the worker pool to stop background goroutines
		wp.Close()
		close(wp.Result)
		
		// Wait for workers to finish
		synctest.Wait()

		require.Equal(t, int64(numTasks), atomic.LoadInt64(&completedTasks))
		require.Len(t, results, numTasks)
	})
}

// TestWorkerPoolSimpleWithSynctest demonstrates basic synctest usage with controlled execution
func TestWorkerPoolSimpleWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Simple test that doesn't rely on long-running background goroutines
		var taskExecuted int64
		
		taskFunc := func() (any, error) {
			// Simulate processing with time
			time.Sleep(50 * time.Millisecond)
			atomic.AddInt64(&taskExecuted, 1)
			return "completed", nil
		}

		// Create a minimal worker pool scenario
		taskQueue := make(chan worker_pool.Task, 1)
		result := make(chan worker_pool.Result, 1)

		// Single worker
		go func() {
			for task := range taskQueue {
				res, err := task()
				result <- worker_pool.Result{Value: res, Error: err}
			}
		}()

		// Submit task
		taskQueue <- taskFunc
		close(taskQueue)

		// Get result
		res := <-result
		close(result)

		// Wait for all operations to complete
		synctest.Wait()

		require.Equal(t, int64(1), atomic.LoadInt64(&taskExecuted))
		require.Equal(t, "completed", res.Value)
		require.NoError(t, res.Error)
	})
}
