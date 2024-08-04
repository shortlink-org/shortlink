package v1

import (
	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
)

type OrderEvent struct {
	Event v2.Event
	Items []*WorkerOrderItem
}
