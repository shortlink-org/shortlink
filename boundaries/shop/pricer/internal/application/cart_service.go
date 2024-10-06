package application

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/domain"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

// NewCartService creates a new CartService
func NewCartService(log logger.Logger, discountPolicy DiscountPolicy, taxPolicy TaxPolicy, policyNames []string) *CartService {
	return &CartService{
		log: log,

		DiscountPolicy: discountPolicy,
		TaxPolicy:      taxPolicy,
		PolicyNames:    policyNames,
	}
}

// CalculateTotal computes the total price, applying discounts and taxes
func (s *CartService) CalculateTotal(ctx context.Context, cart *domain.Cart, discountParams, taxParams map[string]interface{}) (CartTotal, error) {
	var total CartTotal

	// Evaluate Discount Policy
	s.log.InfoWithContext(ctx, "Evaluating Discount Policy for CustomerID: %s", field.Fields{"customerID": cart.CustomerID})
	totalDiscountFloat, err := s.DiscountPolicy.Evaluate(ctx, cart, discountParams)
	if err != nil {
		return total, fmt.Errorf("failed to evaluate discount policy: %w", err)
	}

	s.log.InfoWithContext(ctx, "Total Discount: %.2f", field.Fields{"totalDiscount": totalDiscountFloat})
	totalDiscount := decimal.NewFromFloat(totalDiscountFloat)

	// Evaluate Tax Policy
	s.log.InfoWithContext(ctx, "Evaluating Tax Policy for CustomerID: %s", field.Fields{"customerID": cart.CustomerID})
	totalTaxFloat, err := s.TaxPolicy.Evaluate(ctx, cart, taxParams)
	if err != nil {
		return total, fmt.Errorf("failed to evaluate tax policy: %w", err)
	}

	s.log.InfoWithContext(ctx, "Total Tax: %.2f", field.Fields{"totalTax": totalTaxFloat})
	totalTax := decimal.NewFromFloat(totalTaxFloat)

	// Calculate Final Price
	s.log.InfoWithContext(ctx, "Calculating Final Price for CustomerID: %s", field.Fields{"customerID": cart.CustomerID})
	finalPrice := decimal.Zero
	for _, item := range cart.Items {
		// Calculate per-item total: (Price + Tax - Discount) * Quantity
		itemTotal := item.Price.Add(totalTax).Sub(totalDiscount).Mul(decimal.NewFromInt32(item.Quantity))
		finalPrice = finalPrice.Add(itemTotal)
		s.log.InfoWithContext(ctx, "ItemID: %s, ItemTotal: %.2f", field.Fields{"itemID": item.ProductID, "itemTotal": itemTotal})
	}

	// Prepare the CartTotal
	total = CartTotal{
		TotalTax:      totalTax,
		TotalDiscount: totalDiscount,
		FinalPrice:    finalPrice,
		Policies:      s.PolicyNames,
	}

	s.log.InfoWithContext(ctx, "Final Price for CustomerID %s: %.2f", field.Fields{"customerID": cart.CustomerID, "finalPrice": finalPrice})

	return total, nil
}
