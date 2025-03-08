package migrate

type MigrationError struct {
	Err         error
	Description string
}

func (e *MigrationError) Error() string {
	return "migration error: " + e.Description + ": " + e.Err.Error()
}

func (e *MigrationError) Unwrap() error {
	return e.Err
}
