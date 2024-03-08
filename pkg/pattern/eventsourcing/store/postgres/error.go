package es_postgres

import (
	"fmt"
)

type IncorrectUpdatedBillingError struct {
	Err     error
	Updated int64
}

func (e *IncorrectUpdatedBillingError) Error() string {
	return fmt.Sprintf(`incorrect updated aggregates. Updated: %d/1`, e.Updated)
}
