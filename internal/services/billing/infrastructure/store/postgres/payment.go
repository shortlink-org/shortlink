package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/batazor/shortlink/internal/pkg/db"
	payment "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

type Payment struct {
	client *pgxpool.Pool
}

func (p *Payment) Init(ctx context.Context, db *db.Store) error {
	p.client = db.Store.GetConn().(*pgxpool.Pool)
	return nil
}

func (p *Payment) Get(ctx context.Context, id string) (*payment.Payment, error) {
	panic("implement me")
}

func (p *Payment) List(ctx context.Context, filter interface{}) ([]*payment.Payment, error) {
	panic("implement me")
}

func (p *Payment) Add(ctx context.Context, in *payment.Payment) (*payment.Payment, error) {
	//// get uncommitted events
	//events := in.AggregateHandler.Uncommitted()
	//
	//// validate
	//if len(events) == 0 {
	//	return nil, fmt.Errorf("Not found events for %s", in.Id)
	//}
	//
	//// query builder
	//entities := psql.Insert("entities").
	//	Columns("id", "type", "version").
	//	Values(in.Id, events[0].Type, events[0].Version)
	//
	//q, args, err := entities.ToSql()
	//if err != nil {
	//	return nil, err
	//}
	//
	//// save first event to entities
	//row := p.client.QueryRow(ctx, q, args...)
	//errScan := row.Scan()
	//if errors.Is(errScan, pgx.ErrNoRows) {
	//	return in, nil
	//}
	//if errScan.Error() != "" {
	//	return nil, errScan
	//}
	//
	//return in, nil
	panic("implement me")
}

func (p *Payment) Update(ctx context.Context, in *payment.Payment) (*payment.Payment, error) {
	panic("implement me")
}

func (p *Payment) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
