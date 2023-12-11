package ram

import (
	"reflect"
	"strings"

	domain "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
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

// TODO: write generator for protoc
func isFilterSuccess(link *domain.Link, filter *domain.FilterLink) bool { //nolint:gocognit,cyclop // ignore
	// Skip empty filter
	if filter == nil {
		return true
	}

	r := reflect.ValueOf(filter)
	l := reflect.ValueOf(link)

	for _, key := range GetKeys() {
		val, okStringFilterInput := reflect.Indirect(r).FieldByName(key).Interface().(*domain.StringFilterInput)
		if !okStringFilterInput {
			continue
		}

		// Skip empty value
		if val == nil {
			continue
		}

		// Eq
		if val.Eq != "" {
			linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if !ok {
				continue
			}

			if linkValue != val.Eq {
				return false
			}
		}

		// Ne
		if val.Ne != "" {
			linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if !ok {
				continue
			}

			if linkValue == val.Ne {
				return false
			}
		}

		// Lt
		if val.Lt != "" {
			linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if !ok {
				continue
			}

			if linkValue > val.Lt {
				return false
			}
		}

		// Le
		if val.Le != "" {
			linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if !ok {
				continue
			}

			if linkValue >= val.Le {
				return false
			}
		}

		// Gt
		if val.Gt != "" {
			linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if !ok {
				continue
			}

			if linkValue < val.Gt {
				return false
			}
		}

		// Ge
		if val.Ge != "" {
			linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if !ok {
				continue
			}

			if linkValue <= val.Ge {
				return false
			}
		}

		// Contains
		if val.Contains != "" {
			linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if !ok {
				continue
			}

			if !strings.Contains(linkValue, val.Contains) {
				return false
			}
		}

		// NotContains
		if val.Contains != "" {
			linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if !ok {
				continue
			}

			if strings.Contains(linkValue, val.NotContains) {
				return false
			}
		}
	}

	return true
}
