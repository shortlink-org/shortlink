package worker_pool

import (
	"fmt"
	"sync"
	"testing"
)

func Test_WorkerPool(t *testing.T) {
	wp := New(10)

	f := func() {
		t.Log("Hello")
	}

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
		for resp := range wp.Result {
			fmt.Println(resp)
			wg.Done()
		}
	}()

	<-done
	wg.Wait()
	close(wp.Result)
}
