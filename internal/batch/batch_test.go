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
		ctx, cancel := context.WithCancel(ctx)
		aggrCB := func(args []*Item) interface{} {
			// Get string
			for _, item := range args {
				//time.Sleep(time.Second * 2) // Emulate long work

				item.CB <- item.Item.(string)
			}

			return nil
		}

		b, err := New(ctx, aggrCB)
		assert.Nil(t, err)

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

		time.Sleep(time.Second * 1)
		cancel()
		for key := range request {
			wg.Add(1)
			res, err := b.Push(request[key])
			assert.Nil(t, err)
			go func(key int) {
				assert.Equal(t, <-res, "ctx close")
				wg.Done()
			}(key)
		}
		wg.Wait()
	})
}
