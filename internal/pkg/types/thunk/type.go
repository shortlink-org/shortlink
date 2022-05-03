package thunk

import (
	"github.com/batazor/shortlink/internal/pkg/types/options"
)

type Thunk[T any] struct {
	doer func() T           // action being thunked
	o    *options.Option[T] // cache for complete thunk data
}
