package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/repository/crud/postgres/schema/crud"
	"github.com/shortlink-org/shortlink/internal/pkg/batch"
)

// Config - config
type Config struct {
	job  *batch.Batch
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client *pgxpool.Pool
	query  *crud.Queries

	config Config
}
