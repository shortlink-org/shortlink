//go:build unit || (database && postgres)

package tariff_repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	db "github.com/shortlink-org/shortlink/pkg/db/postgres"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/config"
	"github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing"
)

func TestMain(m *testing.M) {
	// TODO: research how correct close store
	// pgxpool: https://github.com/jackc/pgx/pull/1642
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"),
		goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("sync.runtime_Semacquire"))

	os.Exit(m.Run())
}

func TestPostgres(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// Init logger
	conf := config.Configuration{}
	log, err := logger.New(logger.Zap, conf)

	// Init store
	st := &db.Store{}
	require.NoError(t, err, "Error init a logger")

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("ghcr.io/dbsystel/postgresql-partman", "16", []string{
		"POSTGRESQL_USERNAME=postgres",
		"POSTGRESQL_PASSWORD=shortlink",
		"POSTGRESQL_DATABASE=billing",
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
		t.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/billing?sslmode=disable", resource.GetPort("5432/tcp")))

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

	// Init event sourcing
	es, err := eventsourcing.New(ctx, log, st)
	require.NotNil(t, es, "Event sourcing is nil")

	// new store
	store, err := New(ctx, st)
	if err != nil {
		t.Fatalf("Could not create store: %s", err)
	}

	require.NotNil(t, store, "Store is nil")
}
