package error_di

type BaseError struct {
	Err error
}

func (e *BaseError) Error() string {
	return e.Err.Error()
}
