package mock

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

var (
	timestamp = timestamppb.Now()
	AddLink   = &domain.Link{
		Url:       "https://example.com",
		Hash:      "5888cabde79b6d7",
		Describe:  "example link",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	GetLink = &domain.Link{
		Url:       "https://example.com",
		Hash:      "5888cabde79b6d7",
		Describe:  "example link",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
)
