package account_repository

import (
	"context"
	"embed"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/db/postgres/migrate"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS

	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func New(ctx context.Context, store db.DB) (Repository, error) {
	client, ok := store.GetConn().(*pgxpool.Pool)
	if !ok {
		return nil, db.ErrGetConnection
	}

	// Migration ---------------------------------------------------------------------------------------------------
	err := migrate.Migration(ctx, store, migrations, "repository_account")
	if err != nil {
		return nil, err
	}

	return &account{
		client: client,
	}, nil
}

func (a *account) Get(ctx context.Context, id string) (*v1.Account, error) {
	panic("implement me")
}

func (a *account) List(ctx context.Context, filter any) ([]*v1.Account, error) {
	panic("implement me")
}

func (a *account) Add(ctx context.Context, in *v1.Account) (*v1.Account, error) {
	id := uuid.MustParse(in.GetId())
	userId := uuid.MustParse(in.GetUserId())
	tariffId := uuid.MustParse(in.GetTariffId())

	account := psql.Insert("billing.account").
		Columns("id", "user_id", "tariff_id").
		Values(id, userId, tariffId)

	q, args, err := account.ToSql()
	if err != nil {
		return nil, err
	}

	row := a.client.QueryRow(ctx, q, args...)
	errScan := row.Scan()
	if errors.Is(errScan, pgx.ErrNoRows) {
		return in, nil
	}
	if errScan.Error() != "" {
		return nil, errScan
	}

	return in, nil
}

func (a *account) Update(ctx context.Context, in *v1.Account) (*v1.Account, error) {
	panic("implement me")
}

func (a *account) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
