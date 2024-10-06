package policy_evaluator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/open-policy-agent/opa/rego"

	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/domain"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

// PolicyEvaluator interface as defined
type PolicyEvaluator interface {
	Evaluate(ctx context.Context, cart *domain.Cart, params map[string]interface{}) (float64, error)
}

// OPAEvaluator implements the PolicyEvaluator interface using OPA's rego package
type OPAEvaluator struct {
	preparedQuery rego.PreparedEvalQuery
	query         string
	policyPath    string
}

func NewOPAEvaluator(log logger.Logger, policyPath string, query string) (*OPAEvaluator, error) {
	// Log the policy path and query
	log.Info("Initializing OPAEvaluator with Policy Path: %s and Query: %s", field.Fields{"policyPath": policyPath, "query": query})

	// Check if the policy directory exists
	if _, err := os.Stat(policyPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("policy directory does not exist: %s", policyPath)
	}

	// Prepare the query
	r := rego.New(
		rego.Query(query),
		rego.Load([]string{policyPath}, nil),
	)

	preparedQuery, err := r.PrepareForEval(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to prepare OPA query: %w", err)
	}

	return &OPAEvaluator{
		preparedQuery: preparedQuery,
		query:         query,
		policyPath:    policyPath,
	}, nil
}

// Evaluate executes the OPA policy against the provided cart and parameters
func (e *OPAEvaluator) Evaluate(ctx context.Context, cart *domain.Cart, params map[string]interface{}) (float64, error) {
	// Transform Cart to OPA input
	input := transformCartToInput(cart, params)

	// Evaluate the policy
	rs, err := e.preparedQuery.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return 0.0, fmt.Errorf("OPA evaluation error: %w", err)
	}

	if len(rs) == 0 {
		return 0.0, nil // No result from policy
	}

	// Assuming the policy returns a single value
	expr := rs[0].Expressions[0].Value
	result, err := parseOPAResult(expr)
	if err != nil {
		return 0.0, err
	}

	return result, nil
}

// transformCartToInput converts the domain.Cart to the input format expected by OPA
func transformCartToInput(cart *domain.Cart, params map[string]interface{}) map[string]interface{} {
	var items []map[string]interface{}
	for _, item := range cart.Items {
		items = append(items, map[string]interface{}{
			"productId": item.ProductID.String(), // Convert UUID to string
			"quantity":  item.Quantity,
			"price":     item.Price.InexactFloat64(), // Convert decimal to float64
			"brand":     item.Brand,
		})
	}

	return map[string]interface{}{
		"items":  items,
		"params": params, // Include additional parameters if needed
	}
}

// parseOPAResult handles different types that OPA might return and converts them to float64
func parseOPAResult(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case string:
		// Attempt to parse string to float64
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0.0, fmt.Errorf("invalid string format for result: %v", err)
		}
		return parsed, nil
	case json.Number:
		parsed, err := v.Float64()
		if err != nil {
			return 0.0, fmt.Errorf("invalid json.Number format for result: %v", err)
		}
		return parsed, nil
	default:
		return 0.0, fmt.Errorf("unexpected type for result: %T", v)
	}
}

// GetPolicyNames retrieves the names of all .rego files in the specified directories.
func GetPolicyNames(dirs ...string) ([]string, error) {
	var policyNames []string
	for _, dir := range dirs {
		// Use filepath.Glob to find all .rego files in the directory
		pattern := filepath.Join(dir, "*.rego")
		files, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("failed to list .rego files in %s: %v", dir, err)
		}

		for _, file := range files {
			// Extract the base name without the directory and extension
			base := filepath.Base(file)
			name := base[:len(base)-len(filepath.Ext(base))]
			policyNames = append(policyNames, name)
		}
	}
	return policyNames, nil
}
