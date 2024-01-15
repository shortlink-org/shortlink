package v1

import (
	"github.com/shortlink-org/shortlink/internal/pkg/types/vector"
)

func Filter[V vector.Type](lValue V, rValue V, operator Operator) bool {
	switch operator {
	case Operator_OPERATOR_EQ:
		return lValue == rValue
	case Operator_OPERATOR_GT:
		return lValue > rValue
	case Operator_OPERATOR_GTE:
		return lValue >= rValue
	case Operator_OPERATOR_LT:
		return lValue < rValue
	case Operator_OPERATOR_LTE:
		return lValue <= rValue
	case Operator_OPERATOR_NE:
		return lValue != rValue
	default:
		return false
	}
}

func FilterBool(lValue, rValue bool, operator Operator) bool {
	switch operator {
	case Operator_OPERATOR_EQ:
		return lValue == rValue
	case Operator_OPERATOR_NE:
		return lValue != rValue
	default:
		return false
	}
}
