// Package io_uring implements a high-level Go wrapper to perform
// file read/write operations using liburing.
package io_uring

/*
#cgo LDFLAGS: -luring
#include <fcntl.h>
#include <liburing.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/stat.h>
extern int queue_init();
extern int push_read_request(int, off_t);
extern int push_write_request(int, void *, off_t);
extern int pop_request();
extern int queue_submit(int);
extern void queue_exit();
*/
import "C" // nolint typecheck
import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"syscall"
	"unsafe"
)

// TODO: move to local struct fields.
var (
	globalMut  sync.RWMutex
	quitChan   chan struct{}
	submitChan chan *request
	pollChan   chan struct{}
	errChan    chan error

	cbMut sync.RWMutex
	cbMap map[uintptr]cbInfo
)

//export read_callback
func read_callback(iovecs *C.struct_iovec, length C.int, fd C.int) {
	// Here be dragons.
	intLen := int(length)
	slice := (*[1 << 28]C.struct_iovec)(unsafe.Pointer(iovecs))[:intLen:intLen]
	// Can be optimized further with more unsafe.
	var buf bytes.Buffer
	for i := 0; i < intLen; i++ {
		_, err := buf.Write(C.GoBytes(slice[i].iov_base, C.int(slice[i].iov_len)))
		if err != nil {
			errChan <- fmt.Errorf("error during buffer write: %v", err)
		}
	}
	cbMut.Lock()
	cbMap[uintptr(fd)].close()
	cbMap[uintptr(fd)].readCb(buf.Bytes())
	delete(cbMap, uintptr(fd))
	cbMut.Unlock()
}

//export write_callback
func write_callback(written C.int, fd C.int) {
	cbMut.Lock()
	cbMap[uintptr(fd)].close()
	cbMap[uintptr(fd)].writeCb(int(written))
	delete(cbMap, uintptr(fd))
	cbMut.Unlock()
}

// Init is used to initialize the ring and setup some global state.
func Init() error {
	ret := int(C.queue_init())
	if ret < 0 {
		return fmt.Errorf("%v", syscall.Errno(-ret))
	}
	globalMut.Lock()
	quitChan = make(chan struct{})
	pollChan = make(chan struct{})
	errChan = make(chan error)
	submitChan = make(chan *request)
	cbMap = make(map[uintptr]cbInfo)
	globalMut.Unlock()
	go startLoop()
	return nil
}

// Cleanup must be called to close the ring.
func Cleanup() {
	quitChan <- struct{}{}
	C.queue_exit()
	close(submitChan)
	close(errChan)
}

// Err is a channel that needs to be read to receive internal errors
// that can happen during interaction with the ring or during callbacks.
// It must be read from after calling Init, otherwise the processing might
// get stuck.
func Err() <-chan error {
	globalMut.RLock()
	defer globalMut.RUnlock()
	return errChan
}

func startLoop() {
	queueSize := 0
	for {
		select {
		case sqe := <-submitChan:
			switch sqe.code {
			case opCodeRead:
				// We populate the cbMap to be called later from the callback from C.
				// No need for locking here.
				cbMap[sqe.f.Fd()] = cbInfo{
					readCb: sqe.readCb,
					close:  sqe.f.Close,
				}

				ret := int(C.push_read_request(C.int(sqe.f.Fd()), C.long(sqe.size)))
				if ret < 0 {
					errChan <- fmt.Errorf("error while pushing read request: %v", syscall.Errno(-ret))
					continue
				}
			case opCodeWrite:
				// No need for locking here.
				cbMap[sqe.f.Fd()] = cbInfo{
					writeCb: sqe.writeCb,
					close:   sqe.f.Close,
				}

				var ptr unsafe.Pointer
				if len(sqe.buf) == 0 {
					// In case it's a zero byte write, we explicitly take the pointer
					// to a zero byte slice. Because we can't do &sqe.buf.
					zeroBytes := []byte("")
					ptr = unsafe.Pointer(&zeroBytes)
				} else {
					ptr = unsafe.Pointer(&sqe.buf[0])
				}

				ret := int(C.push_write_request(C.int(sqe.f.Fd()), ptr, C.long(len(sqe.buf))))
				if ret < 0 {
					errChan <- fmt.Errorf("error while pushing write_request: %v", syscall.Errno(-ret))
					continue
				}
			}

			queueSize++
			if queueSize > queueThreshold { // if queue_size > threshold, then pop all.
				// TODO: maybe just pop one
				submitAndPop(queueSize)
				queueSize = 0
			}
		case <-pollChan:
			if queueSize > 0 {
				submitAndPop(queueSize)
				queueSize = 0
			}
		case <-quitChan:
			// possibly drain channel.
			// pop_request till everything is done.
			return
		}
	}
}

// ReadFile reads a file from the given path and returns the result as a byte slice
// in the passed callback function.
func ReadFile(path string, cb func(buf []byte)) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	submitChan <- &request{
		code:   opCodeRead,
		f:      f,
		size:   fi.Size(),
		readCb: cb,
	}
	return nil
}

// WriteFile writes data to a file at the given path. After the file is written,
// it then calls the callback with the number of bytes written.
func WriteFile(path string, data []byte, perm os.FileMode, cb func(written int)) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}

	submitChan <- &request{
		code:    opCodeWrite,
		buf:     data,
		f:       f,
		writeCb: cb,
	}
	return nil
}

// Poll signals the kernel to read all pending entries from the submission queue
// and waits until all entries have been read from the completion queue.
func Poll() {
	// TODO: do we allow user to set wait_nr ?
	pollChan <- struct{}{}
}

func submitAndPop(queueSize int) {
	// Submit the queue with wait_nr set to current pending queue size.
	ret := int(C.queue_submit(C.int(queueSize)))
	if ret < 0 {
		errChan <- fmt.Errorf("error while submitting: %v", syscall.Errno(-ret))
		return
	}
	// Pop until the queue is empty.
	for queueSize > 0 {
		ret := int(C.pop_request())
		if ret != 0 {
			errChan <- fmt.Errorf("error while popping: %v", syscall.Errno(-ret))
			if syscall.Errno(-ret) != syscall.EAGAIN { // Do not decrement if nothing was read.
				queueSize--
			}
			continue
		}
		queueSize--
	}
}
