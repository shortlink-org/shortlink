package specification

import (
	"errors"
)

type OrSpecification[T any] struct {
	specs []Specification[T]
}

func (o *OrSpecification[T]) IsSatisfiedBy(account *T) error {
	var errs error

	for _, spec := range o.specs {
		err := spec.IsSatisfiedBy(account)
		if err == nil {
			return nil
		}

		errs = errors.Join(errs, err)
	}

	return errs
}

func NewOrSpecification[T any](specs ...Specification[T]) *OrSpecification[T] {
	return &OrSpecification[T]{
		specs: specs,
	}
}
