package dto

import (
	"time"

	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToLinkUpdatedEvent converts LinkData to LinkUpdated event
func ToLinkUpdatedEvent(link LinkData) *linkpb.LinkUpdated {
	return &linkpb.LinkUpdated{
		Url:        link.URL,
		Hash:       link.Hash,
		Describe:   link.Describe,
		CreatedAt:  timestamppb.New(link.CreatedAt),
		UpdatedAt:  timestamppb.New(link.UpdatedAt),
		OccurredAt: timestamppb.New(time.Now()),
	}
}

