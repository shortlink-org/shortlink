//go:build unit || (mq && kafka)

package kafka

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/pkg/mq/query"
)

func TestRedPanda(t *testing.T) {
	// Set configuration
	viper.SetDefault("SERVICE_NAME", "shortlink")
	t.Setenv("MQ_KAFKA_SARAMA_VERSION", "DEFAULT")

	ctx, cancel := context.WithCancel(context.Background())
	mq := Kafka{}

	log, err := logger.New(logger.Zap, config.Configuration{})
	require.NoError(t, err, "Cannot create logger")

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// create a network with Client.CreateNetwork()
	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
		Name: "shortlink-test-redpanda",
	})
	if err != nil {
		t.Fatalf("Error create docker network: %s", err)
	}

	RED_PANDA, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "docker.redpanda.com/vectorized/redpanda",
		Tag:        "v23.2.14",
		Name:       "test-redpanda-server",
		Cmd: []string{
			"redpanda",
			"start",
			"--smp",
			"1",
			"--overprovisioned",
			"--reserve-memory",
			"0M",
			"--node-id",
			"0",
			"--kafka-addr",
			"internal://0.0.0.0:9092,external://0.0.0.0:19092",
			"--advertise-kafka-addr",
			"internal://redpanda:9092,external://localhost:19092",
			"--pandaproxy-addr",
			"internal://0.0.0.0:8082,external://0.0.0.0:18082",
			"--advertise-pandaproxy-addr",
			"internal://redpanda:8082,external://localhost:18082",
		},
		ExposedPorts: []string{"19092/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"19092/tcp": {{HostIP: "localhost", HostPort: "19092/tcp"}},
		},
		NetworkID: network.ID,
	})
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("MQ_KAFKA_URI", fmt.Sprintf("localhost:%s", RED_PANDA.GetPort("19092/tcp")))

		err = mq.Init(ctx, log)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(RED_PANDA); errPurge != nil {
			require.Errorf(t, errPurge, "Could not purge resource")
		}

		if err := pool.Client.RemoveNetwork(network.ID); err != nil {
			require.Errorf(t, err, "Could not purge resource")
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	t.Run("Subscribe", func(t *testing.T) {
		respCh := make(chan query.ResponseMessage)
		msg := query.Response{
			Chan: respCh,
		}

		err := mq.Subscribe(ctx, "test", msg)
		require.Nil(t, err, "Cannot subscribe")

		err = mq.Publish(ctx, "test", []byte("test"), []byte("test"))
		require.Nil(t, err, "Cannot publish")

		select {
		case <-ctx.Done():
			t.Fatal("Timeout")
		case resp := <-respCh:
			require.Equal(t, []byte("test"), resp.Body, "Payloads are not equal")
		}

		err = mq.UnSubscribe("test")
		require.Nil(t, err, "Cannot unsubscribe")
	})

	t.Cleanup(func() {
		cancel()

		if err := pool.Purge(RED_PANDA); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}

		if err := pool.Client.RemoveNetwork(network.ID); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})
}
