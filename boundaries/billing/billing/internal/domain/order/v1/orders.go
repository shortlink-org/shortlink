package v1

// Order list
type Orders struct {
	// order list
	list []*Order
}

// GetList
func (o *Orders) GetList() []*Order {
	return o.list
}
