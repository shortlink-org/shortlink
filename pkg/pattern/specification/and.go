package specification

import (
	"errors"
)

// AndSpecification is a composite specification that represents the logical AND of two other specifications.
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
