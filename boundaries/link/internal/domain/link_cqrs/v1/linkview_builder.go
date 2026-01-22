package v1

import (
	"errors"
	"net/url"
	"time"
)

// LinkViewBuilder constructs LinkView instances.
type LinkViewBuilder struct {
	view  *LinkView
	build error
}

// NewLinkViewBuilder returns a new builder for LinkView.
func NewLinkViewBuilder() *LinkViewBuilder {
	return &LinkViewBuilder{view: &LinkView{}}
}

// SetURL sets the URL of the link view.
func (b *LinkViewBuilder) SetURL(raw string) *LinkViewBuilder {
	if raw == "" {
		return b
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		b.build = errors.Join(b.build, err)
		return b
	}

	b.view.url = Url(*parsed)

	return b
}

// SetHash sets the hash of the link view.
func (b *LinkViewBuilder) SetHash(hash string) *LinkViewBuilder {
	b.view.hash = hash

	return b
}

// SetDescribe sets the description of the link view.
func (b *LinkViewBuilder) SetDescribe(describe string) *LinkViewBuilder {
	b.view.describe = describe

	return b
}

// SetImageUrl sets the image URL of the link view.
func (b *LinkViewBuilder) SetImageUrl(imageURL string) *LinkViewBuilder {
	b.view.imageUrl = imageURL

	return b
}

// SetMetaDescription sets the meta description of the link view.
func (b *LinkViewBuilder) SetMetaDescription(description string) *LinkViewBuilder {
	b.view.metaDescription = description

	return b
}

// SetMetaKeywords sets the meta keywords of the link view.
func (b *LinkViewBuilder) SetMetaKeywords(keywords string) *LinkViewBuilder {
	b.view.metaKeywords = keywords

	return b
}

// SetCreatedAt sets the created timestamp of the link view.
func (b *LinkViewBuilder) SetCreatedAt(createdAt time.Time) *LinkViewBuilder {
	if !createdAt.IsZero() {
		b.view.createdAt = Time(createdAt)
	}

	return b
}

// SetUpdatedAt sets the updated timestamp of the link view.
func (b *LinkViewBuilder) SetUpdatedAt(updatedAt time.Time) *LinkViewBuilder {
	if !updatedAt.IsZero() {
		b.view.updatedAt = Time(updatedAt)
	}

	return b
}

// Build finalizes the LinkView.
func (b *LinkViewBuilder) Build() (*LinkView, error) {
	if b.build != nil {
		return nil, b.build
	}

	return b.view, nil
}
