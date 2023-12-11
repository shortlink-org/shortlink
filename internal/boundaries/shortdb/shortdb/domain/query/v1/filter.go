package v1

import (
	"github.com/shortlink-org/shortlink/internal/pkg/types/vector"
)

func Filter[V vector.Type](LValue, RValue V, operator Operator) bool {
	switch operator {
	case Operator_OPERATOR_EQ:
		return LValue == RValue
	case Operator_OPERATOR_GT:
		return LValue > RValue
	case Operator_OPERATOR_GTE:
		return LValue >= RValue
	case Operator_OPERATOR_LT:
		return LValue < RValue
	case Operator_OPERATOR_LTE:
		return LValue <= RValue
	case Operator_OPERATOR_NE:
		return LValue != RValue
	default:
		return false
	}
}

func FilterBool(LValue, RValue bool, operator Operator) bool {
	switch operator {
	case Operator_OPERATOR_EQ:
		return LValue == RValue
	case Operator_OPERATOR_NE:
		return LValue != RValue
	default:
		return false
	}
}
