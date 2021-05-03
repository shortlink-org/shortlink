package mock

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/batazor/shortlink/internal/services/link/domain/link"
)

var (
	timestamp = timestamppb.Now()
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
