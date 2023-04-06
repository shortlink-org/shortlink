//go:build unit || (mq && kafka)
// +build unit mq,kafka

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
)

func TestKafka(t *testing.T) {
	mq := Kafka{}
	ctx := context.Background()

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

	// pulls an image, creates a container based on it and runs it
	ZOOKEEPER, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "confluentinc/cp-zookeeper",
		Tag:          "7.0.0",
		ExposedPorts: []string{"2181"},
		Name:         "test-kafka-zookeeper",
		Env:          []string{"ZOOKEEPER_CLIENT_PORT=2181", "ZOOKEEPER_TICK_TIME=2000"},
		NetworkID:    network.ID,
	})
	require.NoError(t, err, "Could not start resource")

	KAFKA, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "confluentinc/cp-kafka",
		Tag:          "7.0.0",
		ExposedPorts: []string{"9092"},
		Name:         "test-kafka-server",
		Env: []string{
			"KAFKA_BROKER_ID=1",
			"KAFKA_ZOOKEEPER_CONNECT=test-kafka-zookeeper:2181",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT",
			"KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
			"KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://test-kafka-server:9092",
		},
		NetworkID: network.ID,
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(ZOOKEEPER); errPurge != nil {
			assert.Errorf(t, errPurge, "Could not purge resource")
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("MQ_KAFKA_URI", fmt.Sprintf("localhost:%s", KAFKA.GetPort("9092/tcp")))
		if err != nil {
			assert.Errorf(t, err, "Cannot set ENV")
			return nil
		}

		err = mq.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		assert.Errorf(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(ZOOKEEPER); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}

		if err := pool.Purge(KAFKA); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	t.Run("Close", func(t *testing.T) {
		require.NoError(t, mq.Close())
	})
}
