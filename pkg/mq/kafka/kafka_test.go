//go:build unit || (mq && kafka)

package kafka

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq/query"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))

	os.Exit(m.Run())
}

func TestKafka(t *testing.T) {
	viper.SetDefault("SERVICE_NAME", "shortlink")

	ctx, cancel := context.WithCancel(context.Background())
	mq := Kafka{}

	log, err := logger.New(logger.Zap, config.Configuration{})
	require.NoError(t, err, "Cannot create logger")

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// create a network with Client.CreateNetwork()
	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
		Name: "shortlink-test-kafka",
	})
	if err != nil {
		t.Fatalf("Error create docker network: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	ZOOKEEPER, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "confluentinc/cp-zookeeper",
		Tag:          "7.5.3",
		ExposedPorts: []string{"2181"},
		Name:         "test-kafka-zookeeper",
		Env:          []string{"ZOOKEEPER_CLIENT_PORT=2181", "ZOOKEEPER_TICK_TIME=2000"},
		NetworkID:    network.ID,
	})
	require.NoError(t, err, "Could not start resource")

	KAFKA, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "confluentinc/cp-kafka",
		Tag:        "7.5.3",
		Name:       "test-kafka-server",
		Hostname:   "kafka",
		Env: []string{
			"KAFKA_BROKER_ID=1",
			"KAFKA_ZOOKEEPER_CONNECT=test-kafka-zookeeper:2181",
			"KAFKA_ADVERTISED_LISTENERS=INSIDE://kafka:9092,OUTSIDE://localhost:19093",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=INSIDE:PLAINTEXT,OUTSIDE:PLAINTEXT",
			"KAFKA_LISTENERS=INSIDE://0.0.0.0:9092,OUTSIDE://0.0.0.0:19093",
			"KAFKA_INTER_BROKER_LISTENER_NAME=INSIDE",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
		},
		ExposedPorts: []string{"19093/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"19093/tcp": {{HostIP: "localhost", HostPort: "19093/tcp"}},
		},
		NetworkID: network.ID,
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(ZOOKEEPER); errPurge != nil {
			require.Errorf(t, errPurge, "Could not purge resource")
		}

		if err := pool.Client.RemoveNetwork(network.ID); err != nil {
			require.Errorf(t, err, "Could not purge resource")
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("MQ_KAFKA_URI", fmt.Sprintf("localhost:%s", KAFKA.GetPort("19093/tcp")))

		err = mq.Init(ctx, log)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
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

		// When you're done, kill and remove the container
		if err := pool.Purge(ZOOKEEPER); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}

		if err := pool.Purge(KAFKA); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}

		if err := pool.Client.RemoveNetwork(network.ID); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})
}
