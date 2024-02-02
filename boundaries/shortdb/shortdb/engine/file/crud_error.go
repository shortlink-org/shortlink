package file

import (
	"fmt"
)

var ErrIncorrectNameFields = fmt.Errorf("at SELECT: expected field to SELECT")
var ErrCreatePage = fmt.Errorf("at INSERT INTO: error create a new page")
var ErrCreateCursor = fmt.Errorf("at INSERT INTO: error create a new cursor")

// NotExistTableError is an error type returned when the table does not exist
type NotExistTableError struct {
	Table string
	Type  string
}

func (e *NotExistTableError) Error() string {
	switch e.Type {
	case "SELECT":
		return fmt.Sprintf("at SELECT: not exist table %s", e.Table)
	case "INSERT":
		return fmt.Sprintf("at INSERT INTO: not exist table %s", e.Table)
	default:
		return fmt.Sprintf("not exist table %s", e.Table)
	}
}

// CreateCursorError is an error type returned when the cursor cannot be created
type CreateCursorError struct {
	Type string
}

func (e *CreateCursorError) Error() string {
	switch e.Type {
	case "SELECT":
		return "at SELECT: error create a new cursor"
	case "INSERT":
		return "at INSERT INTO: error create a new cursor"
	default:
		return "error create a new cursor"
	}
}

// IncorrectNameFieldsError is an error type returned when the name of the field is incorrect
type IncorrectNameFieldsError struct {
	Field string
	Table string
}

func (e *IncorrectNameFieldsError) Error() string {
	return fmt.Sprintf("at SELECT: incorrect name fields %s in table %s", e.Field, e.Table)
}

// IncorrectTypeFieldsError is an error type returned when the type of the field is incorrect
type IncorrectTypeFieldsError struct {
	Field string
	Table string
}

func (e *IncorrectTypeFieldsError) Error() string {
	return fmt.Sprintf("at INSERT INTO: incorrect type fields %s in table %s", e.Field, e.Table)
}
