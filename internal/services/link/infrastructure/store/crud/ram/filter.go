package ram

import (
	"reflect"
	"strings"

	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

func isFilterSuccess(link *v1.Link, filter *query.Filter) bool { // nolint:gocognit
	// Skip empty filter
	if filter == nil {
		return true
	}

	r := reflect.ValueOf(filter)
	l := reflect.ValueOf(link)

	for _, key := range filter.GetKeys() {
		val := reflect.Indirect(r).FieldByName(key).Interface().(*query.StringFilterInput) // nolint:errcheck

		// Skip empty value
		if val == nil {
			continue
		}

		// Eq
		if val.Eq != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) // nolint:errcheck
			if linkValue != *val.Eq {
				return false
			}
		}

		// Ne
		if val.Ne != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) // nolint:errcheck
			if linkValue == *val.Ne {
				return false
			}
		}

		// Lt
		if val.Lt != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) // nolint:errcheck
			if linkValue > *val.Lt {
				return false
			}
		}

		// Le
		if val.Le != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) // nolint:errcheck
			if linkValue >= *val.Le {
				return false
			}
		}

		// Gt
		if val.Gt != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) // nolint:errcheck
			if linkValue < *val.Gt {
				return false
			}
		}

		// Ge
		if val.Ge != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) // nolint:errcheck
			if linkValue <= *val.Ge {
				return false
			}
		}

		// Contains
		if val.Contains != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) // nolint:errcheck
			if !strings.Contains(linkValue, *val.Contains) {
				return false
			}
		}

		// NotContains
		if val.Contains != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) // nolint:errcheck
			if strings.Contains(linkValue, *val.NotContains) {
				return false
			}
		}
	}

	return true
}
