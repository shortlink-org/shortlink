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
	t.Attr("type", "unit")
	t.Attr("package", "freeport")
	t.Attr("component", "freeport")
	
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
	b.Attr("type", "benchmark")
	b.Attr("package", "freeport")
	b.Attr("component", "freeport")
	
	for i := 0; i < b.N; i++ {
		GetFreePort()
	}
}
