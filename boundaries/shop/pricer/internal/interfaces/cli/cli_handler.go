package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/application"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/domain"
)

// CLIHandler handles command-line interactions
type CLIHandler struct {
	CartService *application.CartService
	OutputDir   string
}

// Run processes a single cart file with provided parameters
func (h *CLIHandler) Run(cartFile string, discountParams, taxParams map[string]interface{}) error {
	// Load the cart
	cart, err := loadCart(cartFile)
	if err != nil {
		return fmt.Errorf("failed to load cart %s: %w", cartFile, err)
	}

	// Calculate totals
	total, err := h.CartService.CalculateTotal(context.Background(), &cart, discountParams, taxParams)
	if err != nil {
		return fmt.Errorf("failed to calculate total for cart %s: %w", cartFile, err)
	}

	// Prepare the result map
	result := map[string]interface{}{
		"customerId":    cart.CustomerID.String(),
		"totalTax":      total.TotalTax.StringFixed(2),
		"totalDiscount": total.TotalDiscount.StringFixed(2),
		"finalPrice":    total.FinalPrice.StringFixed(2),
		"policies":      total.Policies,
	}

	// Save the result
	filename := fmt.Sprintf("cart_result_%s.json", cart.CustomerID.String())
	if err := saveResultToFile(result, h.OutputDir, filename); err != nil {
		return fmt.Errorf("failed to save result for cart %s: %w", cartFile, err)
	}

	fmt.Printf("Final result saved to %s\n", filepath.Join(h.OutputDir, filename))
	return nil
}

// loadCart reads and unmarshals the cart JSON file
func loadCart(filePath string) (domain.Cart, error) {
	var cart domain.Cart
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cart, err
	}
	err = json.Unmarshal(file, &cart)
	return cart, err
}

// saveResultToFile marshals the result to JSON and writes it to a file
func saveResultToFile(result map[string]interface{}, outDir string, filename string) error {
	// Ensure the output directory exists
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
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
	if err := ioutil.WriteFile(outputFile, data, 0644); err != nil {
		return err
	}

	return nil
}
