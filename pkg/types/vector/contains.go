package vector

// Contains tells whether a slice `a` contains the element `x`.
func Contains[V comparable](a []V, x V) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}
