package v1

// Payments - list payments
type Payments struct {
	// list of payments
	list []*Payment
}

// GetList
func (o *Payments) GetList() []*Payment {
	return o.list
}
