package csi_driver

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/kubernetes-csi/csi-test/v4/pkg/sanity"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"

	"github.com/batazor/shortlink/internal/logger"
)

func TestDriver(t *testing.T) {
	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	socket := "/tmp/csi.sock"
	endpoint := "unix://" + socket
	if err := os.Remove(socket); err != nil && !os.IsNotExist(err) {
		t.Fatalf("failed to remove unix domain socket file %s, error: %s", socket, err)
	}

	// Setup the full driver and its environment
	driver := &Driver{
		name:     DefaultDriverName,
		endpoint: endpoint,
		nodeID:   "myNode",
		srv:      nil,
		log:      log,
		readyMu:  sync.Mutex{},
		ready:    false,
	}

	ctx, cancel := context.WithCancel(context.Background())

	var eg errgroup.Group
	eg.Go(func() error {
		return driver.Run(ctx)
	})

	cfg := sanity.NewTestConfig()
	if err := os.RemoveAll(cfg.TargetPath); err != nil {
		t.Fatalf("failed to delete target path %s: %s", cfg.TargetPath, err)
	}
	if err := os.RemoveAll(cfg.StagingPath); err != nil {
		t.Fatalf("failed to delete staging path %s: %s", cfg.StagingPath, err)
	}
	cfg.Address = endpoint

	// Now call the test suite
	sanity.Test(t, cfg)
	cancel()
	if err := eg.Wait(); err != nil {
		t.Errorf("driver run failed: %s", err)
	}
}
