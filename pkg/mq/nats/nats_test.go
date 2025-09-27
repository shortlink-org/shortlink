//go:build unit || (mq && nats)

package nats

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/pkg/mq/query"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestNATS(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "nats")
	t.Attr("component", "mq")

		t.Attr("type", "unit")
		t.Attr("package", "nats")
		t.Attr("component", "mq")
	
	ctx, cancel := context.WithCancel(context.Background())
	mq := New()

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.Nil(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("nats", "2.10-alpine", nil)
	require.Nil(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("MQ_NATS_URI", fmt.Sprintf("nats://localhost:%s", resource.GetPort("4222/tcp")))

		err = mq.Init(ctx, nil)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		require.Nil(t, err, "Could not connect to docker")
	}

	t.Run("Subscribe", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "nats")
		t.Attr("component", "mq")

		respCh := make(chan query.ResponseMessage)
		msg := query.Response{
			Chan: respCh,
		}

		err := mq.Subscribe(ctx, "test", msg)
		require.Nil(t, err, "Cannot subscribe")

		err = mq.Publish(ctx, "", []byte("test"), []byte("test"))
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
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})
}
