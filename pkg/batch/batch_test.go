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
	t.Attr("type", "unit")
	t.Attr("package", "batch")
	t.Attr("component", "batch")

		t.Attr("type", "unit")
		t.Attr("package", "batch")
		t.Attr("component", "batch")
	
	t.Run("Create a new batch", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "batch")
		t.Attr("component", "batch")

		ctx, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()
		eg, ctx := errgroup.WithContext(ctx)

		aggrCB := func(args []*Item[string]) error {
			for _, item := range args {
				time.Sleep(100 * time.Microsecond) // Emulate long work

				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}

			return nil
		}

		b, errChan := New(ctx, aggrCB)
		require.NotNil(t, b)
		require.NotNil(t, errChan)

		requests := []string{"A", "B", "C", "D"}
		for _, request := range requests {
			res := b.Push(request)

			eg.Go(func() error {
				val, ok := <-res
				require.True(t, ok)
				require.Equal(t, request, val)

				return nil
			})
		}

		require.NoError(t, eg.Wait())

		// Cancel the context to trigger cleanup.
		cancelFunc()

		// Drain the error channel (should be closed without any error)
		for range errChan {
			// No errors expected.
		}
	})

	t.Run("Check context cancellation", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "batch")
		t.Attr("component", "batch")

		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelFunc()

		eg, ctx := errgroup.WithContext(ctx)

		aggrCB := func(args []*Item[string]) error {
			for _, item := range args {
				time.Sleep(10 * time.Second) // Emulate long work

				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}

			return nil
		}

		requests := []string{"A", "B", "C", "D"}
		b, errChan := New(ctx, aggrCB)
		require.NotNil(t, b)
		require.NotNil(t, errChan)

		for _, request := range requests {
			res := b.Push(request)

			eg.Go(func() error {
				_, ok := <-res
				require.False(t, ok)

				return nil
			})
		}

		require.NoError(t, eg.Wait())
	})
}
