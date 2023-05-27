package channel

import (
	"sync"
)

// Merge merges multiple channels into one.
func Merge[T any](cs ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan T) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
