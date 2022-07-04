package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	// create
	tree := New(func(a, b int) int {
		return a - b
	})

	// crud operation
	expected := new(*int)
	assert.Equal(t, *expected, tree.find(1).Value())

	tree.insert(1)
	tree.insert(3)
	tree.insert(2)
	tree.insert(5)
	tree.insert(4)
	tree.insert(7)
	tree.insert(9)

	// aggregation operation
	assert.Equal(t, 4, *tree.find(4).Value())
	assert.Equal(t, 1, *tree.Min().Value())
	assert.Equal(t, 9, *tree.Max().Value())

	// drop operation
	tree.delete(5)
	assert.Equal(t, 9, *tree.Max().Value())
	tree.delete(9)
	tree.delete(1)
	tree.insert(10)
	tree.insert(11)
	assert.Equal(t, 11, *tree.Max().Value())
}
