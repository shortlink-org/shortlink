package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config - config
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool
}
