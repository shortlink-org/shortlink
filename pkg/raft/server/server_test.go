package server_test

import (
	"context"
	"sync"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/config"
	"github.com/shortlink-org/shortlink/pkg/raft/server"
	v1 "github.com/shortlink-org/shortlink/pkg/raft/v1"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

func Test_Raft(t *testing.T) {
	ctx := context.Background()

	// Init logger
	conf := config.Configuration{
		Level: config.INFO_LEVEL,
	}
	log, err := logger.New(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Step 1. Create 3 nodes ===================================================
	peers := []string{"127.0.0.1:50051", "127.0.0.1:50052", "127.0.0.1:50053"}

	// node 1 -----------------------------------------------------
	//nolint:mnd,revive // It's okay to have magic numbers here
	viper.Set("GRPC_SERVER_PORT", 50051)
	serverRPC1, err := rpc.InitServer(ctx, log, nil, nil)
	require.NoError(t, err)

	node1, err := server.New(ctx, serverRPC1, peers, server.WithLogger(log))
	require.NoError(t, err)

	// node 2 -----------------------------------------------------
	//nolint:mnd,revive // It's okay to have magic numbers here
	viper.Set("GRPC_SERVER_PORT", 50052)
	serverRPC2, err := rpc.InitServer(ctx, log, nil, nil)
	require.NoError(t, err)

	node2, err := server.New(ctx, serverRPC2, peers, server.WithLogger(log))
	require.NoError(t, err)

	// node 3 -----------------------------------------------------
	//nolint:mnd,revive // It's okay to have magic numbers here
	viper.Set("GRPC_SERVER_PORT", 50053)
	serverRPC3, err := rpc.InitServer(ctx, log, nil, nil)
	require.NoError(t, err)

	node3, err := server.New(ctx, serverRPC3, peers, server.WithLogger(log))
	require.NoError(t, err)

	// Check the status of nodes. All nodes should be in follower status
	require.Equal(t, v1.RaftStatus_RAFT_STATUS_FOLLOWER, node1.GetStatus())
	require.Equal(t, v1.RaftStatus_RAFT_STATUS_FOLLOWER, node2.GetStatus())
	require.Equal(t, v1.RaftStatus_RAFT_STATUS_FOLLOWER, node3.GetStatus())

	// Step 2. We wait for the election process to complete =====================
	wg := sync.WaitGroup{}
	wg.Add(1)

	for {
		if node1.GetStatus() == v1.RaftStatus_RAFT_STATUS_LEADER {
			log.InfoWithContext(ctx, "Node 1 is the leader")

			wg.Done()

			break
		}

		if node2.GetStatus() == v1.RaftStatus_RAFT_STATUS_LEADER {
			log.InfoWithContext(ctx, "Node 2 is the leader")

			wg.Done()

			break
		}

		if node3.GetStatus() == v1.RaftStatus_RAFT_STATUS_LEADER {
			log.InfoWithContext(ctx, "Node 3 is the leader")

			wg.Done()

			break
		}
	}

	wg.Wait()
}
