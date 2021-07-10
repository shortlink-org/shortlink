package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	v1 "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
)

type Tariff struct {
	client *pgxpool.Pool
}

func (t Tariff) Init(ctx context.Context, db *db.Store) error {
	t.client = db.Store.GetConn().(*pgxpool.Pool)
	return nil
}

func (t Tariff) Get(ctx context.Context, id string) (*v1.Tariff, error) {
	panic("implement me")
}

func (t Tariff) List(ctx context.Context, filter interface{}) ([]*v1.Tariff, error) {
	panic("implement me")
}

func (t Tariff) Add(ctx context.Context, in *v1.Tariff) (*v1.Tariff, error) {
	panic("implement me")
}

func (t Tariff) Update(ctx context.Context, in *v1.Tariff) (*v1.Tariff, error) {
	panic("implement me")
}

func (t Tariff) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
