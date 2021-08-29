package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

// New ...
func New(_ context.Context, db *db.Store) (*Store, error) {
	s := &Store{}

	// Set configuration
	s.client = db.Store.GetConn().(*pgxpool.Pool)

	return s, nil
}
