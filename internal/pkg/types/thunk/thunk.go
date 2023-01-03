package thunk

import (
	"github.com/shortlink-org/shortlink/internal/pkg/types/options"
)

func New[T any](doer func() T) *Thunk[T] {
	return &Thunk[T]{
		doer: doer,
		o:    options.New[T](),
	}
}

func (t *Thunk[T]) Force() T {
	if t.o.IsSome() {
		return t.o.Apply()
	}

	t.o.Set(t.doer())

	return t.o.Apply()
}
