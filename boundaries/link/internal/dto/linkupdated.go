package dto

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// ToLinkUpdatedEvent converts LinkData to LinkUpdated event
func ToLinkUpdatedEvent(link *LinkData) *linkpb.LinkUpdated {
	if link == nil {
		return nil
	}

	return &linkpb.LinkUpdated{
		Url:        link.URL,
		Hash:       link.Hash,
		Describe:   link.Describe,
		CreatedAt:  timestamppb.New(link.CreatedAt),
		UpdatedAt:  timestamppb.New(link.UpdatedAt),
		OccurredAt: timestamppb.New(time.Now()),
	}
}
