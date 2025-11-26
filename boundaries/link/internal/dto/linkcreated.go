package dto

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// LinkData represents link data for conversion (avoids domain import cycle)
type LinkData struct {
	URL       string
	Hash      string
	Describe  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ToLinkCreatedEvent converts LinkData to LinkCreated event
func ToLinkCreatedEvent(link *LinkData) *linkpb.LinkCreated {
	if link == nil {
		return nil
	}

	return &linkpb.LinkCreated{
		Url:        link.URL,
		Hash:       link.Hash,
		Describe:   link.Describe,
		CreatedAt:  timestamppb.New(link.CreatedAt),
		UpdatedAt:  timestamppb.New(link.UpdatedAt),
		OccurredAt: timestamppb.New(time.Now()),
	}
}
