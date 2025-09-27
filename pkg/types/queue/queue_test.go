package queue

import (
	"testing"
)

// TestQueue conducts a series of subtests on the Queue.
func TestQueue(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "queue")
	t.Attr("component", "types")

	// Subtest for Push and Pop
	t.Run("PushPop", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "queue")
		t.Attr("component", "types")

		q := New[int]()
		q.Push(1)
		q.Push(2)

		if size := q.Size(); size != 2 {
			t.Errorf("Expected size 2, got %d", size)
		}

		if val, ok := q.Pop(); !ok || val != 1 {
			t.Errorf("Expected Pop to return 1, got %d", val)
		}

		if val, ok := q.Pop(); !ok || val != 2 {
			t.Errorf("Expected Pop to return 2, got %d", val)
		}

		if size := q.Size(); size != 0 {
			t.Errorf("Expected size 0, got %d", size)
		}
	})

	// Subtest for Size
	t.Run("Size", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "queue")
		t.Attr("component", "types")

		q := New[int]()
		q.Push(1)
		q.Push(2)

		if size := q.Size(); size != 2 {
			t.Errorf("Expected size 2, got %d", size)
		}
	})

	// Subtest for Clean
	t.Run("Clean", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "queue")
		t.Attr("component", "types")

		q := New[int]()
		q.Push(1)
		q.Push(2)
		q.Clean()

		if size := q.Size(); size != 0 {
			t.Errorf("Expected size 0 after Clean, got %d", size)
		}
	})
}
