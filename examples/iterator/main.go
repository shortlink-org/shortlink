package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shortlink-org/shortlink/pkg/iterator"
)

func main() {
	fmt.Println("=== Go 1.24 Iterator Demo ===\n")

	// Basic range iteration
	fmt.Println("1. Basic Range Iterator:")
	for i := range iterator.Range(1, 6) {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// String processing (Go 1.24 style)
	fmt.Println("2. String Line Processing:")
	logData := "INFO: Server started\nWARN: High memory usage\nERROR: Database connection failed"
	for line := range iterator.Lines(logData) {
		fmt.Printf("  Log: %s\n", line)
	}
	fmt.Println()

	// CSV-like processing
	fmt.Println("3. CSV Processing:")
	csvData := "apple,banana,cherry,date"
	for fruit := range iterator.SplitSeq(csvData, ",") {
		fmt.Printf("  Fruit: %s\n", fruit)
	}
	fmt.Println()

	// Functional pipeline
	fmt.Println("4. Functional Pipeline:")
	numbers := iterator.Range(1, 21)                                       // 1-20
	evens := iterator.Filter(numbers, func(x int) bool { return x%2 == 0 }) // even numbers
	squared := iterator.Map(evens, func(x int) int { return x * x })        // square them
	first5 := iterator.Take(squared, 5)                                     // first 5
	result := iterator.Collect(first5)                                      // collect to slice
	
	fmt.Printf("  First 5 squares of even numbers: %v\n", result)
	fmt.Println()

	// Slice operations
	fmt.Println("5. Slice Operations:")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	fmt.Print("  Backwards: ")
	for v := range iterator.Backwards(data) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	
	fmt.Print("  Chunks of 3: ")
	for chunk := range iterator.Chunk(data, 3) {
		fmt.Printf("%v ", chunk)
	}
	fmt.Println()
	
	fmt.Print("  Window of 3: ")
	for window := range iterator.Window(data, 3) {
		fmt.Printf("%v ", window)
	}
	fmt.Println("\n")

	// Map operations
	fmt.Println("6. Map Operations:")
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
		"Diana":   88,
	}
	
	fmt.Println("  Sorted by name:")
	for name := range iterator.SortedKeys(scores) {
		fmt.Printf("    %s: %d\n", name, scores[name])
	}
	
	fmt.Println("  High scorers (>90):")
	highScorers := iterator.FilterMaps(scores, func(name string, score int) bool {
		return score > 90
	})
	for name, score := range highScorers {
		fmt.Printf("    %s: %d\n", name, score)
	}
	fmt.Println()

	// Text analysis
	fmt.Println("7. Text Analysis:")
	text := "The quick brown fox jumps over the lazy dog"
	words := iterator.FieldsSeq(text)
	longWords := iterator.Filter(words, func(word string) bool { return len(word) > 4 })
	upperWords := iterator.Map(longWords, strings.ToUpper)
	
	fmt.Print("  Long words (>4 chars): ")
	for word := range upperWords {
		fmt.Printf("%s ", word)
	}
	fmt.Println("\n")

	// Data aggregation
	fmt.Println("8. Data Aggregation:")
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	sum := iterator.Reduce(iterator.Values(values), 0, func(acc, x int) int {
		return acc + x
	})
	
	product := iterator.Reduce(
		iterator.Filter(iterator.Values(values), func(x int) bool { return x%2 == 0 }),
		1,
		func(acc, x int) int { return acc * x },
	)
	
	fmt.Printf("  Sum of all: %d\n", sum)
	fmt.Printf("  Product of evens: %d\n", product)
	fmt.Println()

	// Finding operations
	fmt.Println("9. Finding Operations:")
	testNumbers := []int{1, 3, 5, 8, 9, 12, 15}
	
	if even, found := iterator.Find(iterator.Values(testNumbers), func(x int) bool { return x%2 == 0 }); found {
		fmt.Printf("  First even number: %d\n", even)
	}
	
	allPositive := iterator.All(iterator.Values(testNumbers), func(x int) bool { return x > 0 })
	fmt.Printf("  All positive: %t\n", allPositive)
	
	anyLarge := iterator.Any(iterator.Values(testNumbers), func(x int) bool { return x > 10 })
	fmt.Printf("  Any > 10: %t\n", anyLarge)
	fmt.Println()

	// Number processing from strings
	fmt.Println("10. Error-Safe Number Processing:")
	numberStrings := []string{"1", "2", "3", "invalid", "4", "5"}
	
	validNumbers := iterator.Map(
		iterator.Filter(
			iterator.Values(numberStrings),
			func(s string) bool {
				_, err := strconv.Atoi(s)
				return err == nil
			},
		),
		func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		},
	)
	
	fmt.Print("  Valid numbers: ")
	for n := range validNumbers {
		fmt.Printf("%d ", n)
	}
	fmt.Println("\n")

	fmt.Println("=== Demo Complete ===")
}