package traicing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Clean up global state before each test
	SetGlobalFlightRecorder(nil)

	t.Run("NoGlobalRecorder", func(t *testing.T) {
		status := HealthCheck()
		
		assert.Equal(t, "not_initialized", status["flight_recorder"])
		assert.Equal(t, "disabled", status["status"])
	})

	t.Run("HealthCheckReturnStructure", func(t *testing.T) {
		status := HealthCheck()
		
		// Verify the structure of the returned map
		assert.Contains(t, status, "flight_recorder")
		assert.Contains(t, status, "status")
		
		// Verify the types
		assert.IsType(t, "", status["flight_recorder"])
		assert.IsType(t, "", status["status"])
	})

	t.Run("HealthCheckConsistency", func(t *testing.T) {
		// Multiple calls should return consistent results
		status1 := HealthCheck()
		status2 := HealthCheck()
		
		assert.Equal(t, status1, status2)
	})
}

func TestGlobalFlightRecorderManagement(t *testing.T) {
	// Clean up global state
	SetGlobalFlightRecorder(nil)

	t.Run("InitialStateIsNil", func(t *testing.T) {
		global := GetGlobalFlightRecorder()
		assert.Nil(t, global)
	})

	t.Run("SetNilFlightRecorder", func(t *testing.T) {
		// Setting nil should work
		SetGlobalFlightRecorder(nil)
		global := GetGlobalFlightRecorder()
		assert.Nil(t, global)
	})

	t.Run("GlobalStateIsolation", func(t *testing.T) {
		// Ensure each test starts with clean state
		global1 := GetGlobalFlightRecorder()
		SetGlobalFlightRecorder(nil)
		global2 := GetGlobalFlightRecorder()
		
		assert.Nil(t, global1)
		assert.Nil(t, global2)
	})
}

func TestFlightRecorderConfig(t *testing.T) {
	t.Run("DefaultConfigValues", func(t *testing.T) {
		config := DefaultFlightRecorderConfig()
		
		assert.True(t, config.Enabled)
		assert.Equal(t, 1*time.Minute, config.MinAge)
		assert.Equal(t, int64(3<<20), config.MaxBytes)
	})

	t.Run("ConfigFieldTypes", func(t *testing.T) {
		config := DefaultFlightRecorderConfig()
		
		// Verify field types
		assert.IsType(t, true, config.Enabled)
		assert.IsType(t, 1*time.Minute, config.MinAge)
		assert.IsType(t, int64(0), config.MaxBytes)
	})

	t.Run("ConfigImmutability", func(t *testing.T) {
		config := DefaultFlightRecorderConfig()
		original := config
		
		// Modify the config
		config.Enabled = false
		config.MinAge = 2 * time.Minute
		config.MaxBytes = 1 << 20
		
		// Original should be unchanged (Go structs are passed by value)
		newConfig := DefaultFlightRecorderConfig()
		assert.Equal(t, original.Enabled, newConfig.Enabled)
		assert.Equal(t, original.MinAge, newConfig.MinAge)
		assert.Equal(t, original.MaxBytes, newConfig.MaxBytes)
	})
}

func TestUtilityFunctionsSafety(t *testing.T) {
	// These tests verify that utility functions don't panic when called
	// with no global recorder or invalid parameters

	t.Run("SaveTraceOnErrorSafety", func(t *testing.T) {
		SetGlobalFlightRecorder(nil)
		
		// Should not panic with nil error
		assert.NotPanics(t, func() {
			SaveTraceOnError(nil, nil)
		})
		
		// Should not panic with nil logger
		assert.NotPanics(t, func() {
			SaveTraceOnError(assert.AnError, nil)
		})
	})

	t.Run("SaveTraceOnSignalSafety", func(t *testing.T) {
		SetGlobalFlightRecorder(nil)
		
		// Should not panic with empty signal
		assert.NotPanics(t, func() {
			SaveTraceOnSignal("", nil)
		})
		
		// Should not panic with nil logger
		assert.NotPanics(t, func() {
			SaveTraceOnSignal("USR1", nil)
		})
	})

	t.Run("SaveTraceWithContextSafety", func(t *testing.T) {
		SetGlobalFlightRecorder(nil)
		
		// Should not panic with nil context
		assert.NotPanics(t, func() {
			SaveTraceWithContext(nil, "test", nil, nil)
		})
		
		// Should not panic with nil metadata
		assert.NotPanics(t, func() {
			SaveTraceWithContext(nil, "test", nil, nil)
		})
	})
}

func TestMiddlewareCreation(t *testing.T) {
	t.Run("NewRecorderMiddlewareWithNilLogger", func(t *testing.T) {
		// Should create middleware even with nil logger
		middleware := NewRecorderMiddleware(nil)
		assert.NotNil(t, middleware)
		assert.Nil(t, middleware.log)
	})

	t.Run("MiddlewareWrapFunctions", func(t *testing.T) {
		middleware := NewRecorderMiddleware(nil)
		
		// Test that wrap functions don't panic during creation
		testFunc := func() error { return nil }
		panicFunc := func() { panic("test") }
		
		wrappedErrorFunc := middleware.WrapWithErrorTracking(testFunc)
		assert.NotNil(t, wrappedErrorFunc)
		
		wrappedPanicFunc := middleware.WrapWithPanicRecovery(panicFunc)
		assert.NotNil(t, wrappedPanicFunc)
	})
}

// Benchmark tests for performance verification
func BenchmarkHealthCheck(b *testing.B) {
	SetGlobalFlightRecorder(nil)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = HealthCheck()
	}
}

func BenchmarkGlobalFlightRecorderAccess(b *testing.B) {
	SetGlobalFlightRecorder(nil)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetGlobalFlightRecorder()
	}
}

func BenchmarkDefaultFlightRecorderConfig(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DefaultFlightRecorderConfig()
	}
}

func BenchmarkNewRecorderMiddleware(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewRecorderMiddleware(nil)
	}
}