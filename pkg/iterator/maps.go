// Package iterator provides map iterator functions that extend the standard library.
package iterator

import (
	"cmp"
	"iter"
)

// All returns an iterator over key-value pairs for the map.
func AllMaps[M ~map[K]V, K comparable, V any](m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Keys returns an iterator over the keys in the map.
func Keys[M ~map[K]V, K comparable, V any](m M) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}

// Values returns an iterator over the values in the map.
func ValuesMaps[M ~map[K]V, K comparable, V any](m M) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

// SortedKeys returns an iterator over the keys in sorted order.
func SortedKeys[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq[K] {
	return func(yield func(K) bool) {
		// Collect keys
		keys := make([]K, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		
		// Sort keys
		if len(keys) <= 20 {
			for i := 1; i < len(keys); i++ {
				key := keys[i]
				j := i - 1
				for j >= 0 && keys[j] > key {
					keys[j+1] = keys[j]
					j--
				}
				keys[j+1] = key
			}
		} else {
			quickSortKeys(keys, 0, len(keys)-1)
		}
		
		// Yield sorted keys
		for _, k := range keys {
			if !yield(k) {
				return
			}
		}
	}
}

// quickSortKeys implements quicksort for keys.
func quickSortKeys[K cmp.Ordered](keys []K, low, high int) {
	if low < high {
		pi := partitionKeys(keys, low, high)
		quickSortKeys(keys, low, pi-1)
		quickSortKeys(keys, pi+1, high)
	}
}

// partitionKeys is a helper function for quicksort.
func partitionKeys[K cmp.Ordered](keys []K, low, high int) int {
	pivot := keys[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if keys[j] <= pivot {
			i++
			keys[i], keys[j] = keys[j], keys[i]
		}
	}
	keys[i+1], keys[high] = keys[high], keys[i+1]
	return i + 1
}

// SortedAll returns an iterator over key-value pairs sorted by key.
func SortedAll[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k := range SortedKeys(m) {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

// SortedKeysFunc returns an iterator over keys sorted by the given comparison function.
func SortedKeysFunc[M ~map[K]V, K comparable, V any](m M, cmp func(K, K) int) iter.Seq[K] {
	return func(yield func(K) bool) {
		// Collect keys
		keys := make([]K, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		
		// Sort keys using insertion sort
		for i := 1; i < len(keys); i++ {
			key := keys[i]
			j := i - 1
			for j >= 0 && cmp(keys[j], key) > 0 {
				keys[j+1] = keys[j]
				j--
			}
			keys[j+1] = key
		}
		
		// Yield sorted keys
		for _, k := range keys {
			if !yield(k) {
				return
			}
		}
	}
}

// SortedAllFunc returns an iterator over key-value pairs sorted by key using the given comparison function.
func SortedAllFunc[M ~map[K]V, K comparable, V any](m M, cmp func(K, K) int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k := range SortedKeysFunc(m, cmp) {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

// Filter returns an iterator over key-value pairs that satisfy the predicate.
func FilterMaps[M ~map[K]V, K comparable, V any](m M, predicate func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// FilterKeys returns an iterator over key-value pairs where keys satisfy the predicate.
func FilterKeys[M ~map[K]V, K comparable, V any](m M, predicate func(K) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if predicate(k) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// FilterValues returns an iterator over key-value pairs where values satisfy the predicate.
func FilterValues[M ~map[K]V, K comparable, V any](m M, predicate func(V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if predicate(v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// MapKeys returns an iterator that transforms keys using the given function.
func MapKeys[M ~map[K]V, K comparable, V any, NK comparable](m M, fn func(K) NK) iter.Seq2[NK, V] {
	return func(yield func(NK, V) bool) {
		for k, v := range m {
			if !yield(fn(k), v) {
				return
			}
		}
	}
}

// MapValues returns an iterator that transforms values using the given function.
func MapValues[M ~map[K]V, K comparable, V any, NV any](m M, fn func(V) NV) iter.Seq2[K, NV] {
	return func(yield func(K, NV) bool) {
		for k, v := range m {
			if !yield(k, fn(v)) {
				return
			}
		}
	}
}

// MapAll returns an iterator that transforms both keys and values using the given function.
func MapAll[M ~map[K]V, K comparable, V any, NK comparable, NV any](m M, fn func(K, V) (NK, NV)) iter.Seq2[NK, NV] {
	return func(yield func(NK, NV) bool) {
		for k, v := range m {
			nk, nv := fn(k, v)
			if !yield(nk, nv) {
				return
			}
		}
	}
}

// Invert returns an iterator over inverted key-value pairs (values become keys, keys become values).
func Invert[M ~map[K]V, K comparable, V comparable](m M) iter.Seq2[V, K] {
	return func(yield func(V, K) bool) {
		for k, v := range m {
			if !yield(v, k) {
				return
			}
		}
	}
}

// Merge returns an iterator over key-value pairs from multiple maps.
// If the same key exists in multiple maps, the value from the last map is used.
func Merge[M ~map[K]V, K comparable, V any](maps ...M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seen := make(map[K]V)
		
		// Collect all key-value pairs, with later maps overriding earlier ones
		for _, m := range maps {
			for k, v := range m {
				seen[k] = v
			}
		}
		
		// Yield the merged pairs
		for k, v := range seen {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Intersect returns an iterator over key-value pairs that exist in all maps.
func IntersectMaps[M ~map[K]V, K comparable, V comparable](maps ...M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if len(maps) == 0 {
			return
		}
		
		// Start with the first map
		candidates := make(map[K]V)
		for k, v := range maps[0] {
			candidates[k] = v
		}
		
		// Check intersection with each subsequent map
		for _, m := range maps[1:] {
			newCandidates := make(map[K]V)
			for k, v := range candidates {
				if mv, exists := m[k]; exists && mv == v {
					newCandidates[k] = v
				}
			}
			candidates = newCandidates
		}
		
		// Yield the intersection
		for k, v := range candidates {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Union returns an iterator over key-value pairs that exist in any of the maps.
func UnionMaps[M ~map[K]V, K comparable, V any](maps ...M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seen := make(map[K]struct{})
		
		for _, m := range maps {
			for k, v := range m {
				if _, exists := seen[k]; !exists {
					seen[k] = struct{}{}
					if !yield(k, v) {
						return
					}
				}
			}
		}
	}
}

// Difference returns an iterator over key-value pairs that exist in the first map but not in any of the others.
func DifferenceMaps[M ~map[K]V, K comparable, V any](first M, others ...M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		excluded := make(map[K]struct{})
		
		// Collect keys to exclude
		for _, m := range others {
			for k := range m {
				excluded[k] = struct{}{}
			}
		}
		
		// Yield pairs from the first map that are not excluded
		for k, v := range first {
			if _, exists := excluded[k]; !exists {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// GroupBy returns an iterator over groups of key-value pairs grouped by the result of the key function.
func GroupByMaps[M ~map[K]V, K comparable, V any, GK comparable](m M, keyFunc func(K, V) GK) iter.Seq2[GK, map[K]V] {
	return func(yield func(GK, map[K]V) bool) {
		groups := make(map[GK]map[K]V)
		
		for k, v := range m {
			gk := keyFunc(k, v)
			if groups[gk] == nil {
				groups[gk] = make(map[K]V)
			}
			groups[gk][k] = v
		}
		
		for gk, group := range groups {
			if !yield(gk, group) {
				return
			}
		}
	}
}

// PartitionBy returns two iterators: one for key-value pairs that satisfy the predicate,
// and one for those that don't.
func PartitionByMaps[M ~map[K]V, K comparable, V any](m M, predicate func(K, V) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	trueSeq := func(yield func(K, V) bool) {
		for k, v := range m {
			if predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
	
	falseSeq := func(yield func(K, V) bool) {
		for k, v := range m {
			if !predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
	
	return trueSeq, falseSeq
}

// ToSlice returns an iterator that converts map key-value pairs to a slice of pairs.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// ToSlice converts a map to an iterator over key-value pairs as structs.
func ToSlice[M ~map[K]V, K comparable, V any](m M) iter.Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for k, v := range m {
			if !yield(Pair[K, V]{Key: k, Value: v}) {
				return
			}
		}
	}
}

// Reduce applies a reduction function to all key-value pairs in the map.
func ReduceMaps[M ~map[K]V, K comparable, V any, Acc any](m M, initial Acc, fn func(Acc, K, V) Acc) Acc {
	acc := initial
	for k, v := range m {
		acc = fn(acc, k, v)
	}
	return acc
}

// FindKey returns the first key-value pair where the key satisfies the predicate.
func FindKey[M ~map[K]V, K comparable, V any](m M, predicate func(K) bool) (K, V, bool) {
	for k, v := range m {
		if predicate(k) {
			return k, v, true
		}
	}
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// FindValue returns the first key-value pair where the value satisfies the predicate.
func FindValue[M ~map[K]V, K comparable, V any](m M, predicate func(V) bool) (K, V, bool) {
	for k, v := range m {
		if predicate(v) {
			return k, v, true
		}
	}
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// Contains checks if the map contains a specific key-value pair.
func ContainsPair[M ~map[K]V, K comparable, V comparable](m M, key K, value V) bool {
	if v, exists := m[key]; exists && v == value {
		return true
	}
	return false
}

// ContainsKey checks if the map contains a specific key.
func ContainsKey[M ~map[K]V, K comparable, V any](m M, key K) bool {
	_, exists := m[key]
	return exists
}

// ContainsValue checks if the map contains a specific value.
func ContainsValue[M ~map[K]V, K comparable, V comparable](m M, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}