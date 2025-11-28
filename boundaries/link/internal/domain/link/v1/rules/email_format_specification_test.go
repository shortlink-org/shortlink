package rules

import (
	"testing"
)

func TestEmailFormatSpecification(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		data      *EmailValidationData
		wantError bool
	}{
		{
			name:      "valid email",
			email:     "luna.smith@example.com",
			data:      &EmailValidationData{Email: "luna.smith@example.com"},
			wantError: false,
		},
		{
			name:      "valid email with plus",
			email:     "nova.ray+tag@galaxy.net",
			data:      &EmailValidationData{Email: "nova.ray+tag@galaxy.net"},
			wantError: false,
		},
		{
			name:      "valid email with subdomain",
			email:     "kai.jordan@mail.orbit.dev",
			data:      &EmailValidationData{Email: "kai.jordan@mail.orbit.dev"},
			wantError: false,
		},
		{
			name:      "invalid email - no @",
			email:     "invalidemail.com",
			data:      &EmailValidationData{Email: "invalidemail.com"},
			wantError: true,
		},
		{
			name:      "invalid email - no domain",
			email:     "invalid@",
			data:      &EmailValidationData{Email: "invalid@"},
			wantError: true,
		},
		{
			name:      "invalid email - no local part",
			email:     "@example.com",
			data:      &EmailValidationData{Email: "@example.com"},
			wantError: true,
		},
		{
			name:      "empty email",
			email:     "",
			data:      &EmailValidationData{Email: ""},
			wantError: true,
		},
		{
			name:      "email with spaces",
			email:     "   echo.light@nebula.sh   ",
			data:      &EmailValidationData{Email: "   echo.light@nebula.sh   "},
			wantError: false,
		},
		{
			name:      "email with uppercase",
			email:     "BLAZE.RIVERA@ORBIT.DEV",
			data:      &EmailValidationData{Email: "BLAZE.RIVERA@ORBIT.DEV"},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := NewEmailFormatSpecification(tt.email)
			err := spec.IsSatisfiedBy(tt.data)

			if (err != nil) != tt.wantError {
				t.Errorf("IsSatisfiedBy() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}
