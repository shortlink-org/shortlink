package csi_driver

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/kubernetes-csi/csi-test/v4/pkg/sanity"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"

	"github.com/batazor/shortlink/internal/logger"
)

func TestDriver(t *testing.T) {
	// Create a new context
	ctx := context.Background()

	// TODO: add test
	t.SkipNow()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	socket := "/tmp/csi.sock"
	endpoint := "unix://" + socket
	if err := os.Remove(socket); err != nil && !os.IsNotExist(err) {
		t.Fatalf("failed to remove unix domain socket file %s, error: %s", socket, err)
	}

	// Run CSI Driver
	driver, err := NewDriver(log, DefaultDriverName, "testNode", endpoint, 0)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialize driver: %s", err.Error()))
	}

	ctx, cancel := context.WithCancel(ctx)

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
