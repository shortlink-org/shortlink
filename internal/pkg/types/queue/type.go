package queue

type Queue[T any] struct {
	data chan T
}
