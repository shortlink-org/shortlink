package worker_pool

import (
	"sync"
	"testing"
)

func Test_WorkerPool(t *testing.T) {
	wp := New(10)

	f := func() {}

	wg := sync.WaitGroup{}
	done := make(chan struct{})

	go func() {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			wp.Push(f)
		}
		close(done)
	}()

	go func() {
		for range wp.Result {
			wg.Done()
		}
	}()

	<-done
	wg.Wait()
	close(wp.Result)
}
