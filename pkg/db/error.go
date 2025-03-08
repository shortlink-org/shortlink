package db

import (
	"errors"
)

// ErrGetConnection - error gets connection
var ErrGetConnection = errors.New("error get connection")

// UnknownStoreTypeError - unknown store type error
type UnknownStoreTypeError struct {
	StoreType string
}

func (e UnknownStoreTypeError) Error() string {
	return "unknown store type: " + e.StoreType
}
