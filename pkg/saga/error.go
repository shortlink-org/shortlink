package saga

import (
	"fmt"
)

type DublicateStepError struct {
	Name string
}

func (e *DublicateStepError) Error() string {
	return fmt.Sprintf("dublicate step: %s", e.Name)
}
