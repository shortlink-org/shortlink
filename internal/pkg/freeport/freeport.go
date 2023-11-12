/*
Get free port
*/
package freeport

import (
	"net"
)

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	defer func() {
		_ = l.Close()
	}()

	port, ok := l.Addr().(*net.TCPAddr)
	if !ok {
		return 0, ErrNoFreePort
	}

	return port.Port, nil
}
