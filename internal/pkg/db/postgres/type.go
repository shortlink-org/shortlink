package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config ...
type Config struct {
	mode   int
	config *pgxpool.Config
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool
	config *Config

	tracer Tracer
}
