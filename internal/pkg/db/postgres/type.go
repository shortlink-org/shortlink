package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config ...
type Config struct {
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool
	config Config

	tracer Tracer
}
