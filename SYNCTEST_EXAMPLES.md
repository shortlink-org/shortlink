# Testing with `testing/synctest` in Go

This repository demonstrates how to use Go's `testing/synctest` package to write deterministic, fast, and reliable tests for concurrent code.

## What is `testing/synctest`?

The `testing/synctest` package provides a controlled environment ("bubble") where:
- Time can be manipulated via a fake clock
- All goroutines use the same fake time
- Time advances only when all goroutines are durably blocked
- Tests execute instantly while maintaining correct time-based behavior

## Key Benefits

1. **Speed**: Tests with time-dependent operations execute instantly
2. **Determinism**: Eliminates flakiness from real-time dependencies  
3. **Reliability**: Consistent results across different environments

## Examples in this Codebase

### 1. Rate Limiter Testing

See: `pkg/concurrency/rate_limiter/rate_limiter_synctest_demo_test.go`

**Traditional Approach** (slow, potentially flaky):
```go
func TestRateLimiterTraditional(t *testing.T) {
    // Takes ~100ms+ to execute
    rl, _ := rate_limiter.New(ctx, 1, 100*time.Millisecond)
    rl.Wait() // First token
    rl.Wait() // Waits real 100ms for refill
}
```

**Synctest Approach** (instant, deterministic):
```go
func TestRateLimiterSynctest(t *testing.T) {
    synctest.Test(t, func(t *testing.T) {
        // Executes instantly, but time behavior is preserved
        rl, _ := rate_limiter.New(ctx, 1, 100*time.Millisecond)
        rl.Wait() // First token
        rl.Wait() // Time advances instantly to next refill
    })
}
```

### 2. Batch Processing Testing

See: `pkg/batch/batch_synctest_demo_test.go`

Shows how to test time-based and size-based batch flushing without waiting for real time to pass.

### 3. Comprehensive Tutorial

See: `pkg/concurrency/synctest_tutorial_test.go`

Contains 11 examples covering:
- Basic time control
- Channel operations
- Timers and tickers
- Context timeouts
- Multiple goroutines
- Resource pools
- Producer-consumer patterns

## Core Synctest Functions

### `synctest.Test(t, testFunc)`
Wraps your test function to run in a controlled bubble environment.

### `synctest.Wait()`
Blocks until all goroutines in the bubble are durably blocked. Use this to wait for operations to complete.

## Best Practices

1. **Use for time-dependent code**: Perfect for testing timeouts, intervals, delays
2. **Ensure proper cleanup**: Cancel contexts and close channels to avoid deadlocks
3. **Avoid infinite loops**: Synctest works best with finite operations
4. **Use `synctest.Wait()`**: To ensure operations complete before assertions

## Common Patterns

### Testing Timeouts
```go
synctest.Test(t, func(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
    defer cancel()
    
    select {
    case <-longOperation():
        t.Fatal("should have timed out")
    case <-ctx.Done():
        // Expected timeout
    }
})
```

### Testing Periodic Operations
```go
synctest.Test(t, func(t *testing.T) {
    ticker := time.NewTicker(50*time.Millisecond)
    defer ticker.Stop()
    
    for i := 0; i < 5; i++ {
        <-ticker.C
        // Process tick
    }
    // All 5 ticks happen instantly
})
```

### Testing Race Conditions
```go
synctest.Test(t, func(t *testing.T) {
    var counter int64
    var wg sync.WaitGroup
    
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1)
        }()
    }
    
    wg.Wait()
    require.Equal(t, int64(100), counter)
})
```

## Running the Examples

```bash
# Run basic synctest tutorial
go test -v ./pkg/concurrency -run "TestSynctest[1-6]"

# Run rate limiter comparisons
go test -v ./pkg/concurrency/rate_limiter -run "TestRateLimiterDemo"

# Run batch processing comparisons  
go test -v ./pkg/batch -run "TestBatchDemo"

# Run all synctest examples
go test -v ./pkg/concurrency ./pkg/batch ./pkg/concurrency/rate_limiter -run "Synctest"
```

## Key Takeaways

- Synctest eliminates `time.Sleep()` in tests
- Tests run instantly while preserving time-based logic
- Perfect for testing concurrent systems with time dependencies
- Requires proper cleanup to avoid deadlocks
- Makes tests more reliable and much faster