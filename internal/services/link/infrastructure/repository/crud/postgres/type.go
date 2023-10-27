package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/postgres/schema/crud"
)

// Config - config
type Config struct {
	job  *batch.Batch
	URI  string
	mode int
}

// Store implementation of db interface
type Store struct {
	client    *pgxpool.Pool
	newClient *crud.Queries

	config Config
}
