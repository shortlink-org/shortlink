package postgres

// PingConnectionError - error ping connection
type PingConnectionError struct {
	Err error
}

func (e *PingConnectionError) Error() string {
	return "failed to ping the database: " + e.Err.Error()
}
