//go:build unit
// +build unit

package db

import (
	"context"
	"testing"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/stretchr/testify/require"
)

// TestLink ...
func TestLink(t *testing.T) {
	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	var st Store
	s, err := st.Use(ctx, log)
	require.NoError(t, err, "Error init a db")

	// Init db
	require.NoError(t, s.Store.Init(ctx), "Error  create a new link list")
}
