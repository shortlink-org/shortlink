package postgres

import (
	"fmt"
)

// ErrorPingConnection - error ping connection
type ErrorPingConnection struct {
	Err error
}

func (e *ErrorPingConnection) Error() string {
	return fmt.Sprintf("failed to ping the database: %s", e.Err.Error())
}
