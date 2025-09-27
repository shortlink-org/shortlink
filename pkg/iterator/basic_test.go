package iterator

import (
	"reflect"
	"testing"
)

func TestBasicRange(t *testing.T) {
	result := Collect(Range(0, 5))
	expected := []int{0, 1, 2, 3, 4}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBasicLines(t *testing.T) {
	input := "line1\nline2\nline3"
	result := Collect(Lines(input))
	expected := []string{"line1", "line2", "line3"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBasicSplitSeq(t *testing.T) {
	input := "a,b,c"
	result := Collect(SplitSeq(input, ","))
	expected := []string{"a", "b", "c"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBasicFilter(t *testing.T) {
	numbers := Range(0, 10)
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	result := Collect(evens)
	expected := []int{0, 2, 4, 6, 8}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBasicMap(t *testing.T) {
	numbers := Range(1, 4)
	doubled := Map(numbers, func(x int) int { return x * 2 })
	result := Collect(doubled)
	expected := []int{2, 4, 6}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBasicChain(t *testing.T) {
	seq1 := Range(1, 3)
	seq2 := Range(10, 12)
	chained := Chain(seq1, seq2)
	result := Collect(chained)
	expected := []int{1, 2, 10, 11}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSliceBackwards(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	result := Collect(Backwards(data))
	expected := []int{5, 4, 3, 2, 1}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSliceChunk(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	var result [][]int
	for chunk := range Chunk(data, 2) {
		// Create a copy to avoid slice reference issues
		chunkCopy := make([]int, len(chunk))
		copy(chunkCopy, chunk)
		result = append(result, chunkCopy)
	}
	expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := Collect(Keys(m))
	
	// Since map iteration order is not guaranteed, check length and presence
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
	
	expectedKeys := map[string]bool{"a": true, "b": true, "c": true}
	for _, key := range keys {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key: %s", key)
		}
	}
}

func TestSortedKeys(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	result := Collect(SortedKeys(m))
	expected := []string{"a", "b", "c"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBytesLines(t *testing.T) {
	data := []byte("line1\nline2\nline3")
	var result []string
	for line := range LinesBytes(data) {
		result = append(result, string(line))
	}
	expected := []string{"line1", "line2", "line3"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFieldsSeqBasic(t *testing.T) {
	input := "hello world test"
	result := Collect(FieldsSeq(input))
	expected := []string{"hello", "world", "test"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}