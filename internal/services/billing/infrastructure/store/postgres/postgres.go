package postgres

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
)

type Store struct {
	db *db.Store
}

func (s Store) Init(ctx context.Context, db *db.Store) error {
	s.db = db
	return nil
}
