package specification

// NotSpecification is a composite specification that represents the logical NOT of another specification.
type NotSpecification[T any] struct {
	spec Specification[T]
}

func (n *NotSpecification[T]) IsSatisfiedBy(account *T) error {
	err := n.spec.IsSatisfiedBy(account)
	if err != nil {
		return nil
	}

	return err
}

func NewNotSpecification[T any](spec Specification[T]) *NotSpecification[T] {
	return &NotSpecification[T]{
		spec: spec,
	}
}
