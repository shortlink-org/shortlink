package v1

import (
	"net/url"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRaftBuilder(t *testing.T) {
	testCases := []struct {
		name          string
		id            uuid.UUID
		peerIDs       []string
		nameField     string
		address       string
		weight        int32
		expectedError error
	}{
		{
			name:      "Valid Raft",
			id:        mustNewV7(t),
			peerIDs:   []string{"peer1", "peer2"},
			nameField: "RaftNode1",
			address:   "http://127.0.0.1:8080",
			weight:    1,
		},
		{
			name:          "Invalid Address",
			id:            mustNewV7(t),
			peerIDs:       []string{"peer1", "peer2"},
			nameField:     "RaftNode2",
			address:       "://invalid-url",
			weight:        2,
			expectedError: ErrInvalidURL,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			builder := NewRaftBuilder().
				SetPeerIDs(tc.peerIDs).
				SetName(tc.nameField).
				SetAddress(tc.address).
				SetWeight(tc.weight)

			raft, err := builder.Build()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotNil(t, raft)
				require.Equal(t, tc.id, raft.id)
				require.Equal(t, tc.peerIDs, raft.peerIDs)
				require.Equal(t, tc.nameField, raft.name)
				require.Equal(t, tc.weight, raft.weight)

				parsedURL, errParse := url.Parse(tc.address)
				require.NoError(t, errParse)
				require.Equal(t, *parsedURL, raft.address)
			}
		})
	}
}

func mustNewV7(t *testing.T) uuid.UUID {
	t.Helper()

	id, err := uuid.NewV7()
	require.NoError(t, err)

	return id
}
