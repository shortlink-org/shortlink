package httpchi

import (
	"context"
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/batazor/shortlink/pkg/link"
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
	var err error

	var createdAt time.Time
	if l.CreatedAt != nil {
		createdAt, err = ptypes.Timestamp(l.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	var updatedAt time.Time
	if l.CreatedAt != nil {
		updatedAt, err = ptypes.Timestamp(l.CreatedAt)
		if err != nil {
			return nil, err
		}
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
