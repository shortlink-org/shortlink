package account_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1"
)

type Repository interface {
	Get(ctx context.Context, id string) (*v1.Account, error)
	List(ctx context.Context, filter any) ([]*v1.Account, error)
	Add(ctx context.Context, in *v1.Account) (*v1.Account, error)
	Update(ctx context.Context, in *v1.Account) (*v1.Account, error)
	Delete(ctx context.Context, id string) error
}

type account struct {
	client *pgxpool.Pool
}
