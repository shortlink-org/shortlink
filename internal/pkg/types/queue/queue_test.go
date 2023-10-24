package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := New[string](5) //nolint:revive // it's ok
	q.Push("hello world!")

	str := q.Pop()

	assert.Equal(t, "hello world!", str)
}
