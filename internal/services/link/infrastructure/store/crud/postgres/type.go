package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
)

// Config ...
type Config struct {
	job  *batch.Config
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool

	config Config
}
