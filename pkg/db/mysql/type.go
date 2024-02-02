package mysql

import (
	"database/sql"
)

// Config - configuration
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *sql.DB
	config Config
}
