package domain

import "github.com/shopspring/decimal"

type CartTotal struct {
	TotalTax      decimal.Decimal
	TotalDiscount decimal.Decimal
	FinalPrice    decimal.Decimal
	Policies      []string
}
