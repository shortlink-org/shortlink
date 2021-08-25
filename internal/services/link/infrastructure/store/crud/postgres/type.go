package postgres

import (
	"github.com/go-redis/cache/v8"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/batch"
	"github.com/batazor/shortlink/internal/pkg/logger"
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
	cache  *cache.Cache
	log    logger.Logger

	config Config
}
