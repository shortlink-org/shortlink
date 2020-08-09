package mock

import (
	"github.com/golang/protobuf/ptypes"

	"github.com/batazor/shortlink/pkg/domain/link"
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
