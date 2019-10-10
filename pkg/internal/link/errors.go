package link

import (
	"fmt"
)

type NotFoundError struct {
	Link Link
	Err  error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found link: %s", e.Link.Url)
}
