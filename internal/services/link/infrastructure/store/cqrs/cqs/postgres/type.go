package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// Config ...
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool
}
