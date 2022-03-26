package query

// NotFoundError - not found link
type StoreError struct {
	Value string
}

func (e *StoreError) Error() string {
	return e.Value
}
