package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

type Payment struct {
	client *pgxpool.Pool
}

func (p *Payment) Init(ctx context.Context, db *db.Store) error {
	p.client = db.Store.GetConn().(*pgxpool.Pool)
	return nil
}

func (p *Payment) Get(ctx context.Context, id string) (*v1.Payment, error) {
	panic("implement me")
}

func (p *Payment) List(ctx context.Context, filter interface{}) ([]*v1.Payment, error) {
	panic("implement me")
}

func (p *Payment) Add(ctx context.Context, in *v1.Payment) (*v1.Payment, error) {
	panic("implement me")
}

func (p *Payment) Update(ctx context.Context, in *v1.Payment) (*v1.Payment, error) {
	panic("implement me")
}

func (p *Payment) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
