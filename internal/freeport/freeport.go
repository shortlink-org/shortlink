/*
Get free port
*/
package freeport

import "net"

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) { // nolint unused
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close() // nolint errcheck
	return l.Addr().(*net.TCPAddr).Port, nil
}
