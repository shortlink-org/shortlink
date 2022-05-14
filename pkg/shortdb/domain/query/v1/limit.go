package v1

import (
	"strconv"

	v1 "github.com/batazor/shortlink/pkg/shortdb/domain/table/v1"
)

func (q *Query) IsLimit() bool {
	return q.Limit != 0
}

func (q *Query) IsFilter(record *v1.Row, fields map[string]v1.Type) bool {
	for _, condition := range q.Conditions {
		var err error
		var LValue interface{}
		var RValue interface{}

		payload := record.Value[condition.LValue]
		switch fields[condition.LValue] {
		case v1.Type_TYPE_INTEGER:
			LValue, err = strconv.Atoi(string(payload))
			if err != nil {
				return false
			}
			RValue, err = strconv.Atoi(condition.RValue)
			if err != nil {
				return false
			}

			return Filter(LValue.(int), RValue.(int), condition.Operator)
		case v1.Type_TYPE_STRING:
			LValue = string(payload)
			return Filter(LValue.(string), condition.RValue, condition.Operator)
		case v1.Type_TYPE_BOOLEAN:
			LValue, err = strconv.ParseBool(string(payload))
			if err != nil {
				return false
			}
			RValue, err = strconv.ParseBool(condition.RValue)
			if err != nil {
				return false
			}
			return FilterBool(LValue.(bool), RValue.(bool), condition.Operator)
		case v1.Type_TYPE_UNSPECIFIED:
			fallthrough
		default:
			return false
		}
	}

	return true
}
