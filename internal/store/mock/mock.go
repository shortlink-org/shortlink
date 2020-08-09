package mock

import (
	"github.com/batazor/shortlink/pkg/link"
	"github.com/golang/protobuf/ptypes"
)

var (
	timestamp = ptypes.TimestampNow()
	AddLink   = &link.Link{ // nolint unused
		Url:       "https://example.com",
		Hash:      "",
		Describe:  "example link",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	GetLink = &link.Link{ // nolint unused
		Url:       "https://example.com",
		Hash:      "5888cabde",
		Describe:  "example link",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
)
