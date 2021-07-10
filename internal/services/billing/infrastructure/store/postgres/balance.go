package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	v1 "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
)

type Balance struct {
	client *pgxpool.Pool
}

func (b *Balance) Init(ctx context.Context, db *db.Store) error {
	b.client = db.Store.GetConn().(*pgxpool.Pool)
	return nil
}

func (b *Balance) Get(ctx context.Context, id *v1.Balance) (*v1.Balance, error) {
	panic("implement me")
}

func (b *Balance) List(ctx context.Context, filter interface{}) ([]*v1.Balance, error) {
	panic("implement me")
}

func (b *Balance) Update(ctx context.Context, in *v1.Balance) (*v1.Balance, error) {
	panic("implement me")
}
