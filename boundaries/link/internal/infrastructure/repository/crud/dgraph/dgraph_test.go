//go:build unit || (database && dgraph)

package dgraph

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dgraph-io/dgo/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/goleak"

	db "github.com/shortlink-org/go-sdk/db/drivers/dgraph"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/mock"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

const (
	dgraphImage     = "dgraph/dgraph:v23.1.0"
	dgraphAlphaPort = "9080/tcp"
)

func TestDgraph(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	st := db.Store{}

	t.Setenv("STORE_DGRAPH_URI", startDgraphCluster(t))
	require.NoError(t, st.Init(ctx))

	store := Store{
		client: st.GetConn().(*dgo.Dgraph),
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}

func startDgraphCluster(tb testing.TB) string {
	tb.Helper()

	ctx := context.Background()

	suffix := time.Now().UnixNano()
	networkName := fmt.Sprintf("shortlink-test-%d", suffix)
	zeroName := fmt.Sprintf("test-dgraph-zero-%d", suffix)
	alphaName := fmt.Sprintf("test-dgraph-alpha-%d", suffix)

	network, err := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
		NetworkRequest: testcontainers.NetworkRequest{
			Name:           networkName,
			CheckDuplicate: true,
		},
	})
	require.NoError(tb, err)

	tb.Cleanup(func() {
		removeCtx, removeCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer removeCancel()

		require.NoError(tb, network.Remove(removeCtx))
	})

	zeroReq := testcontainers.ContainerRequest{
		Name:  zeroName,
		Image: dgraphImage,
		Cmd: []string{
			"dgraph", "zero", fmt.Sprintf("--my=%s:5080", zeroName),
		},
		ExposedPorts: []string{"5080/tcp"},
		Networks:     []string{networkName},
		NetworkAliases: map[string][]string{
			networkName: {zeroName},
		},
		WaitingFor: wait.ForListeningPort("5080/tcp").WithStartupTimeout(2 * time.Minute),
	}

	zero, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: zeroReq,
		Started:          true,
	})
	require.NoError(tb, err)

	tb.Cleanup(func() {
		terminateCtx, terminateCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer terminateCancel()

		require.NoError(tb, zero.Terminate(terminateCtx))
	})

	alphaReq := testcontainers.ContainerRequest{
		Name:  alphaName,
		Image: dgraphImage,
		Cmd: []string{
			"dgraph", "alpha",
			fmt.Sprintf("--my=%s:7080", alphaName),
			"--lru_mb=2048",
			fmt.Sprintf("--zero=%s:5080", zeroName),
		},
		ExposedPorts: []string{dgraphAlphaPort},
		Networks:     []string{networkName},
		NetworkAliases: map[string][]string{
			networkName: {alphaName},
		},
		WaitingFor: wait.ForListeningPort(dgraphAlphaPort).WithStartupTimeout(2 * time.Minute),
	}

	alpha, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: alphaReq,
		Started:          true,
	})
	require.NoError(tb, err)

	tb.Cleanup(func() {
		terminateCtx, terminateCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer terminateCancel()

		require.NoError(tb, alpha.Terminate(terminateCtx))
	})

	host, err := alpha.Host(ctx)
	require.NoError(tb, err)

	port, err := alpha.MappedPort(ctx, dgraphAlphaPort)
	require.NoError(tb, err)

	return fmt.Sprintf("%s:%s", host, port.Port())
}
