package vector

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
// It takes a slice of any comparable type and an element of the same type.
// If the element is found, it returns the index of the first occurrence.
// If the element is not found, it returns the length of the slice.
func Find[V comparable](a []V, x V) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}

	return len(a)
}
