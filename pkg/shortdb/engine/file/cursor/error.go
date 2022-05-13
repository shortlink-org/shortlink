package cursor

import (
	"fmt"
)

type ErrorGetPage struct{}

func (e *ErrorGetPage) Error() string {
	return fmt.Sprintf("not found page")
}
