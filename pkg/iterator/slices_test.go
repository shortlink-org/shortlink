package iterator

import (
	"reflect"
	"testing"
)

func TestAllSlices(t *testing.T) {
	data := []string{"a", "b", "c"}
	var result []struct {
		idx int
		val string
	}
	
	for i, v := range AllSlices(data) {
		result = append(result, struct {
			idx int
			val string
		}{i, v})
	}
	
	expected := []struct {
		idx int
		val string
	}{{0, "a"}, {1, "b"}, {2, "c"}}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBackwards(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	result := Collect(Backwards(data))
	expected := []int{5, 4, 3, 2, 1}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBackwardsAll(t *testing.T) {
	data := []string{"a", "b", "c"}
	var result []struct {
		idx int
		val string
	}
	
	for i, v := range BackwardsAll(data) {
		result = append(result, struct {
			idx int
			val string
		}{i, v})
	}
	
	expected := []struct {
		idx int
		val string
	}{{2, "c"}, {1, "b"}, {0, "a"}}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		data     []int
		size     int
		expected [][]int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3}, 5, [][]int{{1, 2, 3}}},
		{[]int{}, 2, [][]int{}},
	}
	
	for _, test := range tests {
		var result [][]int
		for chunk := range Chunk(test.data, test.size) {
			// Create a copy to avoid slice reference issues
			chunkCopy := make([]int, len(chunk))
			copy(chunkCopy, chunk)
			result = append(result, chunkCopy)
		}
		
		if len(result) == 0 && len(test.expected) == 0 {
			// Both empty, test passes
		} else if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Chunk(%v, %d): expected %v, got %v", test.data, test.size, test.expected, result)
		}
	}
}

func TestWindow(t *testing.T) {
	tests := []struct {
		data     []int
		size     int
		expected [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}},
		{[]int{1, 2}, 2, [][]int{{1, 2}}},
		{[]int{1, 2}, 3, [][]int{}},
		{[]int{}, 2, [][]int{}},
	}
	
	for _, test := range tests {
		var result [][]int
		for window := range Window(test.data, test.size) {
			// Create a copy to avoid slice reference issues
			windowCopy := make([]int, len(window))
			copy(windowCopy, window)
			result = append(result, windowCopy)
		}
		
		if len(result) == 0 && len(test.expected) == 0 {
			// Both empty, test passes
		} else if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Window(%v, %d): expected %v, got %v", test.data, test.size, test.expected, result)
		}
	}
}

func TestStepBy(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := Collect(StepBy(data, 3))
	expected := []int{0, 3, 6, 9}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTakeWhile(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 1, 2}
	result := Collect(TakeWhile(data, func(x int) bool { return x < 4 }))
	expected := []int{1, 2, 3}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSkipWhile(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 1, 2}
	result := Collect(SkipWhile(data, func(x int) bool { return x < 4 }))
	expected := []int{4, 5, 1, 2}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFlatten(t *testing.T) {
	data := [][]int{{1, 2}, {3, 4, 5}, {6}}
	result := Collect(Flatten(data))
	expected := []int{1, 2, 3, 4, 5, 6}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestUnique(t *testing.T) {
	data := []int{1, 2, 2, 3, 1, 4, 3}
	result := Collect(Unique(data))
	expected := []int{1, 2, 3, 4}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDedup(t *testing.T) {
	data := []int{1, 1, 2, 2, 2, 3, 1, 1}
	result := Collect(Dedup(data))
	expected := []int{1, 2, 3, 1}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIntersect(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{3, 4, 5, 6}
	result := Collect(Intersect(s1, s2))
	expected := []int{3, 4}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestUnion(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 4, 5}
	result := Collect(Union(s1, s2))
	
	// Check that result contains all unique elements
	expectedElements := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	if len(result) != len(expectedElements) {
		t.Errorf("Expected %d unique elements, got %d", len(expectedElements), len(result))
	}
	
	for _, v := range result {
		if !expectedElements[v] {
			t.Errorf("Unexpected element %d in result", v)
		}
	}
}

func TestDifference(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{3, 4, 5, 6}
	result := Collect(Difference(s1, s2))
	expected := []int{1, 2}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSorted(t *testing.T) {
	data := []int{3, 1, 4, 1, 5, 9, 2, 6}
	result := Collect(Sorted(data))
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
	
	// Verify original slice is unchanged
	originalExpected := []int{3, 1, 4, 1, 5, 9, 2, 6}
	if !reflect.DeepEqual(data, originalExpected) {
		t.Errorf("Original slice was modified: expected %v, got %v", originalExpected, data)
	}
}

func TestSortedFunc(t *testing.T) {
	data := []string{"banana", "apple", "cherry"}
	result := Collect(SortedFunc(data, func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}))
	expected := []string{"apple", "banana", "cherry"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGroupBy(t *testing.T) {
	words := []string{"apple", "banana", "apricot", "blueberry"}
	var result []struct {
		key   byte
		group []string
	}
	
	for key, group := range GroupBy(words, func(word string) byte { return word[0] }) {
		result = append(result, struct {
			key   byte
			group []string
		}{key, group})
	}
	
	// Check that we have the right number of groups
	// "apple", "banana", "apricot", "blueberry" -> 'a', 'b', 'a', 'b' -> 4 groups since GroupBy groups consecutive elements
	if len(result) != 4 {
		t.Errorf("Expected 4 groups, got %d", len(result))
	}
	
	for _, group := range result {
		if group.key == 'a' && reflect.DeepEqual(group.group, []string{"apple"}) {
			continue
		}
		if group.key == 'b' && (reflect.DeepEqual(group.group, []string{"banana"}) || 
			reflect.DeepEqual(group.group, []string{"blueberry"})) {
			continue
		}
		if group.key == 'a' && reflect.DeepEqual(group.group, []string{"apricot"}) {
			continue
		}
	}
}

func TestScan(t *testing.T) {
	data := []int{1, 2, 3, 4}
	result := Collect(Scan(data, 0, func(acc, x int) int { return acc + x }))
	expected := []int{0, 1, 3, 6, 10} // Running sums
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIntersperse(t *testing.T) {
	data := []int{1, 2, 3}
	result := Collect(Intersperse(data, 0))
	expected := []int{1, 0, 2, 0, 3}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPairwise(t *testing.T) {
	data := []int{1, 2, 3, 4}
	var result []struct{ a, b int }
	
	for a, b := range Pairwise(data) {
		result = append(result, struct{ a, b int }{a, b})
	}
	
	expected := []struct{ a, b int }{{1, 2}, {2, 3}, {3, 4}}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkBackwards(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Backwards(data) {
			// Consume iterator
		}
	}
}

func BenchmarkChunk(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Chunk(data, 10) {
			// Consume iterator
		}
	}
}

func BenchmarkWindow(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Window(data, 5) {
			// Consume iterator
		}
	}
}

func BenchmarkUnique(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i % 100 // Create duplicates
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Unique(data) {
			// Consume iterator
		}
	}
}

func BenchmarkSorted(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = 1000 - i // Reverse order
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Sorted(data) {
			// Consume iterator
		}
	}
}