package ram

import (
	"reflect"
	"strings"

	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
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
		val := reflect.Indirect(r).FieldByName(key).Interface().(*domain.StringFilterInput) //nolint:errcheck // ignore

		// Skip empty value
		if val == nil {
			continue
		}

		// Eq
		if val.Eq != "" {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue != val.Eq {
				return false
			}
		}

		// Ne
		if val.Ne != "" {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue == val.Ne {
				return false
			}
		}

		// Lt
		if val.Lt != "" {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue > val.Lt {
				return false
			}
		}

		// Le
		if val.Le != "" {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue >= val.Le {
				return false
			}
		}

		// Gt
		if val.Gt != "" {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue < val.Gt {
				return false
			}
		}

		// Ge
		if val.Ge != "" {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if linkValue <= val.Ge {
				return false
			}
		}

		// Contains
		if val.Contains != "" {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if !strings.Contains(linkValue, val.Contains) {
				return false
			}
		}

		// NotContains
		if val.Contains != "" {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string) //nolint:errcheck // ignore
			if strings.Contains(linkValue, val.NotContains) {
				return false
			}
		}
	}

	return true
}
