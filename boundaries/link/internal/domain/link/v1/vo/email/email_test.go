//go:build unit

package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantValue   string
		wantError   bool
		description string
	}{
		{
			name:        "valid email",
			input:       "user@example.com",
			wantValue:   "user@example.com",
			wantError:   false,
			description: "valid email should be normalized",
		},
		{
			name:        "email with uppercase",
			input:       "User@Example.com",
			wantValue:   "user@example.com",
			wantError:   false,
			description: "email should be converted to lowercase",
		},
		{
			name:        "email with spaces",
			input:       "  user@example.com  ",
			wantValue:   "user@example.com",
			wantError:   false,
			description: "spaces should be trimmed",
		},
		{
			name:        "email with uppercase and spaces",
			input:       "  User@Example.com  ",
			wantValue:   "user@example.com",
			wantError:   false,
			description: "both uppercase and spaces should be normalized",
		},
		{
			name:        "empty string",
			input:       "",
			wantValue:   "",
			wantError:   true,
			description: "empty string should return error",
		},
		{
			name:        "only spaces",
			input:       "   ",
			wantValue:   "",
			wantError:   true,
			description: "only spaces should return error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := NewEmail(tt.input)

			if tt.wantError {
				require.Error(t, err, "Expected error but got none")
				assert.Equal(t, Email(""), email, "Email should be empty on error")
			} else {
				require.NoError(t, err, "Unexpected error: %v", err)
				assert.Equal(t, tt.wantValue, email.String(), tt.description)
			}
		})
	}
}

func TestEmail_String(t *testing.T) {
	t.Run("returns normalized value", func(t *testing.T) {
		email, err := NewEmail("User@Example.com")
		require.NoError(t, err)
		assert.Equal(t, "user@example.com", email.String())
	})

	t.Run("returns empty string for empty email", func(t *testing.T) {
		var email Email
		assert.Equal(t, "", email.String())
	})
}

func TestEmail_Value(t *testing.T) {
	t.Run("returns same as String", func(t *testing.T) {
		email, err := NewEmail("user@example.com")
		require.NoError(t, err)
		assert.Equal(t, email.String(), email.Value())
	})
}

func TestEmail_Equals(t *testing.T) {
	tests := []struct {
		name      string
		email1    string
		email2    string
		wantEqual bool
	}{
		{
			name:      "same normalized emails",
			email1:    "user@example.com",
			email2:    "user@example.com",
			wantEqual: true,
		},
		{
			name:      "different case",
			email1:    "User@Example.com",
			email2:    "user@example.com",
			wantEqual: true,
		},
		{
			name:      "different emails",
			email1:    "user1@example.com",
			email2:    "user2@example.com",
			wantEqual: false,
		},
		{
			name:      "both nil",
			email1:    "",
			email2:    "",
			wantEqual: true,
		},
		{
			name:      "one nil",
			email1:    "user@example.com",
			email2:    "",
			wantEqual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var email1, email2 Email
			var err error

			if tt.email1 != "" {
				email1, err = NewEmail(tt.email1)
				require.NoError(t, err)
			}

			if tt.email2 != "" {
				email2, err = NewEmail(tt.email2)
				require.NoError(t, err)
			}

			assert.Equal(t, tt.wantEqual, email1.Equals(email2))
		})
	}
}

func TestEmail_IsEmpty(t *testing.T) {
	t.Run("empty email is empty", func(t *testing.T) {
		var email Email
		assert.True(t, email.IsEmpty())
	})

	t.Run("valid email is not empty", func(t *testing.T) {
		email, err := NewEmail("user@example.com")
		require.NoError(t, err)
		assert.False(t, email.IsEmpty())
	})
}

func TestEmail_Normalization(t *testing.T) {
	t.Run("normalization is consistent", func(t *testing.T) {
		inputs := []string{
			"User@Example.com",
			"  user@example.com  ",
			"USER@EXAMPLE.COM",
			"  User@Example.com  ",
		}

		var normalized Email
		for i, input := range inputs {
			email, err := NewEmail(input)
			require.NoError(t, err, "Input %d: %s", i, input)

			if normalized == "" {
				normalized = email
			} else {
				assert.True(t, normalized.Equals(email), "All variations should normalize to same value")
			}
		}

		assert.Equal(t, "user@example.com", normalized.String())
	})
}

