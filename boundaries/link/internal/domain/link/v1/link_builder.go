package v1

import (
	"strings"
	"time"
)

var (
	errInvalidURL       = ErrInvalidInput("invalid URL")
	errInvalidCreatedAt = ErrInvalidInput("invalid timestamp: created at is nil")
	errInvalidUpdatedAt = ErrInvalidInput("invalid timestamp: updated at is nil")
)

// LinkBuilder is used to build a new Link
type LinkBuilder struct {
	link *Link
	errs []*LinkError
}

// NewLinkBuilder returns a new instance of LinkBuilder
func NewLinkBuilder() *LinkBuilder {
	return &LinkBuilder{link: &Link{}}
}

func (b *LinkBuilder) appendError(err *LinkError) {
	if err == nil {
		return
	}

	b.errs = append(b.errs, err)
}

// SetURL sets the URL of the link and calculates the hash
func (b *LinkBuilder) SetURL(newURL string) *LinkBuilder {
	link, err := newUrl(newURL)
	if err != nil {
		b.appendError(errInvalidURL)
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
		b.appendError(errInvalidCreatedAt)
		return b
	}

	b.link.createdAt = Time(createdAt)

	return b
}

// SetUpdatedAt sets the update timestamp of the link
func (b *LinkBuilder) SetUpdatedAt(updatedAt time.Time) *LinkBuilder {
	if updatedAt.IsZero() {
		b.appendError(errInvalidUpdatedAt)
		return b
	}

	b.link.updatedAt = Time(updatedAt)

	return b
}

// Build finalizes the building process and returns the built Link
func (b *LinkBuilder) Build() (*Link, error) {
	if len(b.errs) > 0 {
		if len(b.errs) == 1 {
			return nil, b.errs[0]
		}

		details := make([]string, 0, len(b.errs))
		for _, err := range b.errs {
			if err == nil {
				continue
			}
			details = append(details, err.Error())
		}

		return nil, ErrInvalidInput(strings.Join(details, "; "))
	}

	if b.link.createdAt.GetTime().IsZero() {
		b.link.createdAt = Time(time.Now())
	}

	if b.link.updatedAt.GetTime().IsZero() {
		b.link.updatedAt = Time(time.Now())
	}

	return b.link, nil
}
