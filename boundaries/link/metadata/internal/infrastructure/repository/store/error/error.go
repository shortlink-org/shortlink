package errors

// MetadataNotFoundByIdError - metadata not found by id
type MetadataNotFoundByIdError struct {
	ID string
}

func (e *MetadataNotFoundByIdError) Error() string {
	return "metadata not found by id: " + e.ID
}
