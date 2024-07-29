package v1

type SignalChannels string

// String returns the string representation of the SignalChannels
func (s SignalChannels) String() string {
	return string(s)
}

var (
	// ADD_TO_CART is a signal to add an item to the cart
	ADD_TO_CART = SignalChannels("ADD_TO_CART")
	// REMOVE_FROM_CART is a signal to remove an item from the cart
	REMOVE_FROM_CART = SignalChannels("REMOVE_FROM_CART")
)

// AddToCartSignal is a signal to add an item to the cart
type AddToCartSignal struct {
	Route SignalChannels

	// Item is the item to add to the cart
	Item CartItem
}

// RemoveFromCartSignal is a signal to remove an item from the cart
type RemoveFromCartSignal struct {
	Route SignalChannels

	// Item is the item to remove from the cart
	Item CartItem
}
