//go:build unit || (database && postgres)

package postgres

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	pgcontainer "github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/shortlink-org/go-sdk/config"
	db "github.com/shortlink-org/go-sdk/db/drivers/postgres"
)

const (
	postgresImage = "ghcr.io/dbsystel/postgresql-partman:16"
)

func newTestStore(tb testing.TB) (context.Context, *db.Store, *pgxpool.Pool) {
	tb.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	tb.Cleanup(cancel)

	cfg, err := config.New()
	require.NoError(tb, err)

	tb.Setenv("STORE_POSTGRES_URI", startPostgresContainer(tb))

	store := db.New(nil, nil, cfg)
	require.NoError(tb, store.Init(ctx))

	pool := store.GetConn().(*pgxpool.Pool)
	setupSchema(tb, pool)

	return ctx, store, pool
}

func setupSchema(tb testing.TB, pool *pgxpool.Pool) {
	tb.Helper()

	ctx := context.Background()

	statements := []string{
		`CREATE SCHEMA IF NOT EXISTS link`,
		`DROP TABLE IF EXISTS link.link_view`,
		`
		CREATE TABLE link.link_view
		(
			id uuid NOT NULL,
			url text NOT NULL,
			hash varchar(20) NOT NULL,
			describe text,
			image_url text,
			meta_description varchar,
			meta_keywords varchar,
			created_at timestamp,
			updated_at timestamp
		)
		`,
		`
		CREATE OR REPLACE FUNCTION make_tsvector_link_view(description TEXT, keywords TEXT)
		RETURNS tsvector AS $$
		BEGIN
			RETURN (setweight(to_tsvector('simple', keywords),'A') ||
				setweight(to_tsvector('simple', description), 'B'));
		END
		$$ LANGUAGE 'plpgsql' IMMUTABLE
		`,
	}

	for _, stmt := range statements {
		_, err := pool.Exec(ctx, stmt)
		require.NoError(tb, err)
	}
}

func startPostgresContainer(tb testing.TB) string {
	tb.Helper()

	if os.Getenv("SKIP_DOCKER") != "" {
		tb.Skip("skipping testcontainers: SKIP_DOCKER is set")
	}

	defer func() {
		if r := recover(); r != nil {
			tb.Skipf("skipping testcontainers: docker unavailable (%v)", r)
		}
	}()

	ctx := context.Background()

	container, err := pgcontainer.Run(
		ctx,
		postgresImage,
		pgcontainer.WithUsername("postgres"),
		pgcontainer.WithPassword("shortlink"),
		pgcontainer.WithDatabase("link"),
	)
	if err != nil {
		tb.Skipf("skipping testcontainers: could not start postgres container (%v)", err)
	}

	tb.Cleanup(func() {
		terminateCtx, terminateCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer terminateCancel()

		require.NoError(tb, container.Terminate(terminateCtx))
	})

	return container.MustConnectionString(ctx, "sslmode=disable")
}
