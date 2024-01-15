package v1

import (
	"strconv"

	field "github.com/shortlink-org/shortlink/internal/boundaries/shortdb/shortdb/domain/field/v1"
	page "github.com/shortlink-org/shortlink/internal/boundaries/shortdb/shortdb/domain/page/v1"
)

func (q *Query) IsLimit() bool {
	return q.GetLimit() != 0
}

func (q *Query) IsFilter(record *page.Row, fields map[string]field.Type) bool {
	for _, condition := range q.GetConditions() {
		var err error
		var LValue any
		var RValue any

		payload := record.GetValue()[condition.GetLValue()]
		switch fields[condition.GetLValue()] {
		case field.Type_TYPE_INTEGER:
			LValue, err = strconv.Atoi(string(payload))
			if err != nil {
				return false
			}
			RValue, err = strconv.Atoi(condition.GetRValue())
			if err != nil {
				return false
			}

			return Filter(LValue.(int), RValue.(int), condition.GetOperator()) //nolint:forcetypeassert // simple type assertion
		case field.Type_TYPE_STRING:
			LValue = string(payload)
			return Filter(LValue.(string), condition.GetRValue(), condition.GetOperator()) //nolint:forcetypeassert // simple type assertion
		case field.Type_TYPE_BOOLEAN:
			LValue, err = strconv.ParseBool(string(payload))
			if err != nil {
				return false
			}
			RValue, err = strconv.ParseBool(condition.GetRValue())
			if err != nil {
				return false
			}

			return FilterBool(LValue.(bool), RValue.(bool), condition.GetOperator()) //nolint:forcetypeassert // simple type assertion
		case field.Type_TYPE_UNSPECIFIED:
			return false
		default:
			return false
		}
	}

	return true
}
