package ram

import (
	"reflect"
	"strings"

	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
)

// TODO: fix this hardcode
func GetKeys() []string {
	return []string{
		"Field_mask",
		"Url",
		"Hash",
		"Describe",
		"Created_at",
		"Updated_at",
		"Link",
	}
}

func isFilterSuccess(link *domain.Link, filter *domain.FilterLink) bool { //nolint:gocognit // ignore
	// Skip empty filter
	if filter == nil {
		return true
	}

	r := reflect.ValueOf(filter)
	l := reflect.ValueOf(link)

	for _, key := range GetKeys() {
		val := reflect.Indirect(r).FieldByName(key).Interface().(*query.StringFilterInput) //nolint:errcheck // ignore

		// Skip empty value
		if val == nil {
			continue
		}

		// Eq
		if val.Eq != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue != *val.Eq {
				return false
			}
		}

		// Ne
		if val.Ne != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue == *val.Ne {
				return false
			}
		}

		// Lt
		if val.Lt != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue > *val.Lt {
				return false
			}
		}

		// Le
		if val.Le != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue >= *val.Le {
				return false
			}
		}

		// Gt
		if val.Gt != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue < *val.Gt {
				return false
			}
		}

		// Ge
		if val.Ge != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue <= *val.Ge {
				return false
			}
		}

		// Contains
		if val.Contains != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if !strings.Contains(linkValue, *val.Contains) {
				return false
			}
		}

		// NotContains
		if val.Contains != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if strings.Contains(linkValue, *val.NotContains) {
				return false
			}
		}
	}

	return true
}
