package vector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "vector")
	t.Attr("component", "types")

		t.Attr("type", "unit")
		t.Attr("package", "vector")
		t.Attr("component", "types")
	
	tests := []struct {
		slice    []int
		element  int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 6, false},
		{[]int{}, 1, false},
		{[]int{1, 1, 1, 1}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Attr("type", "unit")
			t.Attr("package", "vector")
			t.Attr("component", "types")

			result := Contains(tt.slice, tt.element)
			require.Equal(t, tt.expected, result)
		})
	}
}
