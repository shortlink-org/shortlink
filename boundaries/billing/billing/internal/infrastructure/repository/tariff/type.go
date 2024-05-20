package tariff_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/tariff/v1"
)

type Repository interface {
	Get(ctx context.Context, id string) (*billing.Tariff, error)
	List(ctx context.Context, filter any) (*billing.Tariffs, error)
	Add(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Update(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Delete(ctx context.Context, id string) error
}

type tariff struct {
	client *pgxpool.Pool
}
