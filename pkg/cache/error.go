package cache

// InitCacheError is an error returned when cache initialization fails.
type InitCacheError struct {
	err error
}

func (e *InitCacheError) Error() string {
	return "error init cache: " + e.err.Error()
}
