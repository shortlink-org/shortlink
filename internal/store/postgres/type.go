package postgres

import "github.com/jackc/pgx/v4/pgxpool"

// PostgresConfig ...
type PostgresConfig struct { // nolint unused
	URI string
}

// PostgresLinkList implementation of store interface
type PostgresLinkList struct { // nolint unused
	client *pgxpool.Pool
	config PostgresConfig
}
