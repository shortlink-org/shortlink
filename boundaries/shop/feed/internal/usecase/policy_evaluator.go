package usecase

import (
	"fmt"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"

	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/domain/entity"
)

type Policy struct {
	Name     string                 `yaml:"name"`
	Operator string                 `yaml:"operator"`
	Rules    []Rule                 `yaml:"rules"`
	Params   map[string]interface{} `yaml:"params"`
}

type Rule struct {
	Name      string `yaml:"name"`
	Condition string `yaml:"condition"`
	Message   string `yaml:"message"`
	Action    string `yaml:"action"`
}

func EvaluatePolicy(policy Policy, goods entity.Goods) (bool, error) {
	// Create CEL environment
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewVar("brand", decls.String),
			decls.NewVar("model", decls.String),
			decls.NewVar("price", decls.Double),
			decls.NewVar("stock", decls.Int),
			decls.NewVar("category", decls.String),
			decls.NewVar("tags", decls.NewListType(decls.String)),
			decls.NewVar("features", decls.NewMapType(decls.String, decls.Dyn)),
			// Policy parameters...
		),
	)
	if err != nil {
		return false, fmt.Errorf("error creating CEL environment: %w", err)
	}

	// Prepare activation
	activation := map[string]interface{}{
		"brand":    goods.Brand,
		"model":    goods.Model,
		"price":    goods.Price.InexactFloat64(),
		"stock":    goods.Stock,
		"category": goods.Category,
		"tags":     goods.Tags,
		"features": goods.Features,
	}

	for k, v := range policy.Params {
		activation[k] = v
	}

	results := make([]bool, len(policy.Rules))
	for i, rule := range policy.Rules {
		ast, issues := env.Parse(rule.Condition)
		if issues != nil && issues.Err() != nil {
			return false, fmt.Errorf("error parsing condition '%s': %v", rule.Condition, issues.Err())
		}

		prg, err := env.Program(ast)
		if err != nil {
			return false, fmt.Errorf("error compiling condition '%s': %v", rule.Condition, err)
		}

		out, _, err := prg.Eval(activation)
		if err != nil {
			return false, fmt.Errorf("error evaluating condition '%s': %v", rule.Condition, err)
		}

		result, ok := out.Value().(bool)
		if !ok {
			return false, fmt.Errorf("condition '%s' did not return a boolean value", rule.Condition)
		}

		results[i] = result
	}

	// Combine results based on the operator
	operator := strings.ToUpper(policy.Operator)
	if operator == "" {
		operator = "OR" // Default operator
	}

	switch operator {
	case "AND":
		for _, res := range results {
			if !res {
				return false, nil
			}
		}
		return true, nil
	case "OR":
		for _, res := range results {
			if res {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, fmt.Errorf("unsupported operator '%s'", policy.Operator)
	}
}
