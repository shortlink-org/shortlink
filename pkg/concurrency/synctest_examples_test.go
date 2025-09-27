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

// This file demonstrates the benefits of using testing/synctest for concurrent code testing.
// It shows side-by-side comparisons of traditional time-based tests vs synctest equivalents.

// Traditional approach: flaky, slow, and non-deterministic
func TestTimeoutTraditional(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		// Simulate work that might take varying amounts of time
		time.Sleep(50 * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		// Success - but timing is unpredictable
	case <-ctx.Done():
		t.Fatal("operation timed out")
	}
}

// Synctest approach: fast, deterministic, and reliable
func TestTimeoutWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		done := make(chan struct{})
		go func() {
			// Simulate work - in synctest, this happens instantly
			time.Sleep(50 * time.Millisecond)
			close(done)
		}()

		select {
		case <-done:
			// Success - timing is deterministic
		case <-ctx.Done():
			t.Fatal("operation timed out")
		}
	})
}

// Traditional approach for testing periodic operations
func TestPeriodicOperationTraditional(t *testing.T) {
	var counter int64
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 5; i++ {
			<-ticker.C
			atomic.AddInt64(&counter, 1)
		}
	}()

	// Wait for completion - this takes ~50ms in real time
	<-done
	require.Equal(t, int64(5), atomic.LoadInt64(&counter))
}

// Synctest approach for periodic operations - instant and deterministic
func TestPeriodicOperationWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var counter int64
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()

		done := make(chan struct{})
		go func() {
			defer close(done)
			for i := 0; i < 5; i++ {
				<-ticker.C
				atomic.AddInt64(&counter, 1)
			}
		}()

		// Wait for completion - this happens instantly in synctest
		<-done
		require.Equal(t, int64(5), atomic.LoadInt64(&counter))
	})
}

// Complex example: testing a background processor with multiple time dependencies
type BackgroundProcessor struct {
	ctx       context.Context
	cancel    context.CancelFunc
	items     chan string
	processed chan string
	done      chan struct{}
}

func NewBackgroundProcessor() *BackgroundProcessor {
	ctx, cancel := context.WithCancel(context.Background())
	bp := &BackgroundProcessor{
		ctx:       ctx,
		cancel:    cancel,
		items:     make(chan string, 10),
		processed: make(chan string, 10),
		done:      make(chan struct{}),
	}

	go bp.run()
	return bp
}

func (bp *BackgroundProcessor) run() {
	defer close(bp.done)
	
	ticker := time.NewTicker(25 * time.Millisecond)
	defer ticker.Stop()

	batch := make([]string, 0, 5)

	for {
		select {
		case <-bp.ctx.Done():
			// Process remaining batch
			if len(batch) > 0 {
				bp.processBatch(batch)
			}
			return
		case item := <-bp.items:
			batch = append(batch, item)
			if len(batch) >= 5 {
				bp.processBatch(batch)
				batch = batch[:0]
			}
		case <-ticker.C:
			if len(batch) > 0 {
				bp.processBatch(batch)
				batch = batch[:0]
			}
		}
	}
}

func (bp *BackgroundProcessor) processBatch(items []string) {
	// Simulate processing time
	time.Sleep(5 * time.Millisecond)
	for _, item := range items {
		bp.processed <- "processed:" + item
	}
}

func (bp *BackgroundProcessor) AddItem(item string) {
	select {
	case bp.items <- item:
	case <-bp.ctx.Done():
	}
}

func (bp *BackgroundProcessor) GetProcessed() <-chan string {
	return bp.processed
}

func (bp *BackgroundProcessor) Stop() {
	bp.cancel()
	<-bp.done
	close(bp.processed)
}

// Traditional test - slow and potentially flaky
func TestBackgroundProcessorTraditional(t *testing.T) {
	processor := NewBackgroundProcessor()
	defer processor.Stop()

	// Add items
	processor.AddItem("item1")
	processor.AddItem("item2")
	processor.AddItem("item3")

	// Wait for processing (time-based, unreliable)
	time.Sleep(50 * time.Millisecond)

	var results []string
	for {
		select {
		case result := <-processor.GetProcessed():
			results = append(results, result)
		case <-time.After(10 * time.Millisecond):
			// Timeout waiting for more results
			goto done
		}
	}
done:

	require.GreaterOrEqual(t, len(results), 3)
}

