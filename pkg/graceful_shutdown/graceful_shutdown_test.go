package graceful_shutdown_test

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/graceful_shutdown"
)

func TestGracefulShutdown(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "graceful_shutdown")
	t.Attr("component", "graceful_shutdown")

		t.Attr("type", "unit")
		t.Attr("package", "graceful_shutdown")
		t.Attr("component", "graceful_shutdown")
	

	go func() {
		time.Sleep(time.Second) // wait for GracefulShutdown to start listening for signals

		err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		require.NoError(t, err)
	}()

	receivedSignal := graceful_shutdown.GracefulShutdown()

	require.Equal(t, os.Signal(syscall.SIGTERM), receivedSignal)
}
