package mysql

import (
	"database/sql"
)

// Config ...
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *sql.DB
	config Config
}
