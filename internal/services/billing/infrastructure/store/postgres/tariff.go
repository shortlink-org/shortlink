package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	v1 "github.com/batazor/shortlink/internal/services/billing/domain/billing/tariff/v1"
	v12 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
)

type Tariff struct {
	client *pgxpool.Pool
}

func (t *Tariff) Init(ctx context.Context, db *db.Store) error {
	t.client = db.Store.GetConn().(*pgxpool.Pool)
	return nil
}

func (t *Tariff) Get(ctx context.Context, id string) (*v1.Tariff, error) {
	resp := &v1.Tariff{}

	// query builder
	tariff := psql.Select("id", "name", "payload").
		From("billing.tariff").
		Where(squirrel.Eq{"id": id})

	q, args, err := tariff.ToSql()
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

func (t *Tariff) List(ctx context.Context, filter interface{}) (*v1.Tariffs, error) {
	// query builder
	tariffs := psql.Select("id", "name", "payload").
		From("billing.tariff")
	q, args, err := tariffs.ToSql()
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
			return nil, &v12.NotFoundError{Link: &v12.Link{}, Err: fmt.Errorf("Not found links")}
		}

		response.List = append(response.List, &result)
	}

	return &response, nil
}

func (t *Tariff) Add(ctx context.Context, in *v1.Tariff) (*v1.Tariff, error) {
	// query builder
	tariff := psql.Insert("billing.tariff").
		Columns("id", "name", "payload").
		Values(in.Id, in.Name, in.Payload)

	q, args, err := tariff.ToSql()
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

func (t *Tariff) Update(ctx context.Context, in *v1.Tariff) (*v1.Tariff, error) {
	// query builder
	tariff := psql.Update("billing.tariff").
		Set("name", in.Name).
		Set("payload", in.Payload).
		Where(squirrel.Eq{"id": in.Id})

	q, args, err := tariff.ToSql()
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

func (t *Tariff) Delete(ctx context.Context, id string) error {
	// query builder
	tariff := psql.Delete("billing.tariff").
		Where(squirrel.Eq{"hash": id})

	q, args, err := tariff.ToSql()
	if err != nil {
		return err
	}

	_, err = t.client.Exec(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
