package rules

import "strings"

// NormalizeEmail normalizes an email address for comparison:
// - Converts to lowercase
// - Trims whitespace
// Returns empty string if email is invalid after normalization.
func NormalizeEmail(email string) string {
	normalized := strings.ToLower(strings.TrimSpace(email))
	if normalized == "" {
		return ""
	}

	return normalized
}

