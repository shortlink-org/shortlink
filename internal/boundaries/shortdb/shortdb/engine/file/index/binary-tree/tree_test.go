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
	assert.Equal(t, *expected, tree.find(1).Value()) //nolint:revive // ignore for test

	tree.insert(1) //nolint:revive // ignore for test
	tree.insert(3) //nolint:revive // ignore for test
	tree.insert(2) //nolint:revive // ignore for test
	tree.insert(5) //nolint:revive // ignore for test
	tree.insert(4) //nolint:revive // ignore for test
	tree.insert(7) //nolint:revive // ignore for test
	tree.insert(9) //nolint:revive // ignore for test

	// aggregation operation
	assert.Equal(t, 4, *tree.find(4).Value()) //nolint:revive // ignore for test
	assert.Equal(t, 1, *tree.Min().Value())   //nolint:revive // ignore for test
	assert.Equal(t, 9, *tree.Max().Value())   //nolint:revive // ignore for test

	// drop operation
	tree.delete(5)                           //nolint:revive // ignore for test
	assert.Equal(t, 9, *tree.Max().Value())  //nolint:revive // ignore for test
	tree.delete(9)                           //nolint:revive // ignore for test
	tree.delete(1)                           //nolint:revive // ignore for test
	tree.insert(10)                          //nolint:revive // ignore for test
	tree.insert(11)                          //nolint:revive // ignore for test
	assert.Equal(t, 11, *tree.Max().Value()) //nolint:revive // ignore for test
}
