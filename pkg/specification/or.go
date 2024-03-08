package specification

import (
	"errors"
)

type OrSpecification[T any] struct {
	specs []Specification[T]
}

func (o *OrSpecification[T]) IsSatisfiedBy(account *T) error {
	for _, spec := range o.specs {
		if err := spec.IsSatisfiedBy(account); err == nil {
			return nil
		}
	}

	return errors.New("no specification is satisfied")
}

func NewOrSpecification[T any](specs ...Specification[T]) *OrSpecification[T] {
	return &OrSpecification[T]{
		specs: specs,
	}
}
