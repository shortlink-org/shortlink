package mysql

import (
	"github.com/jmoiron/sqlx"
)

// Config ...
type Config struct { // nolint unused
	URI string
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *sqlx.DB
	config Config
}
