//go:build unit

package batch

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
	"golang.org/x/sync/errgroup"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	t.Run("Create new a batch", func(t *testing.T) {
		ctx := context.Background()
		eg, ctx := errgroup.WithContext(ctx)

		aggrCB := func(args []*Item) any {
			for _, item := range args {
				time.Sleep(time.Microsecond * 100) // Emulate long work

				item.CallbackChannel <- item.Item.(string)
			}

			return nil
		}

		b, err := New(ctx, aggrCB)
		require.NoError(t, err)
		t.Cleanup(b.Stop)

		requests := []string{"A", "B", "C", "D"}
		for key := range requests {
			request := requests[key]
			res := b.Push(request)

			eg.Go(func() error {
				assert.Equal(t, <-res, request)
				return nil
			})
		}

		err = eg.Wait()
		require.NoError(t, err)
	})

	t.Run("Check close context", func(t *testing.T) {
		ctx := context.Background()
		ctx, cancelFunc := context.WithTimeout(ctx, time.Millisecond*10)
		defer cancelFunc()

		eg, ctx := errgroup.WithContext(ctx)

		aggrCB := func(args []*Item) any {
			// Get string
			for _, item := range args {
				time.Sleep(time.Second * 10) // Emulate long work

				item.CallbackChannel <- item.Item.(string)
			}

			return nil
		}

		requests := []string{"A", "B", "C", "D"}

		b, err := New(ctx, aggrCB)
		require.NoError(t, err)
		t.Cleanup(b.Stop)

		for key := range requests {
			request := requests[key]
			res := b.Push(request)

			eg.Go(func() error {
				assert.Equal(t, struct{}{}, <-res)
				return nil
			})
		}

		err = eg.Wait()
		require.NoError(t, err)
	})
}
