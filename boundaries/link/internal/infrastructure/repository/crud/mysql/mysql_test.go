//go:build unit || (database && mysql)

package mysql

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mock"
	db "github.com/shortlink-org/shortlink/pkg/db/drivers/mysql"
)

func TestMain(m *testing.M) {

	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("github.com/go-sql-driver/mysql.(*mysqlConn).startWatcher.func1"))

	os.Exit(m.Run())
}

func TestMysql(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "mysql")
	t.Attr("component", "link")
	t.Attr("driver", "mysql")

	st := &db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "8", []string{
		"MYSQL_DATABASE=link",
		"MYSQL_USER=shortlink",
		"MYSQL_PASSWORD=shortlink",
		"MYSQL_ROOT_PASSWORD=shortlink",
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
		t.Setenv("STORE_MYSQL_URI", fmt.Sprintf("shortlink:shortlink@(localhost:%s)/link", resource.GetPort("3306/tcp")))

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

	t.Run("Create", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mysql")
		t.Attr("component", "link")
		t.Attr("driver", "mysql")

		link, err := store.Add(ctx, mock.AddLink)
		if err != nil {
			t.Fatalf("Could not add link: %s", err)
		}
		assert.Equal(t, mock.AddLink.Hash, link.Hash)
		assert.Equal(t, mock.AddLink.Describe, link.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mysql")
		t.Attr("component", "link")
		t.Attr("driver", "mysql")

		link, err := store.Get(ctx, mock.GetLink.Hash)
		if err != nil {
			t.Fatalf("Could not get link: %s", err)
		}
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mysql")
		t.Attr("component", "link")
		t.Attr("driver", "mysql")

		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mysql")
		t.Attr("component", "link")
		t.Attr("driver", "mysql")

		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}
