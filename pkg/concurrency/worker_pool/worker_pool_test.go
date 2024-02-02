package worker_pool

import (
	"os"
	"sync"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func Test_WorkerPool(t *testing.T) {
	wp := New(10)

	f := func() (any, error) {
		// some operation
		return nil, nil
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
		for range wp.Result {
			wg.Done()
		}
	}()

	<-done
	wg.Wait()
	close(wp.Result)

	t.Cleanup(func() {
		wp.Close()
	})
}
