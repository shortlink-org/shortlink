package cache

// InitCacheError is an error returned when cache initialization fails.
type InitCacheError struct {
	err error
}

func (e *InitCacheError) Error() string {
	return "error init cache: " + e.err.Error()
}

// NewCacheError creates a new cache error.
func NewCacheError(op string, err error) error {
	return &BaseError{
		op:  op,
		err: err,
	}
}

// BaseError is an error returned by cache operations.
type BaseError struct {
	op  string
	err error
}

func (e *BaseError) Error() string {
	return "cache: " + e.op + ": " + e.err.Error()
}
