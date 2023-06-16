package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

// New ...
func New(_ context.Context, db *db.Store) (*Store, error) {
	var ok bool
	s := &Store{}

	// Set configuration
	s.client, ok = db.Store.GetConn().(*pgxpool.Pool)
	if !ok {
		return nil, fmt.Errorf("error get connection")
	}

	return s, nil
}
