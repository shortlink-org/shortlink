package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	v1 "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

type Account struct {
	client *pgxpool.Pool
}

func (a *Account) Init(ctx context.Context, db *db.Store) error {
	a.client = db.Store.GetConn().(*pgxpool.Pool)
	return nil
}

func (a *Account) Get(ctx context.Context, id string) (*v1.Account, error) {
	panic("implement me")
}

func (a *Account) List(ctx context.Context, filter interface{}) ([]*v1.Account, error) {
	panic("implement me")
}

func (a *Account) Add(ctx context.Context, in *v1.Account) (*v1.Account, error) {
	id := uuid.MustParse(in.Id)
	userId := uuid.MustParse(in.UserId)
	tariffId := uuid.MustParse(in.TariffId)

	// query builder
	account := psql.Insert("billing.account").
		Columns("id", "user_id", "tariff_id").
		Values(id, userId, tariffId)

	q, args, err := account.ToSql()
	if err != nil {
		return nil, err
	}

	row := a.client.QueryRow(ctx, q, args...)
	errScan := row.Scan().Error()
	if errScan == "no rows in result set" {
		return in, nil
	}
	if errScan != "" {
		return nil, fmt.Errorf(errScan)
	}

	return in, nil
}

func (a *Account) Update(ctx context.Context, in *v1.Account) (*v1.Account, error) {
	panic("implement me")
}

func (a *Account) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
