package payment

import (
	events "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
)

type Payment struct {
	events.Payment
}

func (p *Payment) Create() {
	panic("implement me")
}

func (p *Payment) Close() {
	panic("implement me")
}
