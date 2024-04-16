package main

import (
	"fmt"
)

type NotSupportDatabaseError struct {
	DbType string
}

func (e NotSupportDatabaseError) Error() string {
	return fmt.Sprintf("database type %s is not supported", e.DbType)
}
