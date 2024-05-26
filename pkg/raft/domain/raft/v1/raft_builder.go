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

// SetID sets the id of the Raft.
func (b *RaftBuilder) SetID(id uuid.UUID) *RaftBuilder {
	b.raft.id = id
	return b
}

// SetPeerIDs sets the peerIDs of the Raft.
func (b *RaftBuilder) SetPeerIDs(peerIDs []uuid.UUID) *RaftBuilder {
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

	return b.raft, nil
}
