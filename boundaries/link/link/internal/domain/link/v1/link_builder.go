package v1

import (
	"errors"
	"net/url"
	"time"
)

// LinkBuilder is used to build a new Link
type LinkBuilder struct {
	link   *Link
	errors error
}

// NewLinkBuilder returns a new instance of LinkBuilder
func NewLinkBuilder() *LinkBuilder {
	return &LinkBuilder{link: &Link{}}
}

// SetURL sets the URL of the link and calculates the hash
func (b *LinkBuilder) SetURL(newURL string) *LinkBuilder {
	url, err := url.Parse(newURL)
	if err != nil {
		b.errors = errors.Join(b.errors, errors.New("invalid URL"))
		return b
	}

	b.link.url = *url
	b.link.hash = newHash(*url)

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
		b.errors = errors.Join(b.errors, errors.New("invalid timestamp: created at is nil"))
		return b
	}

	b.link.createdAt = createdAt

	return b
}

// SetUpdatedAt sets the update timestamp of the link
func (b *LinkBuilder) SetUpdatedAt(updatedAt time.Time) *LinkBuilder {
	if updatedAt.IsZero() {
		b.errors = errors.Join(b.errors, errors.New("invalid timestamp: updated at is nil"))
		return b
	}

	b.link.updatedAt = updatedAt

	return b
}

// Build finalizes the building process and returns the built Link
func (b *LinkBuilder) Build() (*Link, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	if b.link.createdAt.IsZero() {
		b.link.createdAt = time.Now()
	}

	if b.link.updatedAt.IsZero() {
		b.link.updatedAt = time.Now()
	}

	return b.link, nil
}
