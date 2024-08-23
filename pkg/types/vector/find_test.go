package vector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
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
			result := Find(tt.slice, tt.element)
			require.Equal(t, tt.expected, result)
		})
	}
}
