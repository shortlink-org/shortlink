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
	t.Attr("type", "unit")
	t.Attr("package", "batch")
	t.Attr("component", "batch")

		t.Attr("type", "unit")
		t.Attr("package", "batch")
		t.Attr("component", "batch")
	
	t.Run("Returns cleanly after context cancellation", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "batch")
		t.Attr("component", "batch"), cancel := context.WithCancel(t.Context())

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

// TestBatchProcessingWithSynctest verifies batch processing behavior with deterministic timing.
// This test validates both size-based and time-based flush mechanisms using synctest
// to eliminate timing dependencies and ensure consistent test execution.
func TestBatchProcessingWithSynctest(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "batch")
	t.Attr("component", "batch")

		t.Attr("type", "unit")
		t.Attr("package", "batch")
		t.Attr("component", "batch")
	
	synctest.Test(t, func(t *testing.T) {, cancel := context.WithCancel(t.Context())
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

		// Configure batch with 100ms flush interval and size threshold of 3 items
		batch, errChan := New(ctx, aggrCB, WithInterval[string](100*time.Millisecond), WithSize[string](3))

		// Add items below the size threshold to test time-based flushing
		ch1 := batch.Push("item1")
		ch2 := batch.Push("item2")

		// Ensure no premature flushing occurs before size threshold is met
		synctest.Wait()

		// Verify no processing has occurred yet since size threshold not reached
		require.Equal(t, int64(0), atomic.LoadInt64(&callbackCount))

		// Add third item to trigger size-based flush mechanism
		ch3 := batch.Push("item3")

		// Allow size-based flush to complete
		synctest.Wait()

		// Verify size-based flush processed all three items
		require.Equal(t, int64(1), atomic.LoadInt64(&callbackCount))
		require.Equal(t, "item1", <-ch1)
		require.Equal(t, "item2", <-ch2)
		require.Equal(t, "item3", <-ch3)

		// Add single item to test time-based flush behavior
		ch4 := batch.Push("item4")

		// Allow time-based flush to occur (100ms interval)
		// synctest advances time instantly when all goroutines reach stable state
		synctest.Wait()

		// Verify time-based flush processed the remaining item
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

// TestBatchCancellationWithSynctest verifies proper resource cleanup and graceful shutdown
// when the batch context is cancelled. Ensures that pending items are handled correctly
// and no goroutines are leaked during cancellation scenarios.
func TestBatchCancellationWithSynctest(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "batch")
	t.Attr("component", "batch")

		t.Attr("type", "unit")
		t.Attr("package", "batch")
		t.Attr("component", "batch")
	
	synctest.Test(t, func(t *testing.T) {, cancel := context.WithCancel(t.Context())
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

// TestBatchTimeBasedFlushWithSynctest validates the time-based flush mechanism.
// Verifies that batches are flushed according to the configured interval when
// the size threshold is not reached, ensuring predictable batch processing behavior.
func TestBatchTimeBasedFlushWithSynctest(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "batch")
	t.Attr("component", "batch")

		t.Attr("type", "unit")
		t.Attr("package", "batch")
		t.Attr("component", "batch")
	
	synctest.Test(t, func(t *testing.T) {, cancel := context.WithCancel(t.Context())
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
