package specification

type Specification[T any] interface {
	IsSatisfiedBy(*T) error
}
