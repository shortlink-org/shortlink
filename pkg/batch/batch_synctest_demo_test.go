package batch

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

// TestBatchDemoTraditional shows traditional batch testing approach
func TestBatchDemoTraditional(t *testing.T) {
	start := time.Now()
	
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var processedCount int64
	callback := func(items []*Item[string]) error {
		atomic.AddInt64(&processedCount, int64(len(items)))
		for _, item := range items {
			item.CallbackChannel <- item.Item
			close(item.CallbackChannel)
		}
		return nil
	}

	// Create batch with 50ms interval
	batch, _ := New(ctx, callback, WithInterval[string](50*time.Millisecond), WithSize[string](10))

	// Add items
	ch1 := batch.Push("item1")
	ch2 := batch.Push("item2")

	// Wait for time-based flush (real time)
	time.Sleep(100 * time.Millisecond)

	require.Equal(t, "item1", <-ch1)
	require.Equal(t, "item2", <-ch2)
	require.Equal(t, int64(2), atomic.LoadInt64(&processedCount))

	elapsed := time.Since(start)
	t.Logf("Traditional batch test took: %v", elapsed)
}

// TestBatchDemoSynctest shows the same test with synctest - instant execution
func TestBatchDemoSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var processedCount int64
		callback := func(items []*Item[string]) error {
			atomic.AddInt64(&processedCount, int64(len(items)))
			for _, item := range items {
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		// Create batch with 50ms interval
		batch, errChan := New(ctx, callback, WithInterval[string](50*time.Millisecond), WithSize[string](10))

		// Add items
		ch1 := batch.Push("item1")
		ch2 := batch.Push("item2")

		// Wait for time-based flush (instant in synctest)
		synctest.Wait()

		require.Equal(t, "item1", <-ch1)
		require.Equal(t, "item2", <-ch2)
		require.Equal(t, int64(2), atomic.LoadInt64(&processedCount))

		// Clean up
		cancel()
		for range errChan {
			// Drain error channel
		}

		elapsed := time.Since(start)
		t.Logf("Synctest batch test took: %v (instant execution)", elapsed)
	})
}

// TestBatchSizeFlushWithSynctest demonstrates size-based flushing
func TestBatchSizeFlushWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var flushes [][]string
		callback := func(items []*Item[string]) error {
			var batch []string
			for _, item := range items {
				batch = append(batch, item.Item)
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			flushes = append(flushes, batch)
			return nil
		}

		// Create batch with size 3 and long interval
		batch, errChan := New(ctx, callback, 
			WithSize[string](3), 
			WithInterval[string](10*time.Second)) // Long interval to force size-based flush

		// Add items one by one
		ch1 := batch.Push("a")
		ch2 := batch.Push("b")
		
		// No flush yet (size not reached)
		synctest.Wait()
		require.Len(t, flushes, 0)

		// Add third item - should trigger flush
		ch3 := batch.Push("c")
		synctest.Wait()

		// Verify flush occurred
		require.Len(t, flushes, 1)
		require.Equal(t, []string{"a", "b", "c"}, flushes[0])
		
		require.Equal(t, "a", <-ch1)
		require.Equal(t, "b", <-ch2)
		require.Equal(t, "c", <-ch3)

		// Clean up
		cancel()
		for range errChan {
			// Drain error channel
		}
	})
}

// TestBatchErrorHandlingWithSynctest shows error handling in batch processing
func TestBatchErrorHandlingWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var callCount int64
		callback := func(items []*Item[string]) error {
			count := atomic.AddInt64(&callCount, 1)
			if count == 2 {
				// Second batch fails
				return context.DeadlineExceeded
			}
			
			for _, item := range items {
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		batch, errChan := New(ctx, callback, WithSize[string](2), WithInterval[string](100*time.Millisecond))

		// First batch - should succeed
		ch1 := batch.Push("item1")
		ch2 := batch.Push("item2")
		synctest.Wait()

		require.Equal(t, "item1", <-ch1)
		require.Equal(t, "item2", <-ch2)

		// Second batch - should fail
		ch3 := batch.Push("item3")
		ch4 := batch.Push("item4")
		synctest.Wait()

		// Channels should be closed without data due to error
		_, ok3 := <-ch3
		_, ok4 := <-ch4
		require.False(t, ok3)
		require.False(t, ok4)

		// Should receive error
		cancel()
		
		var errors []error
		for err := range errChan {
			if err != nil {
				errors = append(errors, err)
			}
		}
		require.Len(t, errors, 1)
		require.Equal(t, context.DeadlineExceeded, errors[0])
	})
}

// TestConcurrentBatchOperationsWithSynctest demonstrates concurrent access patterns
func TestConcurrentBatchOperationsWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var totalProcessed int64
		callback := func(items []*Item[string]) error {
			atomic.AddInt64(&totalProcessed, int64(len(items)))
			for _, item := range items {
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		batch, errChan := New(ctx, callback, WithSize[string](5), WithInterval[string](25*time.Millisecond))

		// Launch multiple goroutines adding items concurrently
		var wg sync.WaitGroup
		const numGoroutines = 10
		const itemsPerGoroutine = 3

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				
				for j := 0; j < itemsPerGoroutine; j++ {
					// Add some delay variation
					time.Sleep(time.Duration(id+j) * time.Millisecond)
					
					ch := batch.Push("data")
					<-ch // Wait for processing
				}
			}(i)
		}

		wg.Wait()

		// All items should be processed
		totalItems := int64(numGoroutines * itemsPerGoroutine)
		require.Equal(t, totalItems, atomic.LoadInt64(&totalProcessed))

		// Clean up
		cancel()
		for range errChan {
			// Drain error channel
		}
	})
}