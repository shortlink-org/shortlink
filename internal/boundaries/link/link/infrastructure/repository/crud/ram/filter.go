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
func isFilterSuccess(link *domain.Link, filter *domain.FilterLink) bool {
	if filter == nil {
		return true
	}

	r := reflect.ValueOf(filter)
	l := reflect.ValueOf(link)

	for _, key := range GetKeys() {
		val, ok := reflect.Indirect(r).FieldByName(key).Interface().(*domain.StringFilterInput)
		if !ok || val == nil {
			continue
		}

		linkValue, ok := reflect.Indirect(l).FieldByName(key).Interface().(string)
		if !ok {
			continue
		}

		if !isStringFilterMatch(linkValue, val) {
			return false
		}
	}

	return true
}

func isStringFilterMatch(linkValue string, filter *domain.StringFilterInput) bool {
	if filter.Eq != "" && linkValue != filter.Eq {
		return false
	}

	if filter.Ne != "" && linkValue == filter.Ne {
		return false
	}

	if filter.Lt != "" && !(linkValue < filter.Lt) {
		return false
	}

	if filter.Le != "" && !(linkValue <= filter.Le) {
		return false
	}

	if filter.Gt != "" && !(linkValue > filter.Gt) {
		return false
	}

	if filter.Ge != "" && !(linkValue >= filter.Ge) {
		return false
	}

	for _, c := range filter.Contains {
		if c != "" && !strings.Contains(linkValue, c) {
			return false
		}
	}

	for _, nc := range filter.NotContains {
		if nc != "" && strings.Contains(linkValue, nc) {
			return false
		}
	}

	return true
}
