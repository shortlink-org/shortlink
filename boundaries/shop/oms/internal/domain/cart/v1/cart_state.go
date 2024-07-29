package v1

import (
	"sync"

	"github.com/google/uuid"
)

// CartState
type CartState struct {
	mu sync.Mutex

	// items is the cart items
	items []CartItem
	// customerId is the customer ID
	customerId uuid.UUID
}

// NewCartState creates a new cart state.
func NewCartState(customerId uuid.UUID) *CartState {
	return &CartState{
		items:      make([]CartItem, 0),
		customerId: customerId,
	}
}

// GetItems returns the value of the items field.
func (m *CartState) GetItems() []CartItem {
	return m.items
}

// GetCustomerId returns the value of the customerId field.
func (m *CartState) GetCustomerId() uuid.UUID {
	return m.customerId
}

// AddItem adds an item to the cart.
func (m *CartState) AddItem(item CartItem) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// if the item already exists in the cart, increment the quantity
	for i, cartItem := range m.items {
		if cartItem.productId == item.productId {
			m.items[i].quantity += item.quantity
			return
		}
	}

	m.items = append(m.items, item)
}

// RemoveItem removes an item from the cart.
func (m *CartState) RemoveItem(item CartItem) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// if the item already exists in the cart, decrement the quantity
	for i, cartItem := range m.items {
		if cartItem.productId == item.productId {
			m.items[i].quantity -= item.quantity
			if m.items[i].quantity <= 0 {
				m.items = append(m.items[:i], m.items[i+1:]...)
			}
			return
		}
	}
}
