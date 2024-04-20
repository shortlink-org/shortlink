package errors

import (
	"fmt"
)

// MetadataNotFoundByIdError - metadata not found by id
type MetadataNotFoundByIdError struct {
	ID string
}

func (e *MetadataNotFoundByIdError) Error() string {
	return fmt.Sprintf("metadata not found by id: %s", e.ID)
}
