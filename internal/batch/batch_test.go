package batch

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

func TestNew(t *testing.T) {
	t.Run("Create new a batch", func(t *testing.T) {
		// Add events
		wg := sync.WaitGroup{}

		ctx := context.Background()
		aggrCB := func(args []*Item) interface{} {
			// Get string
			for _, item := range args {
				time.Sleep(time.Microsecond * 100) // Emulate long work

				item.CB <- item.Item.(string)
			}

			return nil
		}

		b, err := New(context.TODO(), aggrCB)
		assert.Nil(t, err)

		go b.Run(ctx)

		request := []string{"A", "B", "C", "D"}
		for key := range request {
			wg.Add(1)
			res, err := b.Push(request[key])
			assert.Nil(t, err)
			go func(key int) {
				assert.Equal(t, <-res, request[key])
				wg.Done()
			}(key)
		}

		wg.Wait()
	})

	// TODO: Fix test with context.Done
	//t.Run("Check close context", func(t *testing.T) {
	//	// Add events
	//	wg := sync.WaitGroup{}
	//
	//	aggrCB := func(args []*Item) interface{} {
	//		// Get string
	//		for _, item := range args {
	//			time.Sleep(time.Microsecond * 100) // Emulate long work
	//
	//			item.CB <- item.Item.(string)
	//		}
	//
	//		return nil
	//	}
	//
	//	ctx := context.Background()
	//	ctx, cancel := context.WithCancel(ctx)
	//	cancel()
	//
	//	request := []string{"A", "B", "C", "D"}
	//
	//	b, err := New(ctx, aggrCB)
	//	assert.Nil(t, err)
	//
	//	go b.Run(ctx)
	//
	//	for key := range request {
	//		wg.Add(1)
	//		res, err := b.Push(request[key])
	//		assert.Nil(t, err)
	//		go func() {
	//			assert.Equal(t, <-res, "ctx close")
	//			wg.Done()
	//		}()
	//	}
	//	wg.Wait()
	//})
}
