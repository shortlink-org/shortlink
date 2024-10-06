package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/open-policy-agent/opa/rego"
	"github.com/shopspring/decimal"
)

// CartItem represents a cart item.
type CartItem struct {
	ProductId uuid.UUID       `json:"productId"`
	Quantity  int32           `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	Discount  decimal.Decimal `json:"discount"`
	Tax       decimal.Decimal `json:"tax"`
	Brand     string          `json:"brand"`
}

// CartItems represents a list of cart items.
type CartItems []CartItem

// CartState represents the cart state.
type CartState struct {
	Items      CartItems `json:"items"`
	CustomerId uuid.UUID `json:"customerId"`
}

// LoadCart loads a cart from a JSON file.
func LoadCart(filePath string) (CartState, error) {
	var cart CartState
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cart, err
	}
	err = json.Unmarshal(file, &cart)
	return cart, err
}

// LoadOPAPolicies loads all OPA policies from a directory with a specific query
func LoadOPAPolicies(dir string, query string) (*rego.PreparedEvalQuery, error) {
	preparedQuery, err := rego.New(
		rego.Query(query),
		rego.Load([]string{dir}, nil),
	).PrepareForEval(context.Background())
	if err != nil {
		return nil, err
	}
	return &preparedQuery, nil
}

// EvaluateTaxPolicy evaluates the tax policy using OPA.
func EvaluateTaxPolicy(cart CartItems, query *rego.PreparedEvalQuery, params map[string]interface{}) (decimal.Decimal, error) {
	var items []map[string]interface{}
	for _, item := range cart {
		items = append(items, map[string]interface{}{
			"productId": item.ProductId.String(), // Convert UUID to string
			"quantity":  item.Quantity,
			"price":     item.Price.InexactFloat64(), // Convert decimal to float64
			"brand":     item.Brand,
		})
	}
	input := map[string]interface{}{
		"items":  items,
		"params": params, // Include tax parameters if needed
	}

	rs, err := query.Eval(context.Background(), rego.EvalInput(input))
	if err != nil {
		return decimal.Zero, err
	}

	if len(rs) == 0 {
		return decimal.Zero, nil
	}

	// Extract the single value from the result
	expr := rs[0].Expressions[0].Value
	var taxFloat float64

	switch v := expr.(type) {
	case float64:
		taxFloat = v
	case string:
		taxFloat, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return decimal.Zero, fmt.Errorf("invalid number format for tax: %v", err)
		}
	case json.Number:
		taxFloat, err = v.Float64()
		if err != nil {
			return decimal.Zero, fmt.Errorf("invalid number for tax: %v", err)
		}
	default:
		return decimal.Zero, fmt.Errorf("unexpected type for tax: %T", expr)
	}

	return decimal.NewFromFloat(taxFloat), nil
}

// EvaluateDiscountPolicy evaluates the discount policy using OPA.
func EvaluateDiscountPolicy(cart CartItems, query *rego.PreparedEvalQuery, params map[string]interface{}) (decimal.Decimal, error) {
	var items []map[string]interface{}
	for _, item := range cart {
		items = append(items, map[string]interface{}{
			"productId": item.ProductId.String(), // Convert UUID to string
			"quantity":  item.Quantity,
			"price":     item.Price.InexactFloat64(), // Convert decimal to float64
			"brand":     item.Brand,
		})
	}
	input := map[string]interface{}{
		"items":  items,
		"params": params, // Include discount parameters
	}

	rs, err := query.Eval(context.Background(), rego.EvalInput(input))
	if err != nil {
		return decimal.Zero, err
	}

	if len(rs) == 0 {
		return decimal.Zero, nil
	}

	// Extract the single value from the result
	expr := rs[0].Expressions[0].Value
	var discountFloat float64

	switch v := expr.(type) {
	case float64:
		discountFloat = v
	case string:
		discountFloat, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return decimal.Zero, fmt.Errorf("invalid number format for discount: %v", err)
		}
	case json.Number:
		discountFloat, err = v.Float64()
		if err != nil {
			return decimal.Zero, fmt.Errorf("invalid number for discount: %v", err)
		}
	default:
		return decimal.Zero, fmt.Errorf("unexpected type for discount: %T", expr)
	}

	return decimal.NewFromFloat(discountFloat), nil
}

// SaveResultToFile saves the final result to the out directory
func SaveResultToFile(result map[string]interface{}, outDir string, filename string) error {
	// Ensure the output directory exists
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Create the output file path
	outputFile := filepath.Join(outDir, filename)

	// Marshal the result into JSON
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	// Write the JSON data to the output file
	err = ioutil.WriteFile(outputFile, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Final result saved to %s\n", outputFile)
	return nil
}

func main() {
	// Cart files to process
	cartFiles := []string{
		"tests/fixtures/cart_1.json",
		"tests/fixtures/cart_2.json",
		"tests/fixtures/cart_3.json",
		"tests/fixtures/cart_4.json",
		"tests/fixtures/cart_5.json",
	}

	// Define discount parameters
	discountParams := map[string]interface{}{
		"apple_samsung_discount": 0.05,
	}

	// Define tax parameters if any (add as needed)
	taxParams := map[string]interface{}{
		// Example: "tax_rate": 0.20,
	}

	// Load tax policies with specific query
	taxQuery, err := LoadOPAPolicies("policies/taxes/", "data.pricing.tax.total_markup")
	if err != nil {
		log.Fatalf("Failed to load tax policies: %v", err)
	}

	// Load discount policies with specific query
	discountQuery, err := LoadOPAPolicies("policies/discounts/", "data.pricing.discount.total_brand_discount")
	if err != nil {
		log.Fatalf("Failed to load discount policies: %v", err)
	}

	// Process each cart
	for _, cartFile := range cartFiles {
		// Load the cart from the file
		cart, err := LoadCart(cartFile)
		if err != nil {
			log.Fatalf("Failed to load cart %s: %v", cartFile, err)
		}

		// Calculate taxes and discounts
		totalTax, err := EvaluateTaxPolicy(cart.Items, taxQuery, taxParams)
		if err != nil {
			log.Fatalf("Failed to evaluate tax for cart %s: %v", cartFile, err)
		}

		totalDiscount, err := EvaluateDiscountPolicy(cart.Items, discountQuery, discountParams)
		if err != nil {
			log.Fatalf("Failed to evaluate discount for cart %s: %v", cartFile, err)
		}

		// Final price calculation
		finalPrice := decimal.Zero
		for _, item := range cart.Items {
			// Calculate per-item total: (Price + Tax - Discount) * Quantity
			itemTotal := item.Price.Add(totalTax).Sub(totalDiscount).Mul(decimal.NewFromInt32(item.Quantity))
			finalPrice = finalPrice.Add(itemTotal)
		}

		// Prepare the result
		result := map[string]interface{}{
			"customerId":    cart.CustomerId.String(),
			"totalTax":      totalTax.StringFixed(2),
			"totalDiscount": totalDiscount.StringFixed(2),
			"finalPrice":    finalPrice.StringFixed(2),
		}

		// Save result to the out folder, one file per customer
		filename := fmt.Sprintf("cart_result_%s.json", cart.CustomerId.String())
		err = SaveResultToFile(result, "out", filename)
		if err != nil {
			log.Fatalf("Failed to save result for cart %s: %v", cartFile, err)
		}
	}
}
