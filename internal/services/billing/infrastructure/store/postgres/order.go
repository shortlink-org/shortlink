package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	v1 "github.com/batazor/shortlink/internal/services/billing/domain/billing/order/v1"
)

type Order struct {
	client *pgxpool.Pool
}

func (o *Order) Init(ctx context.Context, db *db.Store) error {
	o.client = db.Store.GetConn().(*pgxpool.Pool)
	return nil
}

func (o *Order) Get(ctx context.Context, id string) (*v1.Order, error) {
	panic("implement me")
}

func (o *Order) List(ctx context.Context, filter interface{}) ([]*v1.Order, error) {
	panic("implement me")
}

func (o *Order) Add(ctx context.Context, in *v1.Order) (*v1.Order, error) {
	panic("implement me")
}

func (o *Order) Update(ctx context.Context, in *v1.Order) (*v1.Order, error) {
	panic("implement me")
}

func (o *Order) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
