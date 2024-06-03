//go:build unit || (database && postgres)

package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/pkg/db/options"
	db "github.com/shortlink-org/shortlink/pkg/db/postgres"
)

func TestMain(m *testing.M) {
	// TODO: research how correct close store
	// pgxpool: https://github.com/jackc/pgx/pull/1642
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"),
		goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("sync.runtime_Semacquire"))

	os.Exit(m.Run())
}

var linkUniqId atomic.Int64

func TestPostgres(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	st := &db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("ghcr.io/dbsystel/postgresql-partman", "16", []string{
		"POSTGRESQL_USERNAME=postgres",
		"POSTGRESQL_PASSWORD=shortlink",
		"POSTGRESQL_DATABASE=link",
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/link?sslmode=disable", resource.GetPort("5432/tcp")))

		errInit := st.Init(ctx)
		if errInit != nil {
			return errInit
		}

		return nil
	}); err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		require.NoError(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	// new store
	store, err := New(ctx, st)
	if err != nil {
		t.Fatalf("Could not create store: %s", err)
	}

	timestamp := time.Now()
	mockLink, err := v1.NewLinkBuilder().
		SetURL("https://example.com").
		SetDescribe("example link").
		SetCreatedAt(timestamp).
		SetUpdatedAt(timestamp).
		Build()

	t.Run("Create [single]", func(t *testing.T) {
		link, err := store.Add(ctx, mockLink)
		require.NoError(t, err)
		assert.Equal(t, mockLink.GetHash(), link.GetHash())
		assert.Equal(t, mockLink.GetDescribe(), link.GetDescribe())
	})

	t.Run("Create [batch]", func(t *testing.T) {
		// Set config
		t.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		newCtx, cancelBatchMode := context.WithCancel(ctx)

		// new store
		storeBatchMode, err := New(newCtx, st)
		if err != nil {
			t.Fatalf("Could not create store: %s", err)
		}

		source, err := getLink()
		require.NoError(t, err)
		_, err = storeBatchMode.Add(ctx, source)
		require.NoError(t, err)
		assert.NotNil(t, source.GetCreatedAt())
		assert.Equal(t, source.GetDescribe(), mockLink.GetDescribe())

		source, err = getLink()
		require.NoError(t, err)
		_, err = storeBatchMode.Add(ctx, source)
		require.NoError(t, err)
		assert.NotNil(t, source.GetCreatedAt())
		assert.Equal(t, source.GetDescribe(), mockLink.GetDescribe())

		source, err = getLink()
		require.NoError(t, err)
		_, err = storeBatchMode.Add(ctx, source)
		require.NoError(t, err)
		assert.NotNil(t, source.GetCreatedAt())
		assert.Equal(t, source.GetDescribe(), mockLink.GetDescribe())

		source, err = getLink()
		require.NoError(t, err)
		_, err = storeBatchMode.Add(ctx, source)
		require.NoError(t, err)
		assert.NotNil(t, source.GetCreatedAt())
		assert.Equal(t, source.GetDescribe(), mockLink.GetDescribe())

		t.Cleanup(func() {
			cancelBatchMode()
		})
	})

	t.Run("Get by hash", func(t *testing.T) {
		link, err := store.Get(ctx, mockLink.GetHash())
		require.NoError(t, err)
		assert.Equal(t, link.GetHash(), mockLink.GetHash())
		assert.Equal(t, link.GetDescribe(), mockLink.GetDescribe())
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, 8, len(links.GetLinks()))
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mockLink.GetHash()))
	})
}

func getLink() (*v1.Link, error) {
	id := linkUniqId.Add(1)

	timestamp := time.Now()
	data, err := v1.NewLinkBuilder().
		SetURL(fmt.Sprintf("%s/%d", "http://example.com", id)).
		SetDescribe("example link").
		SetCreatedAt(timestamp).
		SetUpdatedAt(timestamp).
		Build()

	if err != nil {
		return nil, err
	}

	return data, nil
}
