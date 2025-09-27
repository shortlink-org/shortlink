package vector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "vector")
	t.Attr("component", "types")

		t.Attr("type", "unit")
		t.Attr("package", "vector")
		t.Attr("component", "types")
	
	tests := []struct {
		slice    []int
		element  int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, 5},
		{[]int{}, 1, 0},
		{[]int{1, 1, 1, 1}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 0, 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Attr("type", "unit")
			t.Attr("package", "vector")
			t.Attr("component", "types")

			result := Find(tt.slice, tt.element)
			require.Equal(t, tt.expected, result)
		})
	}
}
