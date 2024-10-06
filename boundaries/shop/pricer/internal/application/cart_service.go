package application

import (
	"context"
	"fmt"
	"log"

	"github.com/shopspring/decimal"

	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/domain"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/infrastructure"
)

// CartTotal represents the total calculation result
type CartTotal struct {
	TotalTax      decimal.Decimal `json:"totalTax"`
	TotalDiscount decimal.Decimal `json:"totalDiscount"`
	FinalPrice    decimal.Decimal `json:"finalPrice"`
	Policies      []string        `json:"policies"`
}

// CartService orchestrates cart operations
type CartService struct {
	DiscountPolicy infrastructure.PolicyEvaluator
	TaxPolicy      infrastructure.PolicyEvaluator
	PolicyNames    []string
}

type DiscountPolicy infrastructure.PolicyEvaluator
type TaxPolicy infrastructure.PolicyEvaluator

// NewCartService creates a new CartService
func NewCartService(discountPolicy DiscountPolicy, taxPolicy TaxPolicy, policyNames []string) *CartService {
	return &CartService{
		DiscountPolicy: discountPolicy,
		TaxPolicy:      taxPolicy,
		PolicyNames:    policyNames,
	}
}

// CalculateTotal computes the total price, applying discounts and taxes
func (s *CartService) CalculateTotal(ctx context.Context, cart *domain.Cart, discountParams, taxParams map[string]interface{}) (CartTotal, error) {
	var total CartTotal

	// Evaluate Discount Policy
	log.Printf("Evaluating Discount Policy for CustomerID: %s", cart.CustomerID)
	totalDiscountFloat, err := s.DiscountPolicy.Evaluate(ctx, cart, discountParams)
	if err != nil {
		return total, fmt.Errorf("failed to evaluate discount policy: %w", err)
	}
	log.Printf("Total Discount: %.2f", totalDiscountFloat)
	totalDiscount := decimal.NewFromFloat(totalDiscountFloat)

	// Evaluate Tax Policy
	log.Printf("Evaluating Tax Policy for CustomerID: %s", cart.CustomerID)
	totalTaxFloat, err := s.TaxPolicy.Evaluate(ctx, cart, taxParams)
	if err != nil {
		return total, fmt.Errorf("failed to evaluate tax policy: %w", err)
	}
	log.Printf("Total Tax: %.2f", totalTaxFloat)
	totalTax := decimal.NewFromFloat(totalTaxFloat)

	// Calculate Final Price
	log.Printf("Calculating Final Price for CustomerID: %s", cart.CustomerID)
	finalPrice := decimal.Zero
	for _, item := range cart.Items {
		// Calculate per-item total: (Price + Tax - Discount) * Quantity
		itemTotal := item.Price.Add(totalTax).Sub(totalDiscount).Mul(decimal.NewFromInt32(item.Quantity))
		finalPrice = finalPrice.Add(itemTotal)
		log.Printf("ItemID: %s, ItemTotal: %.2f", item.ProductID, itemTotal)
	}

	// Prepare the CartTotal
	total = CartTotal{
		TotalTax:      totalTax,
		TotalDiscount: totalDiscount,
		FinalPrice:    finalPrice,
		Policies:      s.PolicyNames,
	}

	log.Printf("Final Price for CustomerID %s: %.2f", cart.CustomerID, finalPrice)
	return total, nil
}
