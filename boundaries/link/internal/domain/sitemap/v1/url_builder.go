package v1

import (
	"errors"
)

var (
	errLocEmpty        = errors.New("sitemap: loc cannot be empty")
	errLastModEmpty    = errors.New("sitemap: lastmod cannot be empty")
	errChangeFreqEmpty = errors.New("sitemap: changefreq cannot be empty")
	errPriorityInvalid = errors.New("sitemap: priority must be between 0.0 and 1.0")
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
		b.errors = errors.Join(b.errors, errLocEmpty)

		return b
	}

	b.url.Loc = loc

	return b
}

// SetLastMod sets the lastmod field of the Url
func (b *UrlBuilder) SetLastMod(lastmod string) *UrlBuilder {
	if lastmod == "" {
		b.errors = errors.Join(b.errors, errLastModEmpty)

		return b
	}

	b.url.LastMod = lastmod

	return b
}

// SetChangeFreq sets the changefreq field of the Url
func (b *UrlBuilder) SetChangeFreq(changefreq string) *UrlBuilder {
	if changefreq == "" {
		b.errors = errors.Join(b.errors, errChangeFreqEmpty)

		return b
	}

	b.url.ChangeFreq = changefreq

	return b
}

// SetPriority sets the priority field of the Url
func (b *UrlBuilder) SetPriority(priority float32) *UrlBuilder {
	if priority < 0.0 || priority > 1.0 {
		b.errors = errors.Join(b.errors, errPriorityInvalid)

		return b
	}

	b.url.Priority = priority

	return b
}

// Build finalizes the building process and returns the built Url and any error
func (b *UrlBuilder) Build() (*Url, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	return b.url, nil
}
