package query

// NotFoundError - not found link
type StoreError struct { // nolint unused
	Value string
}

func (e *StoreError) Error() string {
	return e.Value
}
