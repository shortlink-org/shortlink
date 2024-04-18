package link_cqrs

import (
	"fmt"
)

var (
	ErrCreateLink = fmt.Errorf("error create a new link")
	ErrNotFound   = fmt.Errorf("error not found link")
)
