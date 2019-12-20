package mock

import (
	"github.com/batazor/shortlink/pkg/link"
	"github.com/golang/protobuf/ptypes"
)

var (
	timestamp = ptypes.TimestampNow()
	AddLink   = link.Link{
		Url:       "https://example.com",
		Hash:      "",
		Describe:  "example link",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	GetLink = link.Link{
		Url:       "https://example.com",
		Hash:      "5888cab",
		Describe:  "example link",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
)
