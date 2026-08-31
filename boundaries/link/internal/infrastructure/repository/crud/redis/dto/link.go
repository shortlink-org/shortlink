package dto

import (
	"time"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// Link is the Data Transfer Object for Link JSON serialization in Redis.
//
// The domain Link keeps every field unexported, so it cannot be handed to
// encoding/json directly: marshalling it yields an empty object and
// unmarshalling into it silently leaves the link blank.
type Link struct {
	URL           string    `json:"url"`
	Hash          string    `json:"hash"`
	Describe      string    `json:"describe"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	AllowedEmails []string  `json:"allowed_emails,omitempty"`
}

// FromDomain converts the domain Link to the DTO Link for JSON serialization.
func FromDomain(d *domain.Link) *Link {
	if d == nil {
		return nil
	}

	return &Link{
		URL:           d.GetUrl().String(),
		Hash:          d.GetHash(),
		Describe:      d.GetDescribe(),
		CreatedAt:     d.GetCreatedAt().GetTime(),
		UpdatedAt:     d.GetUpdatedAt().GetTime(),
		AllowedEmails: d.GetAllowedEmails(),
	}
}

// ToDomain rebuilds the domain Link from its stored representation.
//
// The hash is derived from the URL by the builder, so it is not set here:
// a round-trip reproduces the same hash that was stored.
func (l *Link) ToDomain() (*domain.Link, error) {
	if l == nil {
		return nil, nil //nolint:nilnil // an absent record maps to an absent link
	}

	builder := domain.NewLinkBuilder().
		SetURL(l.URL).
		SetDescribe(l.Describe)

	// The builder rejects zero timestamps, so only set what was actually stored.
	if !l.CreatedAt.IsZero() {
		builder = builder.SetCreatedAt(l.CreatedAt)
	}

	if !l.UpdatedAt.IsZero() {
		builder = builder.SetUpdatedAt(l.UpdatedAt)
	}

	if len(l.AllowedEmails) > 0 {
		builder = builder.SetAllowedEmails(l.AllowedEmails)
	}

	return builder.Build()
}
