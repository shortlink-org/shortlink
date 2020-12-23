package mysql

import (
	"github.com/jmoiron/sqlx"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *sqlx.DB
}
