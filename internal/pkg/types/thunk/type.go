package thunk

import (
	"github.com/shortlink-org/shortlink/internal/pkg/types/options"
)

type Thunk[T any] struct {
	doer func() T           // action being thunked
	o    *options.Option[T] // cache for complete thunk data
}
