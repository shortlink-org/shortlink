package io_uring

import (
	"os"
)

type opCode int

const (
	opCodeRead opCode = iota + 1
	opCodeWrite
)

const queueThreshold = 5

type readCallback func([]byte)
type writeCallback func(int)

// request contains info to send to the submission queue.
type request struct {
	code    opCode
	f       *os.File
	buf     []byte
	size    int64
	readCb  readCallback
	writeCb writeCallback
}

type cbInfo struct {
	readCb  readCallback
	writeCb writeCallback
	close   func() error
}
