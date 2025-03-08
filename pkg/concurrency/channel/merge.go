package channel

import (
	"sync"
)

// Merge merges multiple channels into one.
func Merge[T any](items ...<-chan T) <-chan T {
	out := make(chan T)

	var wg sync.WaitGroup

	wg.Add(len(items))

	for _, item := range items {
		go func(c <-chan T) {
			for n := range c {
				out <- n
			}

			wg.Done()
		}(item)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