// Synctest approach - fast, deterministic, and reliable
func TestBackgroundProcessorWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		processor := NewBackgroundProcessor()
		defer processor.Stop()

		// Add items that will trigger size-based flush
		processor.AddItem("item1")
		processor.AddItem("item2")
		processor.AddItem("item3")
		processor.AddItem("item4")
		processor.AddItem("item5")

		// Wait for processing to complete
		synctest.Wait()

		var results []string
		for {
			select {
			case result := <-processor.GetProcessed():
				results = append(results, result)
			default:
				goto done
			}
		}
	done:

		require.Len(t, results, 5)
		require.Contains(t, results, "processed:item1")
		require.Contains(t, results, "processed:item5")

		// Test time-based flushing
		processor.AddItem("item6")

		// Wait for time-based flush (25ms ticker)
		synctest.Wait()

		// Should have one more processed item
		result := <-processor.GetProcessed()
		require.Equal(t, "processed:item6", result)
	})
}

// Example of testing race conditions with synctest
func TestRaceConditionDetection(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var sharedCounter int64
		var mu sync.Mutex
		numGoroutines := 100

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		// Launch concurrent goroutines that modify shared state
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				
				// Simulate some work
				time.Sleep(time.Duration(id) * time.Microsecond)
				
				// Critical section
				mu.Lock()
				current := atomic.LoadInt64(&sharedCounter)
				time.Sleep(1 * time.Microsecond) // Increase chance of race
				atomic.StoreInt64(&sharedCounter, current+1)
				mu.Unlock()
			}(i)
		}

		wg.Wait()

		// In synctest, this executes deterministically
		require.Equal(t, int64(numGoroutines), atomic.LoadInt64(&sharedCounter))
	})
}

// Example: testing complex state machine with time transitions
type StateMachine struct {
	state     string
	mu        sync.RWMutex
	timer     *time.Timer
	ctx       context.Context
	cancel    context.CancelFunc
	stateLog  []string
}

func NewStateMachine() *StateMachine {
	ctx, cancel := context.WithCancel(context.Background())
	sm := &StateMachine{
		state:    "idle",
		ctx:      ctx,
		cancel:   cancel,
		stateLog: []string{"idle"},
	}
	return sm
}

func (sm *StateMachine) Transition(newState string, delay time.Duration) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if sm.timer != nil {
		sm.timer.Stop()
	}

	sm.timer = time.AfterFunc(delay, func() {
		sm.mu.Lock()
		defer sm.mu.Unlock()
		sm.state = newState
		sm.stateLog = append(sm.stateLog, newState)
	})
}

func (sm *StateMachine) GetState() string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.state
}

func (sm *StateMachine) GetStateLog() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return append([]string{}, sm.stateLog...)
}

func (sm *StateMachine) Stop() {
	sm.cancel()
	sm.mu.Lock()
	if sm.timer != nil {
		sm.timer.Stop()
	}
	sm.mu.Unlock()
}

// Testing state machine transitions with synctest
func TestStateMachineWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		sm := NewStateMachine()
		defer sm.Stop()

		require.Equal(t, "idle", sm.GetState())

		// Schedule transition to "processing" after 100ms
		sm.Transition("processing", 100*time.Millisecond)
		
		// State should still be idle
		require.Equal(t, "idle", sm.GetState())

		// Wait for time to advance
		synctest.Wait()

		// Now state should have transitioned
		require.Equal(t, "processing", sm.GetState())
		require.Equal(t, []string{"idle", "processing"}, sm.GetStateLog())

		// Schedule another transition
		sm.Transition("completed", 50*time.Millisecond)
		synctest.Wait()

		require.Equal(t, "completed", sm.GetState())
		require.Equal(t, []string{"idle", "processing", "completed"}, sm.GetStateLog())
	})
}