package v1

import (
	"errors"
	"net/url"

	"github.com/google/uuid"
)

// RaftBuilder is used to build a new Raft.
type RaftBuilder struct {
	raft   *Raft
	errors error
}

// NewRaftBuilder returns a new instance of RaftBuilder.
func NewRaftBuilder() *RaftBuilder {
	return &RaftBuilder{raft: &Raft{}}
}

// SetPeerIDs sets the peerIDs of the Raft.
func (b *RaftBuilder) SetPeerIDs(peerIDs []string) *RaftBuilder {
	b.raft.peerIDs = peerIDs
	return b
}

// SetName sets the name of the Raft.
func (b *RaftBuilder) SetName(name string) *RaftBuilder {
	b.raft.name = name
	return b
}

// SetAddress sets the address of the Raft.
func (b *RaftBuilder) SetAddress(address string) *RaftBuilder {
	parsedURL, err := url.Parse(address)
	if err != nil {
		b.errors = errors.Join(b.errors, ErrInvalidURL)
		return b
	}

	b.raft.address = *parsedURL

	return b
}

// SetWeight sets the weight of the Raft.
func (b *RaftBuilder) SetWeight(weight int32) *RaftBuilder {
	b.raft.weight = weight
	return b
}

// Build finalizes the building process and returns the built Raft.
func (b *RaftBuilder) Build() (*Raft, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	// set default values
	if b.raft.weight == 0 {
		b.raft.weight = 1
	}

	b.raft.status = RaftStatus_RAFT_STATUS_FOLLOWER

	// generate a new UUID for the Raft
	b.raft.id = uuid.New()

	return b.raft, nil
}
