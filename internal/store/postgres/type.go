package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/batch"
)

// PostgresConfig ...
type PostgresConfig struct { // nolint unused
	URI  string
	mode int
	job  *batch.Config
}

// PostgresLinkList implementation of store interface
type PostgresLinkList struct { // nolint unused
	client *pgxpool.Pool
	config PostgresConfig
}
