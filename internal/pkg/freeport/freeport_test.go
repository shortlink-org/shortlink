//go:build unit
// +build unit

package freeport

import (
	"net"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	require.NoError(t, err)
	assert.NotEqual(t, port, 0)

	// Try to listen on the port
	l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
	defer func() {
		_ = l.Close()
	}()
	require.NoError(t, err)
}

func BenchmarkGetFreePort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFreePort()
	}
}
