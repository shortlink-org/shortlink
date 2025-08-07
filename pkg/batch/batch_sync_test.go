package batch

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewSync(t *testing.T) {
	t.Run("Returns cleanly after context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		aggrCB := func(args []*Item[string]) error {
			for _, item := range args {
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		// Call NewSync in a goroutine because it blocks until ctx is done.
		done := make(chan struct{})
		var (
			b   *Batch[string]
			err error
		)
		go func() {
			b, err = NewSync(ctx, aggrCB)
			close(done)
		}()

		// Give the goroutine a moment to start.
		time.Sleep(5 * time.Millisecond)
		// Trigger shutdown.
		cancel()

		// Wait for NewSync to return.
		<-done

		require.NotNil(t, b)
		require.NoError(t, err)
	})
}
