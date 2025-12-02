package rules

import (
	"testing"

	"github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/vo/email"
)

func TestValidateEmailAllowlist(t *testing.T) {
	tests := []struct {
		name           string
		emails         []string
		wantNormalized []string
		wantError      bool
	}{
		{
			name:           "empty allowlist",
			emails:         []string{},
			wantNormalized: []string{},
			wantError:      false,
		},
		{
			name:           "valid single email",
			emails:         []string{"luna.smith@example.com"},
			wantNormalized: []string{"luna.smith@example.com"},
			wantError:      false,
		},
		{
			name:           "valid multiple emails",
			emails:         []string{"kai.jordan@zeta.io", "mira.fernandez@orbit.dev"},
			wantNormalized: []string{"kai.jordan@zeta.io", "mira.fernandez@orbit.dev"},
			wantError:      false,
		},
		{
			name:           "emails with spaces and uppercase",
			emails:         []string{"  Blaze.RiverA@Example.org ", "SKY.NOVA@GALAXY.NET  "},
			wantNormalized: []string{"blaze.rivera@example.org", "sky.nova@galaxy.net"},
			wantError:      false,
		},
		{
			name:      "duplicate emails",
			emails:    []string{"echo.wright@neon.sh", "echo.wright@neon.sh"},
			wantError: true,
		},
		{
			name:      "duplicate emails with different case",
			emails:    []string{"orion.vega@stellar.io", "ORION.VEGA@STELLAR.IO"},
			wantError: true,
		},
		{
			name:      "invalid email format",
			emails:    []string{"not-an-email-123"},
			wantError: true,
		},
		{
			name:      "allowlist exceeds limit",
			emails:    make([]string, email.MaxAllowlistSize+1),
			wantError: true,
		},
		{
			name:           "mixed valid and invalid - should fail on first invalid",
			emails:         []string{"ion.kade@dune.app", "invalid!!", "nova.ray@lumen.dev"},
			wantNormalized: nil,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalized, err := ValidateEmailAllowlist(tt.emails)

			if (err != nil) != tt.wantError {
				t.Errorf("ValidateEmailAllowlist() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError {
				if len(normalized) != len(tt.wantNormalized) {
					t.Errorf("ValidateEmailAllowlist() normalized length = %v, want %v", len(normalized), len(tt.wantNormalized))
					return
				}

				for i := range normalized {
					if normalized[i] != tt.wantNormalized[i] {
						t.Errorf("ValidateEmailAllowlist() normalized[%d] = %v, want %v", i, normalized[i], tt.wantNormalized[i])
					}
				}
			}
		})
	}
}
