package dag

import (
	"fmt"
)

type VertexAlreadyExistsError struct {
	Id string
}

func (e *VertexAlreadyExistsError) Error() string {
	return fmt.Sprintf("dag already contains a vertex with the id: %s", e.Id)
}

type VertexNotFoundError struct {
	Id string
}

func (e *VertexNotFoundError) Error() string {
	return fmt.Sprintf("not found vertex by id: %s", e.Id)
}

var ErrIncorrectTypeAssertion = fmt.Errorf("incorrect type assertion")
