//go:build darwin || linux

package io_uring

import (
	"os"
)

func Init() error {
	return nil
}

func Cleanup() {}

func Err() <-chan error {
	return nil
}

func ReadFile(path string, cb func(buf []byte)) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	cb(content)

	return nil
}

func Poll() {}

func WriteFile(path string, data []byte, perm os.FileMode, cb func(written int)) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm) // #nosec
	if err != nil {
		return err
	}

	written, err := f.Write(data)
	if err != nil {
		return err
	}

	cb(written)

	return nil
}
