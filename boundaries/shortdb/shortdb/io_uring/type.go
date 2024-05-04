//go:build !darwin

package io_uring

import (
	"os"
)

type opCode int //nolint:unused // false positive

const (
	opCodeRead  opCode = iota + 1 //nolint:unused // false positive
	opCodeWrite                   //nolint:unused // false positive
)

const queueThreshold = 5 //nolint:unused // false positive

type (
	readCallback  func([]byte) //nolint:unused // false positive
	writeCallback func(int)    //nolint:unused // false positive
)

// request contains info to send to the submission queue.
//
//nolint:unused // false positive
type request struct {
	code    opCode
	f       *os.File
	buf     []byte
	size    int64
	readCb  readCallback
	writeCb writeCallback
}

//nolint:unused // false positive
type cbInfo struct {
	readCb  readCallback
	writeCb writeCallback
	close   func() error
}
