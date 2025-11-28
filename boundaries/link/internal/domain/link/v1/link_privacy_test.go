//go:build unit

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLink_IsPublic(t *testing.T) {
	tests := []struct {
		name           string
		allowedEmails  []string
		wantIsPublic   bool
		wantDescription string
	}{
		{
			name:           "public link with empty allowlist",
			allowedEmails:  []string{},
			wantIsPublic:   true,
			wantDescription: "empty slice means public",
		},
		{
			name:           "public link with nil allowlist",
			allowedEmails:  nil,
			wantIsPublic:   true,
			wantDescription: "nil means public",
		},
		{
			name:           "private link with single email",
			allowedEmails:  []string{"user@example.com"},
			wantIsPublic:   false,
			wantDescription: "non-empty slice means private",
		},
		{
			name:           "private link with multiple emails",
			allowedEmails:  []string{"user1@example.com", "user2@example.com"},
			wantIsPublic:   false,
			wantDescription: "multiple emails means private",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link, err := NewLinkBuilder().
				SetURL("https://example.com").
				SetAllowedEmails(tt.allowedEmails).
				Build()

			require.NoError(t, err, "Failed to build link")
			assert.Equal(t, tt.wantIsPublic, link.IsPublic(), tt.wantDescription)
		})
	}
}

func TestLink_CanBeViewedByEmail(t *testing.T) {
	tests := []struct {
		name           string
		allowedEmails  []string
		viewerEmail    string
		wantCanView    bool
		wantDescription string
	}{
		// Public links - anyone can view
		{
			name:           "public link - any email can view",
			allowedEmails:  []string{},
			viewerEmail:    "anyone@example.com",
			wantCanView:    true,
			wantDescription: "public links are accessible to everyone",
		},
		{
			name:           "public link - empty email can view",
			allowedEmails:  []string{},
			viewerEmail:    "",
			wantCanView:    true,
			wantDescription: "public links are accessible even without email",
		},
		{
			name:           "public link - nil email can view",
			allowedEmails:  []string{},
			viewerEmail:    "",
			wantCanView:    true,
			wantDescription: "public links are accessible",
		},

		// Private links - only allowlist can view
		{
			name:           "private link - email in allowlist can view",
			allowedEmails:  []string{"user@example.com"},
			viewerEmail:    "user@example.com",
			wantCanView:    true,
			wantDescription: "email in allowlist can view",
		},
		{
			name:           "private link - email not in allowlist cannot view",
			allowedEmails:  []string{"user@example.com"},
			viewerEmail:    "other@example.com",
			wantCanView:    false,
			wantDescription: "email not in allowlist cannot view",
		},
		{
			name:           "private link - empty email cannot view",
			allowedEmails:  []string{"user@example.com"},
			viewerEmail:    "",
			wantCanView:    false,
			wantDescription: "empty email cannot view private link",
		},

		// Email normalization
		{
			name:           "private link - case insensitive match",
			allowedEmails:  []string{"User@Example.com"},
			viewerEmail:    "user@example.com",
			wantCanView:    true,
			wantDescription: "email matching is case insensitive",
		},
		{
			name:           "private link - spaces are normalized",
			allowedEmails:  []string{"  user@example.com  "},
			viewerEmail:    "user@example.com",
			wantCanView:    true,
			wantDescription: "spaces in email are normalized",
		},
		{
			name:           "private link - multiple emails in allowlist",
			allowedEmails:  []string{"user1@example.com", "user2@example.com"},
			viewerEmail:    "user2@example.com",
			wantCanView:    true,
			wantDescription: "any email in allowlist can view",
		},
		{
			name:           "private link - viewer email not in multiple allowlist",
			allowedEmails:  []string{"user1@example.com", "user2@example.com"},
			viewerEmail:    "user3@example.com",
			wantCanView:    false,
			wantDescription: "email must be in allowlist to view",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link, err := NewLinkBuilder().
				SetURL("https://example.com").
				SetAllowedEmails(tt.allowedEmails).
				Build()

			require.NoError(t, err, "Failed to build link")
			assert.Equal(t, tt.wantCanView, link.CanBeViewedByEmail(tt.viewerEmail), tt.wantDescription)
		})
	}
}

func TestLink_GetAllowedEmails(t *testing.T) {
	t.Run("returns copy of allowed emails", func(t *testing.T) {
		allowedEmails := []string{"user1@example.com", "user2@example.com"}
		link, err := NewLinkBuilder().
			SetURL("https://example.com").
			SetAllowedEmails(allowedEmails).
			Build()

		require.NoError(t, err)

		// Get allowed emails
		result := link.GetAllowedEmails()

		// Modify the returned slice
		result[0] = "modified@example.com"

		// Original link should not be affected
		original := link.GetAllowedEmails()
		assert.Equal(t, "user1@example.com", original[0], "Original link should not be modified")
		assert.Equal(t, "modified@example.com", result[0], "Returned slice can be modified")
	})

	t.Run("returns nil for nil link", func(t *testing.T) {
		var link *Link
		assert.Nil(t, link.GetAllowedEmails())
	})
}

