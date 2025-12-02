package v1

import (
	"errors"
	"time"

	"github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/rules"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/vo/email"
	vo_time "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/vo/time"
	vo_url "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/vo/url"
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
	link, err := vo_url.NewURL(newURL)
	if err != nil {
		b.build = errors.Join(b.build, errInvalidURL)
		return b
	}

	b.link.url = link
	b.link.hash = newHash(link.GetURL())

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

	b.link.createdAt = vo_time.NewTime(createdAt)

	return b
}

// SetUpdatedAt sets the update timestamp of the link
func (b *LinkBuilder) SetUpdatedAt(updatedAt time.Time) *LinkBuilder {
	if updatedAt.IsZero() {
		b.build = errors.Join(b.build, errInvalidUpdatedAt)
		return b
	}

	b.link.updatedAt = vo_time.NewTime(updatedAt)

	return b
}

// SetAllowedEmails sets the list of allowed emails for private link access.
// Validates email format, checks for duplicates, and enforces size limit using specification pattern.
func (b *LinkBuilder) SetAllowedEmails(emails []string) *LinkBuilder {
	normalizedEmails, err := rules.ValidateEmailAllowlist(emails)
	if err != nil {
		var linkErr *LinkError
		var allowlistErr *email.AllowlistTooLargeError
		var duplicateErr *email.DuplicateEmailError
		var invalidErr *email.InvalidEmailError

		switch {
		case errors.As(err, &allowlistErr):
			linkErr = NewLinkError(CodeInvalidInput, allowlistErr.Error(), nil)
		case errors.As(err, &duplicateErr):
			linkErr = NewLinkError(CodeConflict, duplicateErr.Error(), nil)
		case errors.As(err, &invalidErr):
			linkErr = NewLinkError(CodeInvalidInput, invalidErr.Error(), nil)
		default:
			linkErr = ErrInvalidInput(err.Error())
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
		b.link.createdAt = vo_time.NewTime(time.Now())
	}

	if b.link.updatedAt.GetTime().IsZero() {
		b.link.updatedAt = vo_time.NewTime(time.Now())
	}

	return b.link, nil
}
