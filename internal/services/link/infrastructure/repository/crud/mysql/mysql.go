//go:generate go run entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema

package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
)

// New store
func New(_ context.Context, store *db.Store) (*Store, error) {
	client, ok := store.Store.GetConn().(*sql.DB)
	if !ok {
		return nil, errors.New("error get connection to MySQL")
	}

	s := &Store{
		client: client,
	}

	return s, nil
}

// TODO: use uuid.FromString(id) for all id
func (s Store) Get(ctx context.Context, id string) (*domain.Link, error) {
	panic("implement me")
}

func (s Store) List(ctx context.Context, filter *query.Filter) (*domain.Links, error) {
	panic("implement me")
}

func (s Store) Add(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	panic("implement me")
}

func (s Store) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	panic("implement me")
}

func (s Store) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
