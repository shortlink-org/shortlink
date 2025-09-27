package iterator

import (
	"fmt"
	"reflect"
	"testing"
	"unicode"
)

func TestRange(t *testing.T) {
	tests := []struct {
		start, end int
		expected   []int
	}{
		{0, 5, []int{0, 1, 2, 3, 4}},
		{1, 4, []int{1, 2, 3}},
		{5, 5, []int{}},
		{3, 1, []int{}}, // Invalid range
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Range(%d,%d)", test.start, test.end), func(t *testing.T) {
			result := Collect(Range(test.start, test.end))
			if len(result) == 0 && len(test.expected) == 0 {
				// Both are empty, test passes
				return
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestTake(t *testing.T) {
	numbers := Range(0, 10)
	result := Collect(Take(numbers, 3))
	expected := []int{0, 1, 2}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSkip(t *testing.T) {
	numbers := Range(0, 5)
	result := Collect(Skip(numbers, 2))
	expected := []int{2, 3, 4}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFilter(t *testing.T) {
	numbers := Range(0, 10)
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	result := Collect(evens)
	expected := []int{0, 2, 4, 6, 8}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMap(t *testing.T) {
	numbers := Range(1, 4)
	doubled := Map(numbers, func(x int) int { return x * 2 })
	result := Collect(doubled)
	expected := []int{2, 4, 6}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestChain(t *testing.T) {
	seq1 := Range(1, 3)
	seq2 := Range(10, 12)
	chained := Chain(seq1, seq2)
	result := Collect(chained)
	expected := []int{1, 2, 10, 11}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestEnumerate(t *testing.T) {
	words := Values([]string{"apple", "banana", "cherry"})
	result := Collect2(Enumerate(words))
	expected := map[int]string{0: "apple", 1: "banana", 2: "cherry"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestZip(t *testing.T) {
	seq1 := Range(1, 4)
	seq2 := Values([]string{"a", "b", "c"})
	zipped := Zip(seq1, seq2)
	
	var result []struct{ int; string }
	for x, y := range zipped {
		result = append(result, struct{ int; string }{x, y})
	}
	
	expected := []struct{ int; string }{{1, "a"}, {2, "b"}, {3, "c"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReduce(t *testing.T) {
	numbers := Range(1, 6)
	sum := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
	expected := 15 // 1+2+3+4+5
	
	if sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}
}

func TestFind(t *testing.T) {
	numbers := Range(1, 10)
	result, found := Find(numbers, func(x int) bool { return x > 5 })
	
	if !found {
		t.Error("Expected to find a number > 5")
	}
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestAll(t *testing.T) {
	positives := Range(1, 5)
	allPositive := All(positives, func(x int) bool { return x > 0 })
	
	if !allPositive {
		t.Error("Expected all numbers to be positive")
	}
	
	mixed := Values([]int{1, 2, -3, 4})
	allPositiveMixed := All(mixed, func(x int) bool { return x > 0 })
	
	if allPositiveMixed {
		t.Error("Expected not all numbers to be positive")
	}
}

func TestAny(t *testing.T) {
	numbers := Range(1, 10)
	hasEven := Any(numbers, func(x int) bool { return x%2 == 0 })
	
	if !hasEven {
		t.Error("Expected to find even numbers")
	}
	
	odds := []int{1, 3, 5, 7}
	hasEvenOdds := Any(Values(odds), func(x int) bool { return x%2 == 0 })
	
	if hasEvenOdds {
		t.Error("Expected no even numbers in odds slice")
	}
}

func TestContains(t *testing.T) {
	numbers := Range(1, 5)
	contains3 := Contains(numbers, 3)
	
	if !contains3 {
		t.Error("Expected to contain 3")
	}
	
	contains10 := Contains(Range(1, 5), 10)
	if contains10 {
		t.Error("Expected not to contain 10")
	}
}

func TestFirst(t *testing.T) {
	numbers := Range(5, 10)
	first, found := First(numbers)
	
	if !found {
		t.Error("Expected to find first element")
	}
	if first != 5 {
		t.Errorf("Expected 5, got %d", first)
	}
	
	// Test empty iterator
	empty := Range(0, 0)
	_, found = First(empty)
	if found {
		t.Error("Expected not to find element in empty iterator")
	}
}

func TestLast(t *testing.T) {
	numbers := Range(5, 10)
	last, found := Last(numbers)
	
	if !found {
		t.Error("Expected to find last element")
	}
	if last != 9 {
		t.Errorf("Expected 9, got %d", last)
	}
}

func TestNth(t *testing.T) {
	numbers := Range(10, 20)
	nth, found := Nth(numbers, 3)
	
	if !found {
		t.Error("Expected to find 3rd element")
	}
	if nth != 13 {
		t.Errorf("Expected 13, got %d", nth)
	}
	
	// Test out of bounds
	_, found = Nth(Range(0, 3), 5)
	if found {
		t.Error("Expected not to find element at index 5")
	}
}

func TestLen(t *testing.T) {
	numbers := Range(0, 10)
	length := Len(numbers)
	expected := 10
	
	if length != expected {
		t.Errorf("Expected %d, got %d", expected, length)
	}
}

func TestLines(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"line1\nline2\nline3", []string{"line1", "line2", "line3"}},
		{"single line", []string{"single line"}},
		{"", []string{}},
		{"line1\nline2\n", []string{"line1", "line2", ""}},
		{"\n\n", []string{"", "", ""}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Lines(%q)", test.input), func(t *testing.T) {
			result := Collect(Lines(test.input))
			if len(result) == 0 && len(test.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestSplitSeq(t *testing.T) {
	tests := []struct {
		input, sep string
		expected   []string
	}{
		{"a,b,c", ",", []string{"a", "b", "c"}},
		{"hello world", " ", []string{"hello", "world"}},
		{"abc", "", []string{"a", "b", "c"}},
		{"", ",", []string{""}},
		{"a,,c", ",", []string{"a", "", "c"}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("SplitSeq(%q,%q)", test.input, test.sep), func(t *testing.T) {
			result := Collect(SplitSeq(test.input, test.sep))
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestFieldsSeq(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"hello world", []string{"hello", "world"}},
		{"  hello   world  ", []string{"hello", "world"}},
		{"", []string{}},
		{"   ", []string{}},
		{"single", []string{"single"}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("FieldsSeq(%q)", test.input), func(t *testing.T) {
			result := Collect(FieldsSeq(test.input))
			if len(result) == 0 && len(test.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestFieldsFuncSeq(t *testing.T) {
	text := "hello,world;test"
	isPunctuation := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}
	
	result := Collect(FieldsFuncSeq(text, isPunctuation))
	expected := []string{"hello", "world", "test"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestEarlyTermination(t *testing.T) {
	// Test that iteration stops when yield returns false
	numbers := Range(0, 1000000) // Large range
	
	var collected []int
	count := 0
	for n := range numbers {
		collected = append(collected, n)
		count++
		if count >= 5 {
			break
		}
	}
	
	expected := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(collected, expected) {
		t.Errorf("Expected %v, got %v", expected, collected)
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for range Range(0, 1000) {
			// Consume iterator
		}
	}
}

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := Range(0, 1000)
		evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
		for range evens {
			// Consume iterator
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := Range(0, 1000)
		doubled := Map(numbers, func(x int) int { return x * 2 })
		for range doubled {
			// Consume iterator
		}
	}
}

func BenchmarkChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq1 := Range(0, 500)
		seq2 := Range(500, 1000)
		chained := Chain(seq1, seq2)
		for range chained {
			// Consume iterator
		}
	}
}

func BenchmarkCollect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := Range(0, 1000)
		_ = Collect(numbers)
	}
}

func BenchmarkLinesLarge(b *testing.B) {
	// Create a large text with many lines
	text := ""
	for i := 0; i < 1000; i++ {
		text += fmt.Sprintf("This is line %d\n", i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Lines(text) {
			// Consume iterator
		}
	}
}

func BenchmarkSplitSeqLarge(b *testing.B) {
	// Create a large CSV-like string
	text := ""
	for i := 0; i < 1000; i++ {
		if i > 0 {
			text += ","
		}
		text += fmt.Sprintf("field%d", i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range SplitSeq(text, ",") {
			// Consume iterator
		}
	}
}