package postgres

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config ...
type Config struct {
	configConnect *pgx.ConnConfig
	mode          int
	poolConfig    *pgxpool.Config
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool
	config *Config

	tracer Tracer
}
