// Package iterator provides support for range-over-func iterators as introduced in Go 1.23/1.24.
// This package contains the core iterator types and utility functions that work with
// the new range-over-func feature.
package iterator

import "iter"

// Seq is an iterator over sequences of individual values.
// When called, it returns a function that can be used in a for-range loop.
// Seq[V] is equivalent to iter.Seq[V] from Go 1.23+.
type Seq[V any] iter.Seq[V]

// Seq2 is an iterator over sequences of pairs of values, most commonly key-value pairs.
// When called, it returns a function that can be used in a for-range loop.
// Seq2[K, V] is equivalent to iter.Seq2[K, V] from Go 1.23+.
type Seq2[K, V any] iter.Seq2[K, V]

// Pull converts a push-style iterator to a pull-style iterator.
// It returns a function that can be called to get the next value and a stop function.
func Pull[V any](seq iter.Seq[V]) (next func() (V, bool), stop func()) {
	return iter.Pull(seq)
}

// Pull2 converts a push-style iterator to a pull-style iterator for pairs.
// It returns a function that can be called to get the next pair and a stop function.
func Pull2[K, V any](seq iter.Seq2[K, V]) (next func() (K, V, bool), stop func()) {
	return iter.Pull2(seq)
}

// Empty returns an empty iterator.
func Empty[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {}
}

// Empty2 returns an empty iterator for pairs.
func Empty2[K, V any]() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {}
}

// Single returns an iterator that yields a single value.
func Single[V any](v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		yield(v)
	}
}

// Single2 returns an iterator that yields a single key-value pair.
func Single2[K, V any](k K, v V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		yield(k, v)
	}
}

// Range returns an iterator over a range of integers from start to end (exclusive).
func Range(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Count returns an iterator that yields integers starting from 0.
func Count() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// CountFrom returns an iterator that yields integers starting from start.
func CountFrom(start int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Repeat returns an iterator that yields the same value repeatedly.
func Repeat[V any](v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			if !yield(v) {
				return
			}
		}
	}
}

// Take returns an iterator that yields at most n values from the input iterator.
func Take[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		count := 0
		for v := range seq {
			if count >= n {
				return
			}
			if !yield(v) {
				return
			}
			count++
		}
	}
}

// Skip returns an iterator that skips the first n values from the input iterator.
func Skip[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		count := 0
		for v := range seq {
			if count < n {
				count++
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Filter returns an iterator that yields only values that satisfy the predicate.
func Filter[V any](seq iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Map returns an iterator that applies a function to each value.
func Map[V, U any](seq iter.Seq[V], fn func(V) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

// Enumerate returns an iterator that yields pairs of index and value.
func Enumerate[V any](seq iter.Seq[V]) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := 0
		for v := range seq {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// Zip returns an iterator that yields pairs from two iterators.
func Zip[V, U any](seq1 iter.Seq[V], seq2 iter.Seq[U]) iter.Seq2[V, U] {
	return func(yield func(V, U) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()
		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if !ok1 || !ok2 {
				return
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}

// Chain concatenates multiple iterators.
func Chain[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Collect gathers all values from an iterator into a slice.
func Collect[V any](seq iter.Seq[V]) []V {
	var result []V
	for v := range seq {
		result = append(result, v)
	}
	return result
}

// Collect2 gathers all key-value pairs from an iterator into a map.
func Collect2[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	result := make(map[K]V)
	for k, v := range seq {
		result[k] = v
	}
	return result
}

// Reduce applies a reduction function to all values in the iterator.
func Reduce[V, Acc any](seq iter.Seq[V], initial Acc, fn func(Acc, V) Acc) Acc {
	acc := initial
	for v := range seq {
		acc = fn(acc, v)
	}
	return acc
}

// Find returns the first value that satisfies the predicate.
func Find[V any](seq iter.Seq[V], predicate func(V) bool) (V, bool) {
	for v := range seq {
		if predicate(v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// All returns true if all values satisfy the predicate.
func All[V any](seq iter.Seq[V], predicate func(V) bool) bool {
	for v := range seq {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Any returns true if any value satisfies the predicate.
func Any[V any](seq iter.Seq[V], predicate func(V) bool) bool {
	for v := range seq {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Contains returns true if the iterator contains the given value.
func Contains[V comparable](seq iter.Seq[V], target V) bool {
	for v := range seq {
		if v == target {
			return true
		}
	}
	return false
}

// First returns the first value from the iterator.
func First[V any](seq iter.Seq[V]) (V, bool) {
	for v := range seq {
		return v, true
	}
	var zero V
	return zero, false
}

// Last returns the last value from the iterator.
func Last[V any](seq iter.Seq[V]) (V, bool) {
	var last V
	found := false
	for v := range seq {
		last = v
		found = true
	}
	return last, found
}

// Nth returns the nth value from the iterator (0-indexed).
func Nth[V any](seq iter.Seq[V], n int) (V, bool) {
	i := 0
	for v := range seq {
		if i == n {
			return v, true
		}
		i++
	}
	var zero V
	return zero, false
}

// Len returns the number of values in the iterator.
func Len[V any](seq iter.Seq[V]) int {
	count := 0
	for range seq {
		count++
	}
	return count
}

// Values returns an iterator over the values in the slice.
func Values[S ~[]E, E any](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}