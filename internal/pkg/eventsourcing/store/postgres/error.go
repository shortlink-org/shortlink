package es_postgres

import (
	"fmt"
)

// Return fmt.Errorf(`incorrect updated billing.aggregates. Updated: %d/1`, row.RowsAffected())

type IncorrectUpdatedBillingError struct {
	Err     error
	Updated int64
}

func (e *IncorrectUpdatedBillingError) Error() string {
	return fmt.Sprintf(`incorrect updated billing.aggregates. Updated: %d/1`, e.Updated)
}
