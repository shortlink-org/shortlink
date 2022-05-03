package options

type Option[T any] struct {
	val *T
}

type OptionError struct{}
