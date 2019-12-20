package freeport

import (
	"net"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	assert.Nil(t, err)
	assert.NotEqual(t, port, 0)

	// Try to listen on the port
	l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
	assert.Nil(t, err)
	defer l.Close()
}
