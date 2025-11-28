//go:build unit

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkBuilder_SetAllowedEmails(t *testing.T) {
	tests := []struct {
		name          string
		emails        []string
		wantError     bool
		wantEmails    []string
		errorContains string
	}{
		{
			name:       "empty allowlist creates public link",
			emails:     []string{},
			wantError:  false,
			wantEmails: []string{},
		},
		{
			name:       "nil allowlist creates public link",
			emails:     nil,
			wantError:  false,
			wantEmails: []string{},
		},
		{
			name:       "valid single email",
			emails:     []string{"user@example.com"},
			wantError:  false,
			wantEmails: []string{"user@example.com"},
		},
		{
			name:       "valid multiple emails",
			emails:     []string{"user1@example.com", "user2@example.com"},
			wantError:  false,
			wantEmails: []string{"user1@example.com", "user2@example.com"},
		},
		{
			name:       "emails are normalized",
			emails:     []string{"  User@Example.com  ", "ANOTHER@TEST.ORG"},
			wantError:  false,
			wantEmails: []string{"user@example.com", "another@test.org"},
		},
		{
			name:          "invalid email format",
			emails:        []string{"not-an-email"},
			wantError:     true,
			errorContains: "invalid email",
		},
		{
			name:          "duplicate emails",
			emails:        []string{"user@example.com", "user@example.com"},
			wantError:     true,
			errorContains: "duplicate email",
		},
		{
			name:          "duplicate emails with different case",
			emails:        []string{"User@Example.com", "user@example.com"},
			wantError:     true,
			errorContains: "duplicate email",
		},
		{
			name:          "allowlist too large",
			emails:        make([]string, 101), // MaxAllowlistSize + 1
			wantError:     true,
			errorContains: "allowlist too large",
		},
		{
			name:          "mixed valid and invalid",
			emails:        []string{"valid@example.com", "invalid-email", "another@example.com"},
			wantError:     true,
			errorContains: "invalid email",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewLinkBuilder().
				SetURL("https://example.com").
				SetAllowedEmails(tt.emails)

			link, err := builder.Build()

			if tt.wantError {
				require.Error(t, err, "Expected error but got none")
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains, "Error message should contain expected text")
				}
				assert.Nil(t, link, "Link should be nil on error")
			} else {
				require.NoError(t, err, "Unexpected error: %v", err)
				require.NotNil(t, link, "Link should not be nil")
				assert.Equal(t, tt.wantEmails, link.GetAllowedEmails(), "Allowed emails should match")
			}
		})
	}
}

func TestLinkBuilder_SetAllowedEmails_Integration(t *testing.T) {
	t.Run("can build link with all fields including allowed emails", func(t *testing.T) {
		link, err := NewLinkBuilder().
			SetURL("https://example.com").
			SetDescribe("Test description").
			SetAllowedEmails([]string{"user1@example.com", "user2@example.com"}).
			Build()

		require.NoError(t, err)
		assert.NotEmpty(t, link.GetHash())
		assert.Equal(t, "https://example.com", link.GetUrl().String())
		assert.Equal(t, "Test description", link.GetDescribe())
		assert.Equal(t, []string{"user1@example.com", "user2@example.com"}, link.GetAllowedEmails())
		assert.False(t, link.IsPublic())
	})

	t.Run("can build public link without allowed emails", func(t *testing.T) {
		link, err := NewLinkBuilder().
			SetURL("https://example.com").
			SetDescribe("Public link").
			SetAllowedEmails([]string{}).
			Build()

		require.NoError(t, err)
		assert.True(t, link.IsPublic())
		assert.Empty(t, link.GetAllowedEmails())
	})
}

