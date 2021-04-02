package httpchi

import (
	"context"
	"encoding/json"
	"time"

	"github.com/batazor/shortlink/internal/services/api/domain/link"
)

// API ...
type API struct { // nolint unused
	ctx context.Context
}

// addRequest ...
type addRequest struct { // nolint unused
	URL      string
	Describe string
}

// getRequest ...
type getRequest struct { // nolint unused
	Hash     string
	Describe string
	Filter   string
}

// deleteRequest ...
type deleteRequest struct { // nolint unused
	Hash string
}

// ResponseLink for custom JSON parsing
type ResponseLink struct { // nolint unused
	*link.Link
}

func (l ResponseLink) MarshalJSON() ([]byte, error) {
	var createdAt time.Time
	if l.CreatedAt != nil {
		createdAt = l.CreatedAt.AsTime()
	}

	var updatedAt time.Time
	if l.CreatedAt != nil {
		updatedAt = l.CreatedAt.AsTime()
	}

	return json.Marshal(&struct {
		Url       string
		Hash      string
		Describe  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}{
		Url:       l.Url,
		Hash:      l.Hash,
		Describe:  l.Describe,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
}
