package worker_pool_test

import (
	"os"
	"sync"
	"testing"

	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/pkg/concurrency/worker_pool"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func Test_WorkerPool(t *testing.T) {
	wp := worker_pool.New(10)

	f := func() (any, error) {
		// some operation
		return nil, nil
	}

	wg := sync.WaitGroup{}
	done := make(chan struct{})

	go func() {
		for range 1000 {
			wp.Push(f)
			wg.Go(func() {
				<-wp.Result
			})
		}

		close(done)
	}()

	<-done
	wg.Wait()
	close(wp.Result)

	t.Cleanup(func() {
		wp.Close()
	})
}
