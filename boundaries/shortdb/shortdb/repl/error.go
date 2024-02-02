package repl

import (
	"errors"
)

// ErrStatus is a sentinel error to indicate that the REPL should exit.
var ErrStatus = errors.New("")
