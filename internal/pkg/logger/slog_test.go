package logger

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
)

func TestNewStructLogger(t *testing.T) {
	cfg := config.Configuration{}

	// Test with Zap logger
	zapLogger, err := New(Zap, cfg)
	require.NoError(t, err, "Expected no error creating Zap logger")

	structZapLogger, err := NewStructLogger(zapLogger)
	require.NoError(t, err, "Expected no error creating StructLogger with Zap logger")
	require.NotNil(t, structZapLogger, "Expected StructLogger instance for Zap logger")

	// Test with Logrus logger
	logrusLogger, err := New(Logrus, cfg)
	require.NoError(t, err, "Expected no error creating Logrus logger")

	structLogrusLogger, err := NewStructLogger(logrusLogger)
	require.NoError(t, err, "Expected no error creating StructLogger with Logrus logger")
	require.NotNil(t, structLogrusLogger, "Expected StructLogger instance for Logrus logger")
}

func TestEnabled(t *testing.T) {
	sl := &StructLogger{}
	require.True(t, sl.Enabled(context.Background(), slog.LevelDebug), "Expected Enabled to return true for LevelDebug")
}
