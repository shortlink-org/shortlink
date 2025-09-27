// Package iterator provides slice iterator functions that extend the standard library.
package iterator

import (
	"cmp"
	"iter"
)

// AllSlices returns an iterator over index-value pairs for the slice.
func AllSlices[S ~[]E, E any](s S) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

// Backwards returns an iterator over the slice in reverse order.
func Backwards[S ~[]E, E any](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// BackwardsAll returns an iterator over index-value pairs for the slice in reverse order.
func BackwardsAll[S ~[]E, E any](s S) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

// ValuesSlices returns an iterator over the values in the slice.
func ValuesSlices[S ~[]E, E any](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

// Chunk returns an iterator that yields chunks of the slice of the given size.
func Chunk[S ~[]E, E any](s S, size int) iter.Seq[S] {
	return func(yield func(S) bool) {
		if size <= 0 {
			return
		}
		
		for i := 0; i < len(s); i += size {
			end := i + size
			if end > len(s) {
				end = len(s)
			}
			if !yield(s[i:end]) {
				return
			}
		}
	}
}

// Window returns an iterator that yields sliding windows of the slice of the given size.
func Window[S ~[]E, E any](s S, size int) iter.Seq[S] {
	return func(yield func(S) bool) {
		if size <= 0 || size > len(s) {
			return
		}
		
		for i := 0; i <= len(s)-size; i++ {
			if !yield(s[i : i+size]) {
				return
			}
		}
	}
}

// Enumerate returns an iterator over index-value pairs starting from a given index.
func EnumerateFrom[S ~[]E, E any](s S, start int) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, v := range s {
			if !yield(start+i, v) {
				return
			}
		}
	}
}

