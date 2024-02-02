package postgres

import (
	"fmt"
)

// PingConnectionError - error ping connection
type PingConnectionError struct {
	Err error
}

func (e *PingConnectionError) Error() string {
	return fmt.Sprintf("failed to ping the database: %s", e.Err.Error())
}
