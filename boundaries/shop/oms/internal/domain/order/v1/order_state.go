package v1

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/looplab/fsm"
)

// OrderState represents the order state.
type OrderState struct {
	mu sync.Mutex

	// orderID is the order ID
	id uuid.UUID
	// items is the list of order items
	items Items
	// customerId is the customer ID
	customerId uuid.UUID
	// fsm is the finite state machine for the order status
	fsm *fsm.FSM
}

// NewOrderState creates a new order state.
func NewOrderState(customerId uuid.UUID) *OrderState {
	return &OrderState{
		id:         uuid.New(),
		items:      make([]Item, 0),
		customerId: customerId,
		fsm: fsm.NewFSM(
			OrderStatus_ORDER_STATUS_PENDING.String(),
			fsm.Events{
				{
					Name: OrderStatus_ORDER_STATUS_PENDING.String(),
					Src:  []string{OrderStatus_ORDER_STATUS_PENDING.String()},
					Dst:  OrderStatus_ORDER_STATUS_PROCESSING.String(),
				},
				{
					Name: OrderStatus_ORDER_STATUS_PROCESSING.String(),
					Src:  []string{OrderStatus_ORDER_STATUS_PROCESSING.String()},
					Dst:  OrderStatus_ORDER_STATUS_PROCESSING.String(),
				},
				{
					Name: OrderStatus_ORDER_STATUS_CANCELLED.String(),
					Src:  []string{OrderStatus_ORDER_STATUS_PENDING.String(), OrderStatus_ORDER_STATUS_PROCESSING.String()},
					Dst:  OrderStatus_ORDER_STATUS_CANCELLED.String(),
				},
				{
					Name: OrderStatus_ORDER_STATUS_COMPLETED.String(),
					Src:  []string{OrderStatus_ORDER_STATUS_PROCESSING.String()},
					Dst:  OrderStatus_ORDER_STATUS_COMPLETED.String(),
				},
			},
			fsm.Callbacks{},
		),
	}
}

// GetOrderID returns the order ID.
func (o *OrderState) GetOrderID() uuid.UUID {
	return o.id
}

// GetItems returns the list of order items.
func (o *OrderState) GetItems() Items {
	return o.items
}

// GetCustomerId returns the customer ID.
func (o *OrderState) GetCustomerId() uuid.UUID {
	return o.customerId
}

// GetStatus returns the current status of the order.
func (o *OrderState) GetStatus() OrderStatus {
	for k, v := range OrderStatus_name {
		if v == o.fsm.Current() {
			return OrderStatus(k)
		}
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

// CreateOrder sets the initial order details.
func (o *OrderState) CreateOrder(ctx context.Context, items Items) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	err := o.fsm.Event(ctx, OrderStatus_ORDER_STATUS_PENDING.String())
	if err != nil {
		return err
	}

	o.items = items
	return nil
}

// UpdateOrder updates the order details.
func (o *OrderState) UpdateOrder(ctx context.Context, items Items) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Update quantities of existing items and add new items
	itemMap := make(map[uuid.UUID]*Item)
	for i := range o.items {
		itemMap[o.items[i].productId] = &o.items[i]
	}
	for _, item := range items {
		if existingItem, exists := itemMap[item.productId]; exists {
			existingItem.quantity = item.quantity
			existingItem.price = item.price
		} else {
			o.items = append(o.items, item)
		}
	}

	return nil
}

// CancelOrder cancels the order.
func (o *OrderState) CancelOrder(ctx context.Context) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	err := o.fsm.Event(ctx, OrderStatus_ORDER_STATUS_CANCELLED.String())
	if err != nil {
		return err
	}

	return nil
}

// CompleteOrder completes the order.
func (o *OrderState) CompleteOrder(ctx context.Context) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	err := o.fsm.Event(ctx, OrderStatus_ORDER_STATUS_COMPLETED.String())
	if err != nil {
		return err
	}

	return nil
}
