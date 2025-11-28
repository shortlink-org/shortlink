package v1

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/rules"
)

var (
	errInvalidURL       = ErrInvalidInput("invalid URL")
	errInvalidCreatedAt = ErrInvalidInput("invalid timestamp: created at is nil")
	errInvalidUpdatedAt = ErrInvalidInput("invalid timestamp: updated at is nil")
)

// LinkBuilder is used to build a new Link
type LinkBuilder struct {
	link  *Link
	build error
}

// NewLinkBuilder returns a new instance of LinkBuilder
func NewLinkBuilder() *LinkBuilder {
	return &LinkBuilder{link: &Link{}}
}

// SetURL sets the URL of the link and calculates the hash
func (b *LinkBuilder) SetURL(newURL string) *LinkBuilder {
	link, err := newUrl(newURL)
	if err != nil {
		b.build = errors.Join(b.build, errInvalidURL)
		return b
	}

	b.link.url = link
	b.link.hash = newHash(link.GetUrl())

	return b
}

// SetDescribe sets the description of the link
func (b *LinkBuilder) SetDescribe(describe string) *LinkBuilder {
	b.link.describe = describe

	return b
}

// SetCreatedAt sets the creation timestamp of the link
func (b *LinkBuilder) SetCreatedAt(createdAt time.Time) *LinkBuilder {
	if createdAt.IsZero() {
		b.build = errors.Join(b.build, errInvalidCreatedAt)
		return b
	}

	b.link.createdAt = Time(createdAt)

	return b
}

// SetUpdatedAt sets the update timestamp of the link
func (b *LinkBuilder) SetUpdatedAt(updatedAt time.Time) *LinkBuilder {
	if updatedAt.IsZero() {
		b.build = errors.Join(b.build, errInvalidUpdatedAt)
		return b
	}

	b.link.updatedAt = Time(updatedAt)

	return b
}

// SetAllowedEmails sets the list of allowed emails for private link access.
// Validates email format, checks for duplicates, and enforces size limit using specification pattern.
func (b *LinkBuilder) SetAllowedEmails(emails []string) *LinkBuilder {
	normalizedEmails, err := rules.ValidateEmailAllowlist(emails)
	if err != nil {
		// Convert standard errors from rules package to LinkError
		errStr := err.Error()
		var linkErr *LinkError

		// Check for specific error patterns and convert to appropriate LinkError
		if strings.Contains(errStr, "allowlist too large") {
			// Parse size from error message: "allowlist too large: X emails (max: Y)"
			var currentSize, maxSize int
			fmt.Sscanf(errStr, "allowlist too large: %d emails (max: %d)", &currentSize, &maxSize)
			linkErr = ErrAllowlistTooLarge(currentSize, maxSize)
		} else if strings.Contains(errStr, "duplicate email") {
			// Extract email from error message: "duplicate email in allowlist: email@example.com"
			email := strings.TrimPrefix(errStr, "duplicate email in allowlist: ")
			linkErr = ErrDuplicateEmail(email)
		} else if strings.Contains(errStr, "invalid email") {
			// Extract email from error message: "invalid email: email@example.com"
			email := strings.TrimPrefix(errStr, "invalid email: ")
			linkErr = ErrInvalidEmail(email)
		} else {
			// Generic error
			linkErr = ErrInvalidInput(errStr)
		}

		b.build = errors.Join(b.build, linkErr)
		return b
	}

	b.link.allowedEmails = normalizedEmails

	return b
}

// Build finalizes the building process and returns the built Link
func (b *LinkBuilder) Build() (*Link, error) {
	if b.build != nil {
		return nil, b.build
	}

	if b.link.createdAt.GetTime().IsZero() {
		b.link.createdAt = Time(time.Now())
	}

	if b.link.updatedAt.GetTime().IsZero() {
		b.link.updatedAt = Time(time.Now())
	}

	return b.link, nil
}
