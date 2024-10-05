package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/open-policy-agent/opa/rego"
)

// Goods represents a product with its attributes
type Goods struct {
	Brand    string  `json:"brand"`
	Model    string  `json:"model"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}

// LoadOPAPolicies loads all OPA policies from a directory
func LoadOPAPolicies(dir string) (*rego.PreparedEvalQuery, error) {
	// Find all rego files in the directory
	var regoFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".rego" {
			regoFiles = append(regoFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Prepare an OPA query that loads all rego files
	ctx := context.TODO()
	query, err := rego.New(
		rego.Query("data.pricing"),
		rego.Load(regoFiles, nil),
	).PrepareForEval(ctx)

	if err != nil {
		return nil, err
	}

	return &query, nil
}

// EvaluateTaxPolicy evaluates all tax policies using OPA
func EvaluateTaxPolicy(goods Goods, query *rego.PreparedEvalQuery) (float64, error) {
	// Provide input to the OPA policy
	input := map[string]interface{}{
		"price": goods.Price,
	}

	// Evaluate the policy
	ctx := context.TODO()
	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return 0, err
	}

	// Sum all tax policies
	tax := 0.0
	for _, result := range rs {
		for _, expr := range result.Expressions {
			if v, ok := expr.Value.(map[string]interface{}); ok {
				if t, ok := v["tax"].(float64); ok {
					tax += t
				}
			}
		}
	}
	return tax, nil
}

// EvaluateDiscountPolicy evaluates all discount policies using OPA
func EvaluateDiscountPolicy(goods Goods, query *rego.PreparedEvalQuery) (float64, error) {
	// Get current time
	currentTime := time.Now().Format("15:04")

	// Provide input to the OPA policy
	input := map[string]interface{}{
		"price": goods.Price,
		"brand": goods.Brand,
		"time":  currentTime,
		"count": goods.Stock,
	}

	// Evaluate the policy
	ctx := context.TODO()
	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return 0, err
	}

	// Sum all discount policies
	discount := 0.0
	for _, result := range rs {
		for _, expr := range result.Expressions {
			if v, ok := expr.Value.(map[string]interface{}); ok {
				if d, ok := v["discount"].(float64); ok {
					discount += d
				}
			}
		}
	}
	return discount, nil
}

// CalculateFinalPrice calculates the final price using OPA for taxes and discounts
func CalculateFinalPrice(goods Goods, taxQuery, discountQuery *rego.PreparedEvalQuery) (float64, error) {
	// Evaluate the tax policy
	tax, err := EvaluateTaxPolicy(goods, taxQuery)
	if err != nil {
		return 0, err
	}

	// Evaluate the discount policy
	discount, err := EvaluateDiscountPolicy(goods, discountQuery)
	if err != nil {
		return 0, err
	}

	// Calculate the final price
	finalPrice := (goods.Price + tax) - discount
	return finalPrice, nil
}

func main() {
	// Load tax policies from the policies/taxes directory
	taxQuery, err := LoadOPAPolicies("policies/taxes")
	if err != nil {
		log.Fatalf("Failed to load tax policies: %v", err)
	}

	// Load discount policies from the policies/discounts directory
	discountQuery, err := LoadOPAPolicies("policies/discounts")
	if err != nil {
		log.Fatalf("Failed to load discount policies: %v", err)
	}

	// Example goods
	goods := Goods{
		Brand:    "Apple",
		Model:    "iPhone 13",
		Price:    999.00,
		Stock:    3,
		Category: "Electronics",
	}

	// Calculate the final price using OPA
	finalPrice, err := CalculateFinalPrice(goods, taxQuery, discountQuery)
	if err != nil {
		log.Fatalf("Failed to calculate final price: %v", err)
	}

	fmt.Printf("Final price for %s %s: $%.2f\n", goods.Brand, goods.Model, finalPrice)
}
