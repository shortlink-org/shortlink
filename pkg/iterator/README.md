# Iterator Package

A comprehensive implementation of Go 1.24 iterator functionality, providing powerful tools for working with sequences of data using the new range-over-func feature.

## Overview

This package implements the iterator functionality introduced in Go 1.23/1.24, including:

- Core iterator types (`Seq`, `Seq2`) and utilities
- String processing iterators (similar to Go 1.24 `strings` package additions)
- Byte slice processing iterators
- Slice manipulation iterators
- Map processing iterators
- Functional programming utilities

## Features

### Core Iterator Functions

- `Range(start, end)` - Generate sequences of integers
- `Take(seq, n)` - Limit iterator to first n elements
- `Skip(seq, n)` - Skip first n elements
- `Filter(seq, predicate)` - Filter elements by predicate
- `Map(seq, fn)` - Transform elements
- `Chain(seqs...)` - Concatenate multiple iterators
- `Zip(seq1, seq2)` - Combine two iterators into pairs
- `Collect(seq)` - Gather all elements into a slice

### String Processing

Based on Go 1.24 additions to the `strings` package:

- `Lines(s)` - Split string into lines
- `SplitSeq(s, sep)` - Split by separator
- `SplitAfterSeq(s, sep)` - Split after separator
- `FieldsSeq(s)` - Split by whitespace
- `FieldsFuncSeq(s, f)` - Split by custom predicate

### Slice Operations

- `All(slice)` - Iterate over index-value pairs
- `Backwards(slice)` - Reverse iteration
- `Chunk(slice, size)` - Split into chunks
- `Window(slice, size)` - Sliding windows
- `Unique(slice)` - Remove duplicates
- `Sorted(slice)` - Sorted iteration
- `GroupBy(slice, keyFunc)` - Group consecutive elements

### Map Operations

- `Keys(map)` - Iterate over keys
- `Values(map)` - Iterate over values
- `SortedKeys(map)` - Iterate over sorted keys
- `Filter(map, predicate)` - Filter key-value pairs
- `Merge(maps...)` - Combine multiple maps

## Basic Usage

```go
package main

import (
    "fmt"
    "github.com/shortlink-org/shortlink/pkg/iterator"
)

func main() {
    // Basic range iteration
    for i := range iterator.Range(0, 5) {
        fmt.Print(i, " ") // Output: 0 1 2 3 4
    }
    
    // String processing
    logData := "INFO: Started\nWARN: High memory\nERROR: Failed"
    for line := range iterator.Lines(logData) {
        fmt.Println("Log:", line)
    }
    
    // Functional pipeline
    numbers := iterator.Range(1, 11)                                    // 1-10
    evens := iterator.Filter(numbers, func(x int) bool { return x%2 == 0 }) // 2,4,6,8,10
    squared := iterator.Map(evens, func(x int) int { return x * x })         // 4,16,36,64,100
    first3 := iterator.Take(squared, 3)                                     // 4,16,36
    result := iterator.Collect(first3)                                     // [4,16,36]
    
    fmt.Println("Result:", result)
}
```

## Advanced Examples

### Text Processing

```go
text := "The quick brown fox jumps over the lazy dog"

// Get words longer than 4 characters, uppercase
longWords := iterator.Filter(
    iterator.FieldsSeq(text),
    func(word string) bool { return len(word) > 4 },
)
upperWords := iterator.Map(longWords, strings.ToUpper)

for word := range upperWords {
    fmt.Println(word) // QUICK, BROWN, JUMPS
}
```

### Data Analysis

```go
data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Calculate sum of squares of even numbers
evens := iterator.Filter(iterator.Values(data), func(x int) bool { return x%2 == 0 })
squares := iterator.Map(evens, func(x int) int { return x * x })
sum := iterator.Reduce(squares, 0, func(acc, x int) int { return acc + x })

fmt.Println("Sum of squares of evens:", sum) // 220 (4+16+36+64+100)
```

### Map Processing

```go
scores := map[string]int{
    "Alice": 95,
    "Bob": 87,
    "Charlie": 92,
    "Diana": 88,
}

// Get high scorers (>90) in sorted order
highScorers := iterator.Filter(iterator.SortedAll(scores), 
    func(name string, score int) bool { return score > 90 })

for name, score := range highScorers {
    fmt.Printf("%s: %d\n", name, score)
}
```

### Chunking and Windows

```go
data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Process in chunks of 3
for chunk := range iterator.Chunk(data, 3) {
    sum := iterator.Reduce(iterator.Values(chunk), 0, func(a, b int) int { return a + b })
    fmt.Printf("Chunk %v sum: %d\n", chunk, sum)
}

// Sliding window of size 3
for window := range iterator.Window(data, 3) {
    avg := float64(iterator.Reduce(iterator.Values(window), 0, func(a, b int) int { return a + b })) / 3
    fmt.Printf("Window %v avg: %.2f\n", window, avg)
}
```

## Performance

The iterator approach provides several benefits:

1. **Memory Efficiency**: No intermediate slice allocations
2. **Lazy Evaluation**: Elements processed on-demand
3. **Early Termination**: Stop processing when needed
4. **Composability**: Chain operations efficiently

### Benchmarks

```go
// Traditional approach - allocates intermediate slices
func processTraditional(data []int) []int {
    var evens []int
    for _, x := range data {
        if x%2 == 0 {
            evens = append(evens, x)
        }
    }
    
    var squares []int
    for _, x := range evens {
        squares = append(squares, x*x)
    }
    
    if len(squares) > 5 {
        squares = squares[:5]
    }
    
    return squares
}

// Iterator approach - no intermediate allocations
func processIterator(data []int) []int {
    return iterator.Collect(
        iterator.Take(
            iterator.Map(
                iterator.Filter(iterator.Values(data), func(x int) bool { return x%2 == 0 }),
                func(x int) int { return x * x },
            ),
            5,
        ),
    )
}
```

## Error Handling

Iterators handle edge cases gracefully:

```go
// Empty iterators
empty := iterator.Range(0, 0)
fmt.Println("Length:", iterator.Len(empty)) // 0

// Out of bounds
data := []int{1, 2, 3}
value, found := iterator.Nth(iterator.Values(data), 10)
fmt.Println("Found:", found) // false

// Early termination
for i, v := range iterator.All(largeSlice) {
    if i >= 5 {
        break // Iterator stops processing
    }
    fmt.Println(v)
}
```

## Integration with Go 1.24

This package is designed to be compatible with Go 1.24's iterator additions:

```go
// When Go 1.24 is available, you can use:
import "strings" // for strings.Lines, strings.SplitSeq, etc.

// This package provides:
import "github.com/shortlink-org/shortlink/pkg/iterator" // for extended functionality
```

## Testing

Run the comprehensive test suite:

```bash
go test ./pkg/iterator -v
go test ./pkg/iterator -bench=.
```

## Contributing

When adding new iterator functions:

1. Follow the established patterns
2. Handle empty inputs gracefully
3. Support early termination (respect `yield` return value)
4. Add comprehensive tests
5. Include benchmarks for performance-critical functions
6. Document with examples

## License

This package is part of the shortlink project and follows the same licensing terms.