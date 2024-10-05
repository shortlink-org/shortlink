package entity

import "github.com/shopspring/decimal"

type Goods struct {
	Brand    string                 `json:"brand"`
	Model    string                 `json:"model"`
	Price    decimal.Decimal        `json:"price"`
	Stock    int                    `json:"stock"`
	Category string                 `json:"category"`
	Tags     []string               `json:"tags"`
	Features map[string]interface{} `json:"features"`
}
