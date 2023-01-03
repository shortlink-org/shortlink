package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
)

// Config ...
type Config struct {
	URI  string
	mode int
	job  *batch.Config
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool

	config Config
}
