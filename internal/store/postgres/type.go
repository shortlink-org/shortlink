package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgresConfig ...
type PostgresConfig struct { // nolint unused
	URI string
}

// PostgresLinkList implementation of store interface
type PostgresLinkList struct { // nolint unused
	ctx context.Context

	client *pgxpool.Pool
	config PostgresConfig
}
