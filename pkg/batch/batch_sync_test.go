package batch

import (
	"context"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewSync(t *testing.T) {
	t.Run("Returns cleanly after context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		aggrCB := func(args []*Item[string]) error {
			for _, item := range args {
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		// Call NewSync in a goroutine because it blocks until ctx is done.
		done := make(chan struct{})
		var (
			b   *Batch[string]
			err error
		)
		go func() {
			b, err = NewSync(ctx, aggrCB)
			close(done)
		}()

		// Give the goroutine a moment to start.
		time.Sleep(5 * time.Millisecond)
		// Trigger shutdown.
		cancel()

		// Wait for NewSync to return.
		<-done

		require.NotNil(t, b)
		require.NoError(t, err)
	})
}

// TestBatchProcessingWithSynctest demonstrates deterministic batch testing
func TestBatchProcessingWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		
		var processedItems []string
		var callbackCount int64

		aggrCB := func(items []*Item[string]) error {
			atomic.AddInt64(&callbackCount, 1)
			for _, item := range items {
				processedItems = append(processedItems, item.Item)
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		// Create batch with 100ms interval and size 3
		batch, errChan := New(ctx, aggrCB, WithInterval[string](100*time.Millisecond), WithSize[string](3))

		// Add items that don't reach the batch size
		ch1 := batch.Push("item1")
		ch2 := batch.Push("item2")

		// Wait for potential flush (shouldn't happen yet)
		synctest.Wait()

		// Verify items haven't been processed yet (batch size not reached)
		require.Equal(t, int64(0), atomic.LoadInt64(&callbackCount))

		// Add third item to trigger size-based flush
		ch3 := batch.Push("item3")

		// Wait for flush to complete
		synctest.Wait()

		// Verify items were processed
		require.Equal(t, int64(1), atomic.LoadInt64(&callbackCount))
		require.Equal(t, "item1", <-ch1)
		require.Equal(t, "item2", <-ch2)
		require.Equal(t, "item3", <-ch3)

		// Add one more item
		ch4 := batch.Push("item4")

		// Wait for time-based flush (100ms interval)
		// In synctest, time advances instantly when all goroutines are blocked
		synctest.Wait()

		// Verify the single item was processed
		require.Equal(t, int64(2), atomic.LoadInt64(&callbackCount))
		require.Equal(t, "item4", <-ch4)

		// Cancel to clean up and wait for error channel to close
		cancel()
		
		// Wait for error channel to close and check no errors occurred
		var errors []error
		for err := range errChan {
			if err != nil {
				errors = append(errors, err)
			}
		}
		require.Empty(t, errors)
	})
}

// TestBatchCancellationWithSynctest tests proper cleanup when context is cancelled
func TestBatchCancellationWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		var processedCount int64

		aggrCB := func(items []*Item[string]) error {
			atomic.AddInt64(&processedCount, int64(len(items)))
			for _, item := range items {
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		// Create batch with long interval to test cancellation
		batch, errChan := New(ctx, aggrCB, WithInterval[string](10*time.Second), WithSize[string](100))

		// Add some items
		ch1 := batch.Push("item1")
		ch2 := batch.Push("item2")

		// Cancel context immediately
		cancel()

		// Wait for batch to process cancellation
		synctest.Wait()

		// Verify channels are closed (items should be dropped on cancellation)
		_, ok1 := <-ch1
		_, ok2 := <-ch2
		require.False(t, ok1, "channel should be closed")
		require.False(t, ok2, "channel should be closed")

		// Wait for error channel to close
		var errors []error
		for err := range errChan {
			if err != nil {
				errors = append(errors, err)
			}
		}

		// Should have no errors from the callback itself
		require.Empty(t, errors)
	})
}

// TestBatchTimeBasedFlushWithSynctest tests time-based flushing behavior
func TestBatchTimeBasedFlushWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		
		var flushCount int64

		aggrCB := func(items []*Item[string]) error {
			atomic.AddInt64(&flushCount, 1)
			for _, item := range items {
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		// Create batch with 50ms interval and high size limit
		batch, errChan := New(ctx, aggrCB, WithInterval[string](50*time.Millisecond), WithSize[string](1000))

		// Add items over multiple intervals
		ch1 := batch.Push("item1")
		
		// Wait for first interval flush
		synctest.Wait()
		require.Equal(t, int64(1), atomic.LoadInt64(&flushCount))
		require.Equal(t, "item1", <-ch1)

		// Add more items
		ch2 := batch.Push("item2")
		ch3 := batch.Push("item3")

		// Wait for second interval flush
		synctest.Wait()
		require.Equal(t, int64(2), atomic.LoadInt64(&flushCount))
		require.Equal(t, "item2", <-ch2)
		require.Equal(t, "item3", <-ch3)

		// Clean up
		cancel()
		for range errChan {
			// Drain error channel
		}
	})
}
