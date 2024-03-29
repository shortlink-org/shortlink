//go:build unit

package db

import (
	"context"
	"testing"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/config"

	"github.com/stretchr/testify/require"
)

// TestLink ...
func TestLink(t *testing.T) {
	ctx := context.Background()

	// Init logger
	conf := config.Configuration{}
	log, err := logger.New(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Init db
	s, err := New(ctx, log, nil, nil)
	require.NoError(t, err, "Error init a db")

	// Init db
	require.NoError(t, s.Init(ctx), "Error  create a new DB connection")
}
