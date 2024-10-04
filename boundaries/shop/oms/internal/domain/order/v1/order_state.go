package v1

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/shortlink-org/shortlink/pkg/fsm"
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

// NewOrderState creates a new OrderState instance with the given customer ID.
func NewOrderState(customerId uuid.UUID) *OrderState {
	order := &OrderState{
		id:         uuid.New(),
		items:      make(Items, 0),
		customerId: customerId,
	}

	// Initialize the FSM with the initial state.
	order.fsm = fsm.New(fsm.State(OrderStatus_ORDER_STATUS_PENDING.String()))

	// Define transition rules.
	order.fsm.AddTransitionRule(
		fsm.State(OrderStatus_ORDER_STATUS_PENDING.String()),
		fsm.Event(OrderStatus_ORDER_STATUS_PENDING.String()),
		fsm.State(OrderStatus_ORDER_STATUS_PROCESSING.String()),
	)
	order.fsm.AddTransitionRule(
		fsm.State(OrderStatus_ORDER_STATUS_PROCESSING.String()),
		fsm.Event(OrderStatus_ORDER_STATUS_PROCESSING.String()),
		fsm.State(OrderStatus_ORDER_STATUS_PROCESSING.String()),
	)
	order.fsm.AddTransitionRule(
		fsm.State(OrderStatus_ORDER_STATUS_PENDING.String()),
		fsm.Event(OrderStatus_ORDER_STATUS_CANCELLED.String()),
		fsm.State(OrderStatus_ORDER_STATUS_CANCELLED.String()),
	)
	order.fsm.AddTransitionRule(
		fsm.State(OrderStatus_ORDER_STATUS_PROCESSING.String()),
		fsm.Event(OrderStatus_ORDER_STATUS_CANCELLED.String()),
		fsm.State(OrderStatus_ORDER_STATUS_CANCELLED.String()),
	)
	order.fsm.AddTransitionRule(
		fsm.State(OrderStatus_ORDER_STATUS_PROCESSING.String()),
		fsm.Event(OrderStatus_ORDER_STATUS_COMPLETED.String()),
		fsm.State(OrderStatus_ORDER_STATUS_COMPLETED.String()),
	)

	// Set up callbacks.
	order.fsm.SetOnEnterState(order.onEnterState)
	order.fsm.SetOnExitState(order.onExitState)

	return order
}

// onEnterState is the callback executed when entering a new state.
func (o *OrderState) onEnterState(ctx context.Context, from, to fsm.State, event fsm.Event) {
	fmt.Printf("Order %s entered state '%s' due to event '%s'\n", o.id, to, event)
}

// onExitState is the callback executed when exiting a state.
func (o *OrderState) onExitState(ctx context.Context, from, to fsm.State, event fsm.Event) {
	fmt.Printf("Order %s exited state '%s' due to event '%s'\n", o.id, from, event)
}

// GetOrderID returns the unique identifier of the order.
func (o *OrderState) GetOrderID() uuid.UUID {
	return o.id
}

// GetItems returns a copy of the list of items in the order.
func (o *OrderState) GetItems() Items {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Return a copy to prevent external modification.
	itemsCopy := make(Items, len(o.items))
	copy(itemsCopy, o.items)

	return itemsCopy
}

// GetCustomerId returns the customer ID associated with the order.
func (o *OrderState) GetCustomerId() uuid.UUID {
	return o.customerId
}

// GetStatus returns the current status of the order.
func (o *OrderState) GetStatus() OrderStatus {
	o.mu.Lock()
	defer o.mu.Unlock()

	currentState := o.fsm.GetCurrentState()
	for k, v := range OrderStatus_name {
		if v == currentState.String() {
			return OrderStatus(k)
		}
	}

	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

// CreateOrder initializes the order with the provided items and transitions it to Processing state.
func (o *OrderState) CreateOrder(ctx context.Context, items Items) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Trigger the transition event to Processing.
	err := o.fsm.TriggerEvent(ctx, fsm.Event(OrderStatus_ORDER_STATUS_PENDING.String()))
	if err != nil {
		return err
	}

	o.items = items
	return nil
}

// UpdateOrder updates the order's items. It modifies existing items and adds new ones as needed.
func (o *OrderState) UpdateOrder(ctx context.Context, items Items) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Create a map for quick lookup of existing items.
	itemMap := make(map[uuid.UUID]*Item)
	for i := range o.items {
		itemMap[o.items[i].productId] = &o.items[i]
	}

	// Update quantities and prices of existing items or add new items.
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

// CancelOrder transitions the order to the Cancelled state.
func (o *OrderState) CancelOrder(ctx context.Context) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Trigger the transition event to Cancel.
	err := o.fsm.TriggerEvent(ctx, fsm.Event(OrderStatus_ORDER_STATUS_CANCELLED.String()))
	if err != nil {
		return err
	}

	return nil
}

// CompleteOrder transitions the order to the Completed state.
func (o *OrderState) CompleteOrder(ctx context.Context) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Trigger the transition event to Complete.
	err := o.fsm.TriggerEvent(ctx, fsm.Event(OrderStatus_ORDER_STATUS_COMPLETED.String()))
	if err != nil {
		return err
	}

	return nil
}
