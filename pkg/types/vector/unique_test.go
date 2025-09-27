package vector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "vector")
	t.Attr("component", "types")

	tests := []struct {
		slice    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 2, 3, 4, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{5, 4, 3, 2, 1, 1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Attr("type", "unit")
			t.Attr("package", "vector")
			t.Attr("component", "types")

			result := Unique(tt.slice)
			require.ElementsMatch(t, tt.expected, result)
		})
	}
}
