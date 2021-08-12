package mock

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
)

var (
	timestamp = timestamppb.Now()
	AddLink   = &v1.Link{ // nolint unused
		Url:       "https://example.com",
		Hash:      "",
		Describe:  "example link",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	GetLink = &v1.Link{ // nolint unused
		Url:       "https://example.com",
		Hash:      "5888cabde",
		Describe:  "example link",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
)
