package tariff_repository

import (
	"context"
	"embed"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/billing/tariff/v1"
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
	err := migrate.Migration(ctx, store, migrations, "repository_tariff")
	if err != nil {
		return nil, err
	}

	return &tariff{
		client: client,
	}, nil
}

func (t *tariff) Get(ctx context.Context, id string) (*v1.Tariff, error) {
	resp := &v1.Tariff{}

	query := psql.Select("id", "name", "payload").
		From("billing.tariff").
		Where(squirrel.Eq{"id": id})

	q, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := t.client.QueryRow(ctx, q, args...)
	errScan := row.Scan(&resp.Id, &resp.Name, &resp.Payload)
	if errors.Is(errScan, pgx.ErrNoRows) {
		return resp, nil
	}
	if errScan.Error() != "" {
		return nil, errScan
	}

	return resp, nil
}

func (t *tariff) List(ctx context.Context, filter any) (*v1.Tariffs, error) {
	query := psql.Select("id", "name", "payload").
		From("billing.tariff")
	q, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := t.client.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	response := v1.Tariffs{}

	for rows.Next() {
		var result v1.Tariff
		err = rows.Scan(&result.Id, &result.Name, &result.Payload)
		if err != nil {
			return nil, err
		}

		response.List = append(response.GetList(), &result)
	}

	return &response, nil
}

func (t *tariff) Add(ctx context.Context, in *v1.Tariff) (*v1.Tariff, error) {
	query := psql.Insert("billing.tariff").
		Columns("id", "name", "payload").
		Values(in.GetId(), in.GetName(), in.GetPayload())

	q, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := t.client.QueryRow(ctx, q, args...)
	errScan := row.Scan()
	if errors.Is(errScan, pgx.ErrNoRows) {
		return in, nil
	}
	if errScan.Error() != "" {
		return nil, errScan
	}

	return in, nil
}

func (t *tariff) Update(ctx context.Context, in *v1.Tariff) (*v1.Tariff, error) {
	query := psql.Update("billing.tariff").
		Set("name", in.GetName()).
		Set("payload", in.GetPayload()).
		Where(squirrel.Eq{"id": in.GetId()})

	q, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := t.client.QueryRow(ctx, q, args...)
	errScan := row.Scan()
	if errors.Is(errScan, pgx.ErrNoRows) {
		return in, nil
	}
	if errScan.Error() != "" {
		return nil, errScan
	}

	return in, nil
}

func (t *tariff) Delete(ctx context.Context, id string) error {
	query := psql.Delete("billing.tariff").
		Where(squirrel.Eq{"hash": id})

	q, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = t.client.Exec(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
