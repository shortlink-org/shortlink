package channel

import (
	"sync"
)

// Merge merges multiple channels into one.
func Merge[T any](items ...<-chan T) <-chan T {
	out := make(chan T)

	var wg sync.WaitGroup

	for _, item := range items {
		wg.Go(func() {
			for n := range item {
				out <- n
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
