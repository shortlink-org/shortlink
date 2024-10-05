package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/open-policy-agent/opa/rego"
)

// Config represents the structure of the config.yaml file
type Config struct {
	Params map[string]interface{} `yaml:"params"`
}

// Goods represents a product with its attributes
type Goods struct {
	Brand    string                 `json:"brand"`
	Model    string                 `json:"model"`
	Price    float64                `json:"price"`
	Stock    int                    `json:"stock"`
	Category string                 `json:"category"`
	Tags     []string               `json:"tags"`
	Features map[string]interface{} `json:"features"`
}

// LoadConfig loads the configuration from config.yaml
func LoadConfig(filePath string) (Config, error) {
	var config Config
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	return config, err
}

// LoadGoods loads the goods data from tests/fixtures/phone.json
func LoadGoods(filePath string) ([]Goods, error) {
	var goods []Goods
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &goods)
	return goods, err
}

// LoadOPAPolicies loads all OPA policies from a directory
func LoadOPAPolicies(dir string) (*rego.PreparedEvalQuery, error) {
	// Load all rego files from the directory
	query, err := rego.New(
		rego.Query("data.pricing"),
		rego.Load([]string{dir}, nil),
	).PrepareForEval(context.TODO())
	if err != nil {
		return nil, err
	}
	return &query, nil
}

// EvaluateTaxPolicy evaluates tax policies using OPA
func EvaluateTaxPolicy(goods Goods, params map[string]interface{}, query *rego.PreparedEvalQuery) (float64, error) {
	// Create input data
	input := map[string]interface{}{
		"price":  goods.Price,
		"params": params,
	}

	// Evaluate OPA policies
	rs, err := query.Eval(context.TODO(), rego.EvalInput(input))
	if err != nil {
		return 0, err
	}

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

// EvaluateDiscountPolicy evaluates discount policies using OPA
func EvaluateDiscountPolicy(goods Goods, params map[string]interface{}, query *rego.PreparedEvalQuery) (float64, error) {
	// Create input data with current time and stock count
	currentTime := time.Now().Format("15:04")
	input := map[string]interface{}{
		"price":  goods.Price,
		"brand":  goods.Brand,
		"count":  goods.Stock,
		"time":   currentTime,
		"params": params,
	}

	// Evaluate OPA policies
	rs, err := query.Eval(context.TODO(), rego.EvalInput(input))
	if err != nil {
		return 0, err
	}

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
	// Load parameters from config.yaml
	config, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Load goods data from tests/fixtures/phone.json
	goodsList, err := LoadGoods("tests/fixtures/phone.json")
	if err != nil {
		log.Fatalf("Failed to load goods: %v", err)
	}

	// Load policies
	taxQuery, err := LoadOPAPolicies("policies/taxes/")
	if err != nil {
		log.Fatalf("Failed to load tax policies: %v", err)
	}
	discountQuery, err := LoadOPAPolicies("policies/discounts/")
	if err != nil {
		log.Fatalf("Failed to load discount policies: %v", err)
	}

	// Iterate over the list of goods and calculate prices
	for _, goods := range goodsList {
		// Calculate taxes and discounts
		tax, err := EvaluateTaxPolicy(goods, config.Params, taxQuery)
		if err != nil {
			log.Fatalf("Failed to evaluate tax: %v", err)
		}
		discount, err := EvaluateDiscountPolicy(goods, config.Params, discountQuery)
		if err != nil {
			log.Fatalf("Failed to evaluate discount: %v", err)
		}

		// Final price calculation
		finalPrice := (goods.Price + tax) - discount

		// Prepare the result
		result := map[string]interface{}{
			"brand":          goods.Brand,
			"model":          goods.Model,
			"category":       goods.Category,
			"original_price": goods.Price,
			"tax":            tax,
			"discount":       discount,
			"final_price":    finalPrice,
		}

		// Save result to the out folder, one file per goods item
		filename := fmt.Sprintf("%s_%s_price.json", goods.Brand, goods.Model)
		err = SaveResultToFile(result, "out", filename)
		if err != nil {
			log.Fatalf("Failed to save result: %v", err)
		}
	}
}
