//go:build unit || (mq && kafka)

package kafka

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
)

func TestRedPanda(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	mq := Kafka{}

	log, err := logger.NewLogger(logger.Zap, config.Configuration{})
	require.NoError(t, err, "Cannot create logger")

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// create a network with Client.CreateNetwork()
	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
		Name: "shortlink-test",
	})
	if err != nil {
		assert.Errorf(t, err, "Error create docker network")
		os.Exit(1)
	}

	RED_PANDA, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "docker.redpanda.com/vectorized/redpanda",
		Tag:          "v23.2.14",
		ExposedPorts: []string{"8081", "8082", "9092", "28082", "29092"},
		Name:         "test-redpanda-server",
		Cmd: []string{
			"redpanda",
			"start",
			"--smp",
			"1",
			"--reserve-memory",
			"0M",
			"--overprovisioned",
			"--node-id",
			"0",
			"--kafka-addr",
			"PLAINTEXT://0.0.0.0:29092,OUTSIDE://0.0.0.0:9092",
			"--advertise-kafka-addr",
			"PLAINTEXT://redpanda:29092,OUTSIDE://localhost:9092",
			"--pandaproxy-addr",
			"PLAINTEXT://0.0.0.0:28082,OUTSIDE://0.0.0.0:8082",
			"--advertise-pandaproxy-addr",
			"PLAINTEXT://redpanda:28082,OUTSIDE://localhost:8082",
		},
		NetworkID: network.ID,
	})
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("MQ_KAFKA_URI", fmt.Sprintf("localhost:%s", RED_PANDA.GetPort("9092/tcp")))
		if err != nil {
			assert.Errorf(t, err, "Cannot set ENV")
			return nil
		}

		err = mq.Init(ctx, log)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		assert.Errorf(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		cancel()

		if err := pool.Purge(RED_PANDA); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})
}
