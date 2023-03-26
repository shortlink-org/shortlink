package binary_tree

import (
	"github.com/segmentio/encoding/json"
)

// Tree is a binary tree.
type Tree[T any] struct {
	// cmp compares two T values.
	cmp     func(T, T) int
	Left    *Tree[T]
	Right   *Tree[T]
	Val     *T
	Payload []byte
}

func (t *Tree[T]) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Tree[T]) UnMarshal(bytes []byte, i interface{}) error {
	// TODO implement me
	panic("implement me")
}

func New[T any](cmp func(T, T) int) *Tree[T] {
	return &Tree[T]{
		cmp: cmp,
	}
}

func (t *Tree[T]) Find(key T) []byte {
	return t.find(key).Payload
}

func (t *Tree[T]) find(key T) *Tree[T] {
	if t.Val == nil {
		return t
	}

	switch cmp := t.cmp(key, *t.Val); {
	case cmp < 0:
		return t.Left.find(key)
	case cmp > 0:
		return t.Right.find(key)
	default:
		return t
	}
}

func (t *Tree[T]) Insert(key T) error {
	t.insert(key)
	return nil
}

func (t *Tree[T]) insert(key T) *Tree[T] {
	if t.Val == nil {
		t.Val = &key
		return t
	}

	switch cmp := t.cmp(key, *t.Val); {
	case cmp < 0:
		if t.Left == nil {
			t.Left = New(t.cmp)
			t.Left.Val = &key

			return t.Left
		}

		return t.Left.insert(key)
	case cmp > 0:
		if t.Right == nil {
			t.Right = New(t.cmp)
			t.Right.Val = &key

			return t.Right
		}

		return t.Right.insert(key)
	default:
		return t
	}
}

func (t *Tree[T]) Delete(key T) error {
	t.delete(key)

	return nil
}

func (t *Tree[T]) delete(key T) *Tree[T] {
	if t == nil {
		return t
	}

	switch cmp := t.cmp(key, *t.Value()); {
	case cmp < 0:
		t.Left = t.Left.delete(key)
	case cmp > 0:
		t.Right = t.Right.delete(key)
	default:
		if t.Left == nil && t.Right == nil {
			t = nil
		} else if t.Left == nil {
			t = t.Right
		} else if t.Right == nil {
			t = t.Left
		} else {
			t.Val = t.Right.Min().Value()
			t.Right = t.Right.delete(*t.Value())
		}
	}

	return t
}

func (t *Tree[T]) Min() *Tree[T] {
	if t.Left == nil {
		return t
	}

	return t.Left.Min()
}

func (t *Tree[T]) Max() *Tree[T] {
	if t.Right == nil {
		return t
	}

	return t.Right.Max()
}

func (t *Tree[T]) Value() *T {
	return t.Val
}
