package rules

import (
	"testing"
)

func TestEmailNonEmptySpecification(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		data      *EmailValidationData
		wantError bool
	}{
		{
			name:      "non-empty email",
			email:     "luna.smith@example.com",
			data:      &EmailValidationData{Email: "luna.smith@example.com"},
			wantError: false,
		},
		{
			name:      "empty email",
			email:     "",
			data:      &EmailValidationData{Email: ""},
			wantError: true,
		},
		{
			name:      "whitespace only",
			email:     "   ",
			data:      &EmailValidationData{Email: "   "},
			wantError: true,
		},
		{
			name:      "email with spaces that normalizes to non-empty",
			email:     "   nova.ray@galaxy.net   ",
			data:      &EmailValidationData{Email: "   nova.ray@galaxy.net   "},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := NewEmailNonEmptySpecification(tt.email)
			err := spec.IsSatisfiedBy(tt.data)

			if (err != nil) != tt.wantError {
				t.Errorf("IsSatisfiedBy() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}
