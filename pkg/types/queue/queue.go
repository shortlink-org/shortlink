package queue

import (
	"sync"
)

// Queue represents a generic FIFO queue.
type Queue[T any] struct {
	data []T
	mu   sync.Mutex
}

// New creates a new Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push adds an element to the end of the queue.
func (q *Queue[T]) Push(v T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, v)
}

// Pop removes and returns the first element of the queue. Returns a zero value and false if the queue is empty.
func (q *Queue[T]) Pop() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) == 0 {
		var zero T
		return zero, false
	}

	v := q.data[0]
	q.data = q.data[1:]

	return v, true
}

// Size returns the current number of elements in the queue.
func (q *Queue[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.data)
}

// Clean clears all elements from the queue.
func (q *Queue[T]) Clean() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = nil // Resets the slice to its zero value, effectively clearing it.
}
