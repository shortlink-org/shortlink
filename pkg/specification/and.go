package specification

import (
	"errors"
)

type AndSpecification[T any] struct {
	specs []Specification[T]
}

func (a *AndSpecification[T]) IsSatisfiedBy(account *T) error {
	var errs error

	for _, spec := range a.specs {
		if err := spec.IsSatisfiedBy(account); err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}

func NewAndSpecification[T any](specs ...Specification[T]) *AndSpecification[T] {
	return &AndSpecification[T]{
		specs: specs,
	}
}
