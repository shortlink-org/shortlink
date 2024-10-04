## Batch Processing

This package offers a robust batch processing system for aggregating and processing items efficiently in batches.
It's designed with concurrency and efficiency in mind, aligning with Go's concurrency patterns.

### Features

- **Batch processing:** Groups items for efficient bulk processing.
- **Concurrency Safe:** Thread-safe for reliable operation under concurrent loads.
- **Configurable:** Allows for custom batch sizes and tick intervals.
- **Context Support:** Supports graceful shutdowns and cancellations.
- **Generics:** Utilizes Go's generics for type safety.

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
	ctx := context.Background()

	// Define the callback function
	callback := func(items []*batch.Item[string]) error {
		for _, item := range items {
			// Process item
			time.Sleep(time.Millisecond * 10) // Simulate work
			item.CallbackChannel <- item.Item + " processed"
			close(item.CallbackChannel)
		}
		return nil
	}

	// Create a new batch processor
	b, err := batch.New(ctx, callback, batch.WithSize, batch.WithInterval[string](time.Second))
	if err != nil {
		panic(err)
	}

	// Push items into the batch processor
	for i := 0; i < 20; i++ {
		resChan := b.Push(fmt.Sprintf("Item %d", i))
		go func(ch chan string) {
			result, ok := <-ch
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("Channel closed before processing")
			}
		}(resChan)
	}

	// Wait to ensure all items are processed
	time.Sleep(2 * time.Second)
}
```

## References

- [Batch Processing](https://en.wikipedia.org/wiki/Batch_processing)
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
