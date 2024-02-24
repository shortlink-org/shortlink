package migrate

import (
	"fmt"
)

type MigrationError struct {
	error

	Table string
}

func (e *MigrationError) Error() string {
	return fmt.Sprintf("Error in migration for table %s: %s", e.Table, e.error.Error())
}
