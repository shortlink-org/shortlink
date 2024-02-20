package es_postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventStore struct {
	db *pgxpool.Pool

	Aggregates
	Events
}
