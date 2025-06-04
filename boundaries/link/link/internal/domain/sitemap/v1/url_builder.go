package v1

import (
	"errors"
)

// UrlBuilder is used to build a new Url
type UrlBuilder struct {
	url    *Url
	errors error
}

// NewUrlBuilder returns a new instance of UrlBuilder
func NewUrlBuilder() *UrlBuilder {
	return &UrlBuilder{url: &Url{}}
}

// SetLoc sets the loc field of the Url
func (b *UrlBuilder) SetLoc(loc string) *UrlBuilder {
	if loc == "" {
		b.errors = errors.Join(b.errors, errors.New("loc cannot be empty"))

		return b
	}

	b.url.loc = loc

	return b
}

// SetLastMod sets the lastmod field of the Url
func (b *UrlBuilder) SetLastMod(lastmod string) *UrlBuilder {
	if lastmod == "" {
		b.errors = errors.Join(b.errors, errors.New("lastmod cannot be empty"))

		return b
	}

	b.url.lastmod = lastmod

	return b
}

// SetChangeFreq sets the changefreq field of the Url
func (b *UrlBuilder) SetChangeFreq(changefreq string) *UrlBuilder {
	if changefreq == "" {
		b.errors = errors.Join(b.errors, errors.New("changefreq cannot be empty"))

		return b
	}

	b.url.changefreq = changefreq

	return b
}

// SetPriority sets the priority field of the Url
func (b *UrlBuilder) SetPriority(priority float32) *UrlBuilder {
	if priority < 0.0 || priority > 1.0 {
		b.errors = errors.Join(b.errors, errors.New("priority must be between 0.0 and 1.0"))

		return b
	}

	b.url.priority = priority

	return b
}

// Build finalizes the building process and returns the built Url and any error
func (b *UrlBuilder) Build() (*Url, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	return b.url, nil
}
