package dto

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// ToLinkDeletedEvent creates LinkDeleted event from hash
func ToLinkDeletedEvent(hash string) *linkpb.LinkDeleted {
	return &linkpb.LinkDeleted{
		Hash:       hash,
		OccurredAt: timestamppb.New(time.Now()),
	}
}
