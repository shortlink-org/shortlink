package rules

import (
	"fmt"
	"testing"

	"github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/errors"
)

// helper to generate deterministic but varied emails
func genEmails(n int) []string {
	out := make([]string, n)
	for i := 0; i < n; i++ {
		out[i] = fmt.Sprintf("user%d@example.com", i)
	}
	return out
}

func TestAllowlistSizeSpecification(t *testing.T) {
	tests := []struct {
		name      string
		emails    []string
		wantError bool
	}{
		{
			name:      "empty allowlist",
			emails:    []string{},
			wantError: false,
		},
		{
			name:      "allowlist within limit",
			emails:    []string{"luna.smith@example.com", "nova.kent@orbit.dev"},
			wantError: false,
		},
		{
			name:      "allowlist at limit",
			emails:    genEmails(errors.MaxAllowlistSize),
			wantError: false,
		},
		{
			name:      "allowlist exceeds limit",
			emails:    genEmails(errors.MaxAllowlistSize + 1),
			wantError: true,
		},
		{
			name:      "allowlist exceeds limit by 10",
			emails:    genEmails(errors.MaxAllowlistSize + 10),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := NewAllowlistSizeSpecification()
			err := spec.IsSatisfiedBy(&tt.emails)

			if (err != nil) != tt.wantError {
				t.Errorf("IsSatisfiedBy() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}
