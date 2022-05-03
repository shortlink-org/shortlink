package queue

func New[T any](size int) Queue[T] {
	return Queue[T]{
		data: make(chan T, size),
	}
}

func (q Queue[T]) Push(v T) {
	q.data <- v
}

func (q Queue[T]) Pop() T {
	return <-q.data
}

func (q Queue[T]) TryPop() (T, bool) {
	select {
	case v := <-q.data:
		return v, true
	default:
		var zero T
		return zero, false
	}
}

func Zero[T any]() T {
	var zero T
	return zero
}
