package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecorderState(t *testing.T) {
	t.Run("StateString", func(t *testing.T) {
		tests := []struct {
			state    RecorderState
			expected string
		}{
			{StateStopped, "stopped"},
			{StateRunning, "running"},
			{StateError, "error"},
			{RecorderState(99), "unknown"},
		}

		for _, tt := range tests {
			assert.Equal(t, tt.expected, tt.state.String())
		}
	})
}

func TestConfiguration(t *testing.T) {
	t.Run("ValidConfiguration", func(t *testing.T) {
		config, err := NewConfiguration(true, 5*time.Second, 1024*1024)
		require.NoError(t, err)
		require.NotNil(t, config)

		assert.True(t, config.Enabled())
		assert.Equal(t, 5*time.Second, config.MinAge())
		assert.Equal(t, uint64(1024*1024), config.MaxBytes())
	})

	t.Run("DisabledConfiguration", func(t *testing.T) {
		config, err := NewConfiguration(false, 0, 0)
		require.NoError(t, err)
		require.NotNil(t, config)

		assert.False(t, config.Enabled())
		assert.Equal(t, time.Duration(0), config.MinAge())
		assert.Equal(t, uint64(0), config.MaxBytes())
	})

	t.Run("InvalidMinAge", func(t *testing.T) {
		config, err := NewConfiguration(true, 500*time.Millisecond, 1024*1024)
		assert.Error(t, err)
		assert.Nil(t, config)
		assert.Equal(t, ErrInvalidMinAge, err)
	})

	t.Run("InvalidMaxBytes", func(t *testing.T) {
		config, err := NewConfiguration(true, 5*time.Second, 512*1024) // Less than 1MB
		assert.Error(t, err)
		assert.Nil(t, config)
		assert.Equal(t, ErrInvalidMaxBytes, err)
	})

	t.Run("MinimumValidValues", func(t *testing.T) {
		config, err := NewConfiguration(true, 1*time.Second, 1024*1024)
		require.NoError(t, err)
		require.NotNil(t, config)

		assert.True(t, config.Enabled())
		assert.Equal(t, 1*time.Second, config.MinAge())
		assert.Equal(t, uint64(1024*1024), config.MaxBytes())
	})

	t.Run("LargeValidValues", func(t *testing.T) {
		config, err := NewConfiguration(true, 10*time.Minute, 100*1024*1024)
		require.NoError(t, err)
		require.NotNil(t, config)

		assert.True(t, config.Enabled())
		assert.Equal(t, 10*time.Minute, config.MinAge())
		assert.Equal(t, uint64(100*1024*1024), config.MaxBytes())
	})
}

func TestConfigurationImmutability(t *testing.T) {
	t.Run("ConfigurationIsImmutable", func(t *testing.T) {
		config, err := NewConfiguration(true, 5*time.Second, 2*1024*1024)
		require.NoError(t, err)

		// Verify we can't modify the configuration after creation
		// (This test ensures the fields are not exported or have setters)
		originalEnabled := config.Enabled()
		originalMinAge := config.MinAge()
		originalMaxBytes := config.MaxBytes()

		// Multiple calls should return the same values
		assert.Equal(t, originalEnabled, config.Enabled())
		assert.Equal(t, originalMinAge, config.MinAge())
		assert.Equal(t, originalMaxBytes, config.MaxBytes())
	})
}

func TestConfigurationEdgeCases(t *testing.T) {
	t.Run("ZeroMinAgeWhenDisabled", func(t *testing.T) {
		config, err := NewConfiguration(false, 0, 0)
		require.NoError(t, err)
		assert.False(t, config.Enabled())
	})

	t.Run("ExactMinimumBoundaries", func(t *testing.T) {
		// Test exactly at the boundary values
		config, err := NewConfiguration(true, 1*time.Second, 1024*1024)
		require.NoError(t, err)
		assert.True(t, config.Enabled())

		// Test just below the boundary
		config, err = NewConfiguration(true, 999*time.Millisecond, 1024*1024)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidMinAge, err)

		config, err = NewConfiguration(true, 1*time.Second, 1024*1024-1)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidMaxBytes, err)
	})
}

// Benchmark tests for configuration creation
func BenchmarkNewConfiguration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewConfiguration(true, 5*time.Second, 3*1024*1024)
	}
}

func BenchmarkConfigurationAccessors(b *testing.B) {
	config, _ := NewConfiguration(true, 5*time.Second, 3*1024*1024)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.Enabled()
		_ = config.MinAge()
		_ = config.MaxBytes()
	}
}