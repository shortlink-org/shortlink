package channel_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/pkg/concurrency/channel"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestMerge(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "channel")
	t.Attr("component", "concurrency")

		t.Attr("type", "unit")
		t.Attr("package", "channel")
		t.Attr("component", "concurrency")
	
	// Create two channels
	ch1 := make(chan any, 5)
	ch2 := make(chan any, 5)

	// Populate channels with some data
	for i := range 5 {
		ch1 <- i
		ch2 <- i + 5
	}

	// Close channels after data has been sent
	close(ch1)
	close(ch2)

	// Merge channels
	chMerged := channel.Merge(ch1, ch2)

	ctx := t.Context()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// We're expecting 10 elements
	for range 10 {
		select {
		case result, ok := <-chMerged:
			require.True(t, ok, "channel was closed prematurely")
			t.Logf("Received: %v", result)
		case <-ctx.Done():
			require.Fail(t, "test timed out")
		}
	}

	// Test if a merged channel is closed after all data is received
	_, ok := <-chMerged
	require.False(t, ok, "channel was not closed as expected")
}
