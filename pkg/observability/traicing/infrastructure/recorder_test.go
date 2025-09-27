package infrastructure

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/observability/traicing/domain"
)

func TestGoFlightRecorder(t *testing.T) {
	t.Run("NewGoFlightRecorderEnabled", func(t *testing.T) {
		config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
		require.NoError(t, err)

		recorder, err := NewGoFlightRecorder(config)
		assert.NoError(t, err)
		assert.NotNil(t, recorder)
		assert.Equal(t, domain.StateStopped, recorder.State())
		assert.Equal(t, config, recorder.Configuration())
	})

	t.Run("NewGoFlightRecorderDisabled", func(t *testing.T) {
		config, err := domain.NewConfiguration(false, 0, 0)
		require.NoError(t, err)

		recorder, err := NewGoFlightRecorder(config)
		assert.NoError(t, err)
		assert.NotNil(t, recorder)
		assert.Equal(t, domain.StateStopped, recorder.State())
		assert.Equal(t, config, recorder.Configuration())
	})

	t.Run("NewGoFlightRecorderNilConfig", func(t *testing.T) {
		recorder, err := NewGoFlightRecorder(nil)
		assert.Error(t, err)
		assert.Nil(t, recorder)
		assert.Equal(t, domain.ErrInvalidConfiguration, err)
	})
}

func TestGoFlightRecorderOperations(t *testing.T) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(t, err)

	recorder, err := NewGoFlightRecorder(config)
	require.NoError(t, err)
	require.NotNil(t, recorder)

	ctx := context.Background()

	t.Run("StartStop", func(t *testing.T) {
		// Initial state should be stopped
		assert.Equal(t, domain.StateStopped, recorder.State())

		// Start recording
		err := recorder.Start(ctx)
		assert.NoError(t, err)
		assert.Equal(t, domain.StateRunning, recorder.State())

		// Start again should return error
		err = recorder.Start(ctx)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrAlreadyRunning, err)
		assert.Equal(t, domain.StateRunning, recorder.State())

		// Stop recording
		err = recorder.Stop(ctx)
		assert.NoError(t, err)
		assert.Equal(t, domain.StateStopped, recorder.State())

		// Stop again should be idempotent
		err = recorder.Stop(ctx)
		assert.NoError(t, err)
		assert.Equal(t, domain.StateStopped, recorder.State())
	})

	t.Run("WriteTo", func(t *testing.T) {
		// Start recording
		err := recorder.Start(ctx)
		require.NoError(t, err)
		defer recorder.Stop(ctx)

		// Give it some time to collect trace data
		time.Sleep(50 * time.Millisecond)

		// Write to buffer
		var buf bytes.Buffer
		n, err := recorder.WriteTo(&buf)
		assert.NoError(t, err)
		assert.Greater(t, n, int64(0))
		assert.Greater(t, buf.Len(), 0)
	})

	t.Run("WriteToNotRunning", func(t *testing.T) {
		// Ensure recorder is stopped
		recorder.Stop(ctx)
		assert.Equal(t, domain.StateStopped, recorder.State())

		var buf bytes.Buffer
		n, err := recorder.WriteTo(&buf)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrNotRunning, err)
		assert.Equal(t, int64(0), n)
	})
}

func TestGoFlightRecorderDisabled(t *testing.T) {
	config, err := domain.NewConfiguration(false, 0, 0)
	require.NoError(t, err)

	recorder, err := NewGoFlightRecorder(config)
	require.NoError(t, err)
	require.NotNil(t, recorder)

	ctx := context.Background()

	t.Run("StartDisabled", func(t *testing.T) {
		err := recorder.Start(ctx)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrRecorderDisabled, err)
		assert.Equal(t, domain.StateStopped, recorder.State())
	})

	t.Run("WriteToDisabled", func(t *testing.T) {
		var buf bytes.Buffer
		n, err := recorder.WriteTo(&buf)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrRecorderDisabled, err)
		assert.Equal(t, int64(0), n)
	})

	t.Run("StopDisabled", func(t *testing.T) {
		// Stop should be idempotent even when disabled
		err := recorder.Stop(ctx)
		assert.NoError(t, err)
		assert.Equal(t, domain.StateStopped, recorder.State())
	})
}

func TestGoFlightRecorderConcurrency(t *testing.T) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(t, err)

	recorder, err := NewGoFlightRecorder(config)
	require.NoError(t, err)
	require.NotNil(t, recorder)

	ctx := context.Background()

	t.Run("ConcurrentStartStop", func(t *testing.T) {
		// Test concurrent access to start/stop operations
		done := make(chan bool, 10)

		// Start multiple goroutines trying to start/stop
		for i := 0; i < 5; i++ {
			go func() {
				defer func() { done <- true }()
				_ = recorder.Start(ctx)
				time.Sleep(10 * time.Millisecond)
				_ = recorder.Stop(ctx)
			}()
		}

		// Wait for all goroutines to complete
		for i := 0; i < 5; i++ {
			<-done
		}

		// Final state should be stopped
		assert.Equal(t, domain.StateStopped, recorder.State())
	})

	t.Run("ConcurrentStateAccess", func(t *testing.T) {
		// Test concurrent access to state reading
		done := make(chan bool, 20)

		for i := 0; i < 10; i++ {
			go func() {
				defer func() { done <- true }()
				for j := 0; j < 100; j++ {
					_ = recorder.State()
					_ = recorder.Configuration()
				}
			}()
		}

		for i := 0; i < 10; i++ {
			go func() {
				defer func() { done <- true }()
				_ = recorder.Start(ctx)
				time.Sleep(time.Millisecond)
				_ = recorder.Stop(ctx)
			}()
		}

		// Wait for all goroutines to complete
		for i := 0; i < 20; i++ {
			<-done
		}
	})
}

func TestGoFlightRecorderConfiguration(t *testing.T) {
	t.Run("ConfigurationImmutability", func(t *testing.T) {
		config, err := domain.NewConfiguration(true, 5*time.Second, 3*1024*1024)
		require.NoError(t, err)

		recorder, err := NewGoFlightRecorder(config)
		require.NoError(t, err)

		// Configuration should be the same object reference
		retrievedConfig := recorder.Configuration()
		assert.Equal(t, config, retrievedConfig)

		// Multiple calls should return the same configuration
		config2 := recorder.Configuration()
		assert.Equal(t, config, config2)
	})
}

// Benchmark tests
func BenchmarkGoFlightRecorderStart(b *testing.B) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(b, err)

	recorder, err := NewGoFlightRecorder(config)
	require.NoError(b, err)

	ctx := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = recorder.Start(ctx)
		_ = recorder.Stop(ctx)
	}
}

func BenchmarkGoFlightRecorderState(b *testing.B) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(b, err)

	recorder, err := NewGoFlightRecorder(config)
	require.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = recorder.State()
	}
}

func BenchmarkGoFlightRecorderWriteTo(b *testing.B) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(b, err)

	recorder, err := NewGoFlightRecorder(config)
	require.NoError(b, err)

	ctx := context.Background()
	err = recorder.Start(ctx)
	require.NoError(b, err)
	defer recorder.Stop(ctx)

	// Let it collect some data
	time.Sleep(50 * time.Millisecond)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_, _ = recorder.WriteTo(&buf)
	}
}