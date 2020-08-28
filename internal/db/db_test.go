package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/logger"
)

// TestLink ...
func TestLink(t *testing.T) { //nolint unused
	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	var st Store
	s, err := st.Use(ctx, log)
	assert.Nil(t, err, "Error init a db")

	// Init db
	assert.Nil(t, s.Init(ctx), "Error  create a new link list")
}
