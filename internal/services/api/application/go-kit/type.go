package gokit

import (
	"encoding/json"
	"time"

	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
)

// API ...
type API struct{}

// getRequest ...
type getRequest struct { // nolint unused
	Hash     string
	Describe string
}

// ResponseLink for custom JSON parsing
type ResponseLink struct { // nolint unused
	*v1.Link
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
