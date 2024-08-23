package vector

// Unique returns unique items in a slice.
// It takes a slice of any comparable type and returns a new slice containing only the unique elements from the original slice.
// The order of elements in the returned slice is not guaranteed to be the same as in the original slice.
func Unique[V comparable](slice []V) []V {
	// create a map with all the values as key
	uniqMap := make(map[V]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]V, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}

	return uniqSlice
}
