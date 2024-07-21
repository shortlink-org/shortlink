package server

import (
	"context"
	"math/rand"
	"time"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	api "github.com/shortlink-org/shortlink/pkg/raft/rpc/v1"
	v1 "github.com/shortlink-org/shortlink/pkg/raft/v1"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type Server struct {
	api *api.Server
	//nolint:unused // TODO: use client
	clients []*api.Client

	// TODO: use interface for logger
	logger logger.Logger

	raft *v1.Raft

	electionResetTimer time.Duration
}

func New(ctx context.Context, serverRPC *rpc.Server, peers []string, options ...Option) (*Server, error) {
	const ElectionResetTimer = 150 * time.Millisecond
	const MaxRandomTime = 51

	rpcServer, err := api.NewServer(serverRPC)
	if err != nil {
		return nil, err
	}

	server := &Server{
		api: rpcServer,
	}

	for _, option := range options {
		option(server)
	}

	// create connecting to all peers
	// TODO: need implement client
	// for _, peer := range peers {
	// 	client, err := api.NewClient(peer)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	server.clients = append(server.clients, client)
	// }

	// if electionResetTimer is nil, set default value
	if server.electionResetTimer == 0 {
		//nolint:gosec // it's not a security issue
		randTime := time.Duration(rand.Intn(MaxRandomTime)) * time.Millisecond
		server.electionResetTimer = ElectionResetTimer + randTime
	}

	// create raft node
	server.raft, err = v1.NewRaftBuilder().
		SetName("node1"). // TODO: set name from config
		SetPeerIDs(peers).
		Build()

	if err != nil {
		return nil, err
	}

	// run timer
	go server.runTimer(ctx)

	server.logger.InfoWithContext(ctx, "raft server started", field.Fields{
		"election_reset_timer": server.electionResetTimer,
		"status":               v1.RaftStatus_name[int32(server.raft.GetStatus())],
	})

	return server, nil
}

// runTimer runs the election timer.
func (s *Server) runTimer(ctx context.Context) {
	electionResetTimer := time.NewTimer(s.electionResetTimer)

	if s.logger != nil {
		s.logger.Info("election timer started", field.Fields{
			"election_reset_timer": s.electionResetTimer.Milliseconds(),
		})
	}

	for {
		select {
		case <-electionResetTimer.C:
			// check status
			switch s.raft.GetStatus() {
			case v1.RaftStatus_RAFT_STATUS_UNSPECIFIED:
				panic("raft status is unspecified")
			case v1.RaftStatus_RAFT_STATUS_FOLLOWER:
				// if the node is a follower, change to a candidate
				s.candidatePromotion(ctx)
			case v1.RaftStatus_RAFT_STATUS_CANDIDATE:
			case v1.RaftStatus_RAFT_STATUS_LEADER:
				s.sendHeartbeat(ctx)
			}
		case <-ctx.Done():
			return
		}
	}
}

// sendHeartbeat sends a heartbeat to all peers.
func (s *Server) sendHeartbeat(ctx context.Context) {
	for _, peerID := range s.raft.GetPeerIDs() {
		_, err := s.api.AppendEntries(ctx, &api.AppendEntriesRequest{
			Id: s.raft.GetID().String(),
		})

		if err != nil && s.logger != nil {
			s.logger.ErrorWithContext(ctx, "failed to send heartbeat", field.Fields{
				"peer_id": peerID,
			})
		}
	}
}

// GetStatus returns the status of the raft node.
func (s *Server) GetStatus() v1.RaftStatus {
	return s.raft.GetStatus()
}

// candidatePromotion promotes the node to a candidate.
func (s *Server) candidatePromotion(ctx context.Context) {
	s.raft.SetStatus(v1.RaftStatus_RAFT_STATUS_CANDIDATE)

	if s.logger != nil {
		s.logger.InfoWithContext(ctx, "node promoted to candidate", field.Fields{
			"id": s.raft.GetID(),
		})
	}

	// sent vote requests to all peers
	// for _, peerID := range s.raft.GetPeerIDs() {
	//
	// }
}
