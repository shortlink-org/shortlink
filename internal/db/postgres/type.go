package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// Config ...
type Config struct { // nolint unused
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *pgxpool.Pool
	config Config
}
