package dto

import (
	"time"

	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToLinkDeletedEvent creates LinkDeleted event from hash
func ToLinkDeletedEvent(hash string) *linkpb.LinkDeleted {
	return &linkpb.LinkDeleted{
		Hash:       hash,
		OccurredAt: timestamppb.New(time.Now()),
	}
}

