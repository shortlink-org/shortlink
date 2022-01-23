package tool

// TODO: add test
// Contains tells whether a contains x.
func Contains[V Type](a []V, x V) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// TODO: add test
// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find[V Type](a []V, x V) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// TODO: add test
// Returns unique items in a slice
func Unique[V Type](slice []V) []V {
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
