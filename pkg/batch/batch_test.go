//go:build unit

package batch

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
	"golang.org/x/sync/errgroup"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	t.Run("Create a new batch", func(t *testing.T) {
		ctx, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()
		eg, ctx := errgroup.WithContext(ctx)

		aggrCB := func(args []*Item[string]) error {
			for _, item := range args {
				time.Sleep(time.Microsecond * 100) // Emulate long work

				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}

			return nil
		}

		b, err := New(ctx, aggrCB)
		require.NoError(t, err)

		requests := []string{"A", "B", "C", "D"}
		for _, request := range requests {
			res := b.Push(request)

			req := request // Capture range variable
			eg.Go(func() error {
				val, ok := <-res
				require.True(t, ok)
				require.Equal(t, req, val)
				return nil
			})
		}

		err = eg.Wait()
		require.NoError(t, err)
	})

	t.Run("Check context cancellation", func(t *testing.T) {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*10)
		defer cancelFunc()

		eg, ctx := errgroup.WithContext(ctx)

		aggrCB := func(args []*Item[string]) error {
			for _, item := range args {
				time.Sleep(time.Second * 10) // Emulate long work

				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}

			return nil
		}

		requests := []string{"A", "B", "C", "D"}

		b, err := New(ctx, aggrCB)
		require.NoError(t, err)

		for _, request := range requests {
			res := b.Push(request)

			eg.Go(func() error {
				_, ok := <-res
				require.False(t, ok)
				return nil
			})
		}

		err = eg.Wait()
		require.NoError(t, err)
	})
}