// StepBy returns an iterator that yields every nth element of the slice.
func StepBy[S ~[]E, E any](s S, step int) iter.Seq[E] {
	return func(yield func(E) bool) {
		if step <= 0 {
			return
		}
		
		for i := 0; i < len(s); i += step {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// TakeWhile returns an iterator that yields elements while they satisfy the predicate.
func TakeWhile[S ~[]E, E any](s S, predicate func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range s {
			if !predicate(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// SkipWhile returns an iterator that skips elements while they satisfy the predicate.
func SkipWhile[S ~[]E, E any](s S, predicate func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		skipping := true
		for _, v := range s {
			if skipping && predicate(v) {
				continue
			}
			skipping = false
			if !yield(v) {
				return
			}
		}
	}
}

// Flatten returns an iterator that flattens a slice of slices.
func Flatten[S ~[][]E, E any](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, inner := range s {
			for _, v := range inner {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Unique returns an iterator that yields unique elements from the slice.
func Unique[S ~[]E, E comparable](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		seen := make(map[E]struct{})
		for _, v := range s {
			if _, exists := seen[v]; !exists {
				seen[v] = struct{}{}
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Dedup returns an iterator that yields elements, removing consecutive duplicates.
func Dedup[S ~[]E, E comparable](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		if len(s) == 0 {
			return
		}
		
		if !yield(s[0]) {
			return
		}
		
		for i := 1; i < len(s); i++ {
			if s[i] != s[i-1] {
				if !yield(s[i]) {
					return
				}
			}
		}
	}
}

// Intersect returns an iterator over elements that are in both slices.
func Intersect[S ~[]E, E comparable](s1, s2 S) iter.Seq[E] {
	return func(yield func(E) bool) {
		set2 := make(map[E]struct{})
		for _, v := range s2 {
			set2[v] = struct{}{}
		}
		
		seen := make(map[E]struct{})
		for _, v := range s1 {
			if _, exists := set2[v]; exists {
				if _, alreadySeen := seen[v]; !alreadySeen {
					seen[v] = struct{}{}
					if !yield(v) {
						return
					}
				}
			}
		}
	}
}

// Union returns an iterator over elements that are in either slice.
func Union[S ~[]E, E comparable](s1, s2 S) iter.Seq[E] {
	return func(yield func(E) bool) {
		seen := make(map[E]struct{})
		
		for _, v := range s1 {
			if _, exists := seen[v]; !exists {
				seen[v] = struct{}{}
				if !yield(v) {
					return
				}
			}
		}
		
		for _, v := range s2 {
			if _, exists := seen[v]; !exists {
				seen[v] = struct{}{}
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Difference returns an iterator over elements that are in s1 but not in s2.
func Difference[S ~[]E, E comparable](s1, s2 S) iter.Seq[E] {
	return func(yield func(E) bool) {
		set2 := make(map[E]struct{})
		for _, v := range s2 {
			set2[v] = struct{}{}
		}
		
		seen := make(map[E]struct{})
		for _, v := range s1 {
			if _, exists := set2[v]; !exists {
				if _, alreadySeen := seen[v]; !alreadySeen {
					seen[v] = struct{}{}
					if !yield(v) {
						return
					}
				}
			}
		}
	}
}

// Sorted returns an iterator over the slice elements in sorted order.
// This creates a copy of the slice to avoid modifying the original.
func Sorted[S ~[]E, E cmp.Ordered](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		// Create a copy to avoid modifying the original
		sorted := make(S, len(s))
		copy(sorted, s)
		
		// Simple insertion sort for small slices, otherwise use a more complex sort
		if len(sorted) <= 20 {
			for i := 1; i < len(sorted); i++ {
				key := sorted[i]
				j := i - 1
				for j >= 0 && sorted[j] > key {
					sorted[j+1] = sorted[j]
					j--
				}
				sorted[j+1] = key
			}
		} else {
			// Use quicksort for larger slices
			quickSort(sorted, 0, len(sorted)-1)
		}
		
		for _, v := range sorted {
			if !yield(v) {
				return
			}
		}
	}
}

// quickSort implements quicksort for ordered types.
func quickSort[S ~[]E, E cmp.Ordered](s S, low, high int) {
	if low < high {
		pi := partition(s, low, high)
		quickSort(s, low, pi-1)
		quickSort(s, pi+1, high)
	}
}

// partition is a helper function for quicksort.
func partition[S ~[]E, E cmp.Ordered](s S, low, high int) int {
	pivot := s[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if s[j] <= pivot {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}
	s[i+1], s[high] = s[high], s[i+1]
	return i + 1
}

// SortedFunc returns an iterator over the slice elements sorted by the given comparison function.
func SortedFunc[S ~[]E, E any](s S, cmp func(E, E) int) iter.Seq[E] {
	return func(yield func(E) bool) {
		// Create a copy to avoid modifying the original
		sorted := make(S, len(s))
		copy(sorted, s)
		
		// Simple insertion sort
		for i := 1; i < len(sorted); i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && cmp(sorted[j], key) > 0 {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		
		for _, v := range sorted {
			if !yield(v) {
				return
			}
		}
	}
}

// GroupBy returns an iterator over groups of consecutive elements that have the same key.
func GroupBy[S ~[]E, E any, K comparable](s S, keyFunc func(E) K) iter.Seq2[K, S] {
	return func(yield func(K, S) bool) {
		if len(s) == 0 {
			return
		}
		
		start := 0
		currentKey := keyFunc(s[0])
		
		for i := 1; i < len(s); i++ {
			key := keyFunc(s[i])
			if key != currentKey {
				if !yield(currentKey, s[start:i]) {
					return
				}
				start = i
				currentKey = key
			}
		}
		
		// Yield the last group
		yield(currentKey, s[start:])
	}
}

// PartitionBy returns an iterator over partitions of the slice where elements are grouped by the result of the key function.
func PartitionBy[S ~[]E, E any, K comparable](s S, keyFunc func(E) K) iter.Seq2[K, S] {
	return func(yield func(K, S) bool) {
		groups := make(map[K][]E)
		
		for _, v := range s {
			key := keyFunc(v)
			groups[key] = append(groups[key], v)
		}
		
		for key, group := range groups {
			if !yield(key, S(group)) {
				return
			}
		}
	}
}

// Scan returns an iterator that yields the intermediate results of applying a function to each element.
func Scan[S ~[]E, E any, Acc any](s S, initial Acc, fn func(Acc, E) Acc) iter.Seq[Acc] {
	return func(yield func(Acc) bool) {
		acc := initial
		if !yield(acc) {
			return
		}
		
		for _, v := range s {
			acc = fn(acc, v)
			if !yield(acc) {
				return
			}
		}
	}
}

// Intersperse returns an iterator that yields elements with a separator between them.
func Intersperse[S ~[]E, E any](s S, separator E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i, v := range s {
			if i > 0 {
				if !yield(separator) {
					return
				}
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Pairwise returns an iterator over pairs of consecutive elements.
func Pairwise[S ~[]E, E any](s S) iter.Seq2[E, E] {
	return func(yield func(E, E) bool) {
		for i := 0; i < len(s)-1; i++ {
			if !yield(s[i], s[i+1]) {
				return
			}
		}
	}
}