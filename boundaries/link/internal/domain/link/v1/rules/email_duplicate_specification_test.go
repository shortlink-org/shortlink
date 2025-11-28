package rules

import (
	"testing"
)

func TestEmailDuplicateSpecification(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		data      *EmailValidationData
		wantError bool
	}{
		{
			name:  "unique email",
			email: "luna.smith@example.com",
			data: &EmailValidationData{
				Email:      "luna.smith@example.com",
				SeenEmails: map[string]bool{},
			},
			wantError: false,
		},
		{
			name:  "unique email with different case",
			email: "Nova.Ray@Galaxy.Net",
			data: &EmailValidationData{
				Email:      "Nova.Ray@Galaxy.Net",
				SeenEmails: map[string]bool{"nova.ray@galaxy.net": true},
			},
			wantError: true,
		},
		{
			name:  "duplicate email",
			email: "blaze.rivera@orbit.dev",
			data: &EmailValidationData{
				Email:      "blaze.rivera@orbit.dev",
				SeenEmails: map[string]bool{"blaze.rivera@orbit.dev": true},
			},
			wantError: true,
		},
		{
			name:  "duplicate email with spaces",
			email: "   echo.light@nebula.sh   ",
			data: &EmailValidationData{
				Email:      "   echo.light@nebula.sh   ",
				SeenEmails: map[string]bool{"echo.light@nebula.sh": true},
			},
			wantError: true,
		},
		{
			name:  "empty email",
			email: "",
			data: &EmailValidationData{
				Email:      "",
				SeenEmails: map[string]bool{},
			},
			wantError: true,
		},
		{
			name:  "multiple seen emails, unique new email",
			email: "sol.zen@lumen.dev",
			data: &EmailValidationData{
				Email: "sol.zen@lumen.dev",
				SeenEmails: map[string]bool{
					"mira.fernandez@orbit.dev": true,
					"kai.jordan@zeta.io":       true,
				},
			},
			wantError: false,
		},
		{
			name:  "nil seen emails",
			email: "terra.kade@starhub.co",
			data: &EmailValidationData{
				Email:      "terra.kade@starhub.co",
				SeenEmails: nil,
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := NewEmailDuplicateSpecification(tt.email)
			err := spec.IsSatisfiedBy(tt.data)

			if (err != nil) != tt.wantError {
				t.Errorf("IsSatisfiedBy() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}
