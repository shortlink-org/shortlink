package gokit

import (
	"google.golang.org/protobuf/encoding/protojson"

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
	resp, err := protojson.Marshal(l.Link)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
