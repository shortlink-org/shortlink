//go:build unit || (database && postgres)

package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pgcontainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"go.uber.org/atomic"
	"go.uber.org/goleak"

	db "github.com/shortlink-org/go-sdk/db/drivers/postgres"
	"github.com/shortlink-org/go-sdk/db/options"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
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

const (
	postgresImage = "ghcr.io/dbsystel/postgresql-partman:16"
)

func TestPostgres(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	st := &db.Store{}

	t.Setenv("STORE_POSTGRES_URI", startPostgresContainer(t))
	require.NoError(t, st.Init(ctx))

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

func startPostgresContainer(tb testing.TB) string {
	tb.Helper()

	ctx := context.Background()

	container, err := pgcontainer.Run(
		ctx,
		postgresImage,
		pgcontainer.WithUsername("postgres"),
		pgcontainer.WithPassword("shortlink"),
		pgcontainer.WithDatabase("link"),
	)
	require.NoError(tb, err)

	tb.Cleanup(func() {
		terminateCtx, terminateCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer terminateCancel()

		require.NoError(tb, container.Terminate(terminateCtx))
	})

	return container.MustConnectionString(ctx, "sslmode=disable")
}
