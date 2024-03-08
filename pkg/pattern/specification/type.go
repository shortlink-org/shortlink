package specification

// Specification is the interface that wraps the IsSatisfiedBy method.
type Specification[T any] interface {
	IsSatisfiedBy(*T) error
}
