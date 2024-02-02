package file

import (
	"fmt"
)

// CreateExistIndexError is an error type returned when the index already exists
type CreateExistIndexError struct {
	Name string
}

func (e *CreateExistIndexError) Error() string {
	return fmt.Sprintf("at CREATE INDEX: exist index %s", e.Name)
}
