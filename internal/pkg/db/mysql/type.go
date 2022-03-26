package mysql

import (
	"github.com/jmoiron/sqlx"
)

// Config ...
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *sqlx.DB
	config Config
}
