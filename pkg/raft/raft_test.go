package raft

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	v1 "github.com/shortlink-org/shortlink/pkg/raft/domain/raft/v1"
)

func Test_Raft(t *testing.T) {
	// create 3 nodes ---------------------------------------------------------
	nodeUUIDOne := uuid.New()
	nodeUUIDTwo := uuid.New()
	nodeUUIDThree := uuid.New()

	listNode := []uuid.UUID{nodeUUIDOne, nodeUUIDTwo, nodeUUIDThree}

	node1, err := v1.NewRaftBuilder().
		SetID(nodeUUIDOne).
		SetPeerIDs(listNode).
		SetName("node1").
		SetAddress("http://localhost:8080").
		Build()
	require.NoError(t, err)

	node2, err := v1.NewRaftBuilder().
		SetID(uuid.New()).
		SetPeerIDs(listNode).
		SetName("node3").
		SetAddress("http://localhost:8081").
		Build()
	require.NoError(t, err)

	node3, err := v1.NewRaftBuilder().
		SetID(uuid.New()).
		SetPeerIDs(listNode).
		SetName("node3").
		SetAddress("http://localhost:8082").
		Build()
	require.NoError(t, err)

	// check status of nodes -------------------------------------------------
	require.Equal(t, node1.GetStatus(), v1.RaftStatus_RAFT_STATUS_FOLLOWER)
	require.Equal(t, node2.GetStatus(), v1.RaftStatus_RAFT_STATUS_FOLLOWER)
	require.Equal(t, node3.GetStatus(), v1.RaftStatus_RAFT_STATUS_FOLLOWER)
}
