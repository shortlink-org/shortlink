package v1

import (
	"net/url"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Raft is the representation of a raft node.
type Raft struct {
	// id is the unique identifier of the raft node.
	id uuid.UUID
	// peerIDs is the list of peer IDs.
	peerIDs []uuid.UUID
	// name is the human-readable name of the raft node.
	name string
	// address is the address of the raft node.
	address url.URL
	// status is the status of the raft node.
	status RaftStatus
	// lastHeartbeat is the last time the raft node sent a heartbeat.
	lastHeartbeat timestamppb.Timestamp
	// weight is the voting weight of the raft node.
	weight int32
}

// GetStatus returns the status of the raft node.
func (r *Raft) GetStatus() RaftStatus {
	return r.status
}
