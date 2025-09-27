//go:build unit

package freeport

import (
	"net"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	require.NoError(t, err)

	// Try to listen on the port
	l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
	defer func() {
		_ = l.Close()
	}()
	require.NoError(t, err)
}

func BenchmarkGetFreePort(b *testing.B) {
	for b.Loop() {
		GetFreePort()
	}
}
