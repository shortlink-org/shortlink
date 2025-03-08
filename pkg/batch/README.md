## Batch Processing

This package offers a robust batch processing system for aggregating and processing items efficiently in batches. 
It's designed with concurrency and efficiency in mind, following Go's concurrency patterns.

### Features

- **Batch processing:** Groups items for efficient bulk processing.
- **Concurrency Safe:** Thread-safe for reliable operation under concurrent loads.
- **Configurable:** Custom batch sizes and flush intervals via options.
- **Context Support:** Graceful shutdowns and cancellations without storing contexts.
- **Generics:** Utilizes Go's generics for type safety.
- **Error Reporting:** Callback errors are reported through an error channel.

### Usage

Here's an example of how to use the batch processing package:

```go
package main

import (
	"context"
	"fmt"
	"time"

  "github.com/shortlink-org/shortlink/pkg/batch"
)

func main() {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  // Define the callback function to process a batch of items.
  callback := func(items []*batch.Item[string]) error {
    for _, item := range items {
      // Simulate processing work.
      time.Sleep(10 * time.Millisecond)
      item.CallbackChannel <- item.Item + " processed"
      close(item.CallbackChannel)
    }
    
    return nil
  }

  // Create a new batch processor with custom options.
  // Note: New returns an error channel to report callback errors.
  b, errChan := batch.New(ctx, callback, batch.WithSize[string](5), batch.WithInterval[string](time.Second))

  // Process errors from the error channel.
  go func() {
    for err := range errChan {
      fmt.Println("Error:", err)
    }
  }()

  // Push items into the batch processor.
  for i := 0; i < 20; i++ {
    resChan := b.Push(fmt.Sprintf("Item %d", i))
    
    go func(ch chan string) {
      if result, ok := <-ch; ok {
        fmt.Println(result)
      } else {
        fmt.Println("Channel closed before processing")
      }
    }(resChan)
  }

  // Wait to ensure all items are processed.
  time.Sleep(2 * time.Second)
}
```

## References

- [Batch Processing](https://en.wikipedia.org/wiki/Batch_processing)
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
