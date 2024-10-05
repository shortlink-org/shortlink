package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v2"
)

type Goods struct {
	Brand    string                 `json:"brand"`
	Model    string                 `json:"model"`
	Price    decimal.Decimal        `json:"price"`
	Stock    int                    `json:"stock"`
	Category string                 `json:"category"`
	Tags     []string               `json:"tags"`
	Features map[string]interface{} `json:"features"`
}

type Policy struct {
	Name        string                 `yaml:"name"`
	Operator    string                 `yaml:"operator"`
	Description string                 `yaml:"description"`
	Rules       []Rule                 `yaml:"rules"`
	Params      map[string]interface{} `yaml:"params"`
}

type Rule struct {
	Name      string `yaml:"name"`
	Condition string `yaml:"condition"`
	Message   string `yaml:"message"`
	Action    string `yaml:"action"`
}

func main() {
	// Load fixtures
	fixtures := loadFixtures("tests/fixtures/phone.json")

	// Load policies
	policyFiles, _ := filepath.Glob("feeds/*.yaml")

	// Process each policy
	for _, policyFile := range policyFiles {
		policy := loadPolicy(policyFile)

		// Filter goods based on policy rules
		filteredGoods := applyPolicy(policy, fixtures)

		// Generate output XML
		if len(filteredGoods) > 0 {
			// Ensure the output directory exists
			os.MkdirAll("out", os.ModePerm)

			outputFileName := fmt.Sprintf("out/feed_%s.xml", strings.TrimSuffix(filepath.Base(policyFile), ".yaml"))
			generateXML(filteredGoods, outputFileName)
		} else {
			fmt.Printf("No goods passed the policy %s, skipping XML generation.\n", policyFile)
		}
	}
}

func loadFixtures(filePath string) []Goods {
	var goods []Goods
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading fixtures:", err)
		return nil
	}
	err = json.Unmarshal(data, &goods)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}
	return goods
}

func loadPolicy(filePath string) Policy {
	var policy Policy
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading policy:", err)
		return policy
	}
	err = yaml.Unmarshal(data, &policy)
	if err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		return policy
	}
	return policy
}

func applyPolicy(policy Policy, goods []Goods) []Goods {
	var filteredGoods []Goods
	for _, item := range goods {
		if evaluatePolicy(policy, item) {
			filteredGoods = append(filteredGoods, item)
		}
	}
	return filteredGoods
}

func evaluatePolicy(policy Policy, goods Goods) bool {
	// Create a new CEL environment
	env, err := cel.NewEnv(
	// Declarations as before...
	)
	if err != nil {
		fmt.Println("Error creating CEL environment:", err)
		return false
	}

	activation := make(map[string]interface{})

	// Set goods attributes
	activation["brand"] = goods.Brand
	activation["model"] = goods.Model
	activation["price"], _ = goods.Price.Float64()
	activation["stock"] = goods.Stock
	activation["category"] = goods.Category
	activation["tags"] = goods.Tags
	activation["features"] = goods.Features

	// Set policy parameters
	for k, v := range policy.Params {
		activation[k] = v
	}

	// Evaluate each rule
	results := make([]bool, len(policy.Rules))
	for i, rule := range policy.Rules {
		// Parse and check the expression
		ast, issues := env.Parse(rule.Condition)
		if issues != nil && issues.Err() != nil {
			fmt.Printf("Error parsing condition '%s': %v\n", rule.Condition, issues.Err())
			return false
		}

		prg, err := env.Program(ast)
		if err != nil {
			fmt.Printf("Error compiling condition '%s': %v\n", rule.Condition, err)
			return false
		}

		// Evaluate the expression
		out, _, err := prg.Eval(activation)
		if err != nil {
			fmt.Printf("Error evaluating condition '%s': %v\n", rule.Condition, err)
			return false
		}

		// Check the result
		result, ok := out.Value().(bool)
		if !ok {
			fmt.Printf("Condition '%s' did not return a boolean value\n", rule.Condition)
			return false
		}

		results[i] = result
	}

	// Combine results based on the operator
	operator := "OR"
	if policy.Operator != "" {
		operator = strings.ToUpper(policy.Operator)
	}
	if operator == "AND" {
		for _, res := range results {
			if !res {
				return false
			}
		}
		return true
	} else { // Default to OR
		for _, res := range results {
			if res {
				return true
			}
		}
		return false
	}
}

// Define a Feature struct for XML encoding
type Feature struct {
	XMLName xml.Name `xml:"feature"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

// Define XMLGoods struct for XML encoding
type XMLGoods struct {
	XMLName  xml.Name  `xml:"goods"`
	Brand    string    `xml:"brand"`
	Model    string    `xml:"model"`
	Price    string    `xml:"price"`
	Stock    int       `xml:"stock"`
	Category string    `xml:"category"`
	Tags     []string  `xml:"tags>tag"`
	Features []Feature `xml:"features>feature"`
}

func generateXML(goods []Goods, filePath string) {
	var xmlGoodsList []XMLGoods
	for _, g := range goods {
		// Collect and sort feature names
		featureNames := make([]string, 0, len(g.Features))
		for k := range g.Features {
			featureNames = append(featureNames, k)
		}
		sort.Strings(featureNames)

		// Build Features slice in sorted order
		features := make([]Feature, 0, len(featureNames))
		for _, k := range featureNames {
			v := g.Features[k]
			valueStr := fmt.Sprintf("%v", v)
			features = append(features, Feature{
				Name:  k,
				Value: valueStr,
			})
		}

		// Optionally, sort tags if needed
		sort.Strings(g.Tags)

		xmlGoods := XMLGoods{
			Brand:    g.Brand,
			Model:    g.Model,
			Price:    g.Price.StringFixed(2), // Convert decimal to string with 2 decimal places
			Stock:    g.Stock,
			Category: g.Category,
			Tags:     g.Tags,
			Features: features,
		}
		xmlGoodsList = append(xmlGoodsList, xmlGoods)
	}

	feed := struct {
		XMLName xml.Name   `xml:"feed"`
		Goods   []XMLGoods `xml:"goods"`
	}{
		Goods: xmlGoodsList,
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	err = encoder.Encode(feed)
	if err != nil {
		fmt.Println("Error encoding XML:", err)
	}
}
