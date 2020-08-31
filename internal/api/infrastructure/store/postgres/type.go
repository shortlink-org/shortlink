package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/batch"
)

// Config ...
type Config struct { // nolint unused
	URI  string
	mode int
	job  *batch.Config
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *pgxpool.Pool
	config Config
}
