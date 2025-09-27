package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainErrors(t *testing.T) {
	t.Run("ErrorMessages", func(t *testing.T) {
		tests := []struct {
			err      error
			expected string
		}{
			{ErrRecorderDisabled, "flight recorder is disabled"},
			{ErrAlreadyRunning, "flight recorder is already running"},
			{ErrNotRunning, "flight recorder is not running"},
			{ErrInvalidConfiguration, "invalid flight recorder configuration"},
			{ErrInvalidMinAge, "minimum age must be at least 1 second"},
			{ErrInvalidMaxBytes, "maximum bytes must be at least 1MB"},
			{ErrTraceNotFound, "trace data not found"},
			{ErrRepositoryUnavailable, "trace repository is unavailable"},
		}

		for _, tt := range tests {
			assert.Equal(t, tt.expected, tt.err.Error())
		}
	})

	t.Run("ErrorIdentity", func(t *testing.T) {
		// Test that errors can be compared using errors.Is
		assert.True(t, errors.Is(ErrRecorderDisabled, ErrRecorderDisabled))
		assert.False(t, errors.Is(ErrRecorderDisabled, ErrAlreadyRunning))
	})
}

func TestValidationError(t *testing.T) {
	t.Run("ValidationErrorWithoutCause", func(t *testing.T) {
		err := NewValidationError("field1", "value1", "invalid value", nil)
		
		assert.Equal(t, "field1", err.Field)
		assert.Equal(t, "value1", err.Value)
		assert.Equal(t, "invalid value", err.Message)
		assert.Nil(t, err.Cause)
		assert.Equal(t, "invalid value", err.Error())
	})

	t.Run("ValidationErrorWithCause", func(t *testing.T) {
		cause := errors.New("underlying error")
		err := NewValidationError("field2", 42, "validation failed", cause)
		
		assert.Equal(t, "field2", err.Field)
		assert.Equal(t, 42, err.Value)
		assert.Equal(t, "validation failed", err.Message)
		assert.Equal(t, cause, err.Cause)
		assert.Equal(t, "validation failed: underlying error", err.Error())
	})

	t.Run("ValidationErrorUnwrap", func(t *testing.T) {
		cause := errors.New("root cause")
		err := NewValidationError("field", "value", "message", cause)
		
		unwrapped := err.Unwrap()
		assert.Equal(t, cause, unwrapped)
	})

	t.Run("ValidationErrorUnwrapNoCause", func(t *testing.T) {
		err := NewValidationError("field", "value", "message", nil)
		
		unwrapped := err.Unwrap()
		assert.Nil(t, unwrapped)
	})

	t.Run("ValidationErrorIsSupport", func(t *testing.T) {
		cause := ErrInvalidMinAge
		err := NewValidationError("minAge", "500ms", "too short", cause)
		
		// Should support errors.Is for the wrapped error
		assert.True(t, errors.Is(err, ErrInvalidMinAge))
		assert.False(t, errors.Is(err, ErrInvalidMaxBytes))
	})

	t.Run("ValidationErrorChaining", func(t *testing.T) {
		// Test error chaining with multiple levels
		rootCause := errors.New("root cause")
		middleErr := NewValidationError("middle", "value", "middle error", rootCause)
		topErr := NewValidationError("top", "value", "top error", middleErr)
		
		assert.Equal(t, "top error: middle error: root cause", topErr.Error())
		assert.True(t, errors.Is(topErr, rootCause))
	})
}

func TestValidationErrorTypes(t *testing.T) {
	t.Run("DifferentValueTypes", func(t *testing.T) {
		tests := []struct {
			name  string
			value interface{}
		}{
			{"string", "test"},
			{"int", 42},
			{"float", 3.14},
			{"bool", true},
			{"nil", nil},
			{"slice", []string{"a", "b"}},
			{"map", map[string]int{"key": 1}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := NewValidationError("field", tt.value, "test message", nil)
				assert.Equal(t, tt.value, err.Value)
				assert.Equal(t, "test message", err.Error())
			})
		}
	})
}

func TestErrorComposition(t *testing.T) {
	t.Run("WrapDomainErrors", func(t *testing.T) {
		// Test wrapping domain errors in validation errors
		wrappedErr := NewValidationError(
			"recorder_state", 
			"disabled", 
			"cannot start recorder", 
			ErrRecorderDisabled,
		)
		
		assert.Equal(t, "cannot start recorder: flight recorder is disabled", wrappedErr.Error())
		assert.True(t, errors.Is(wrappedErr, ErrRecorderDisabled))
	})

	t.Run("MultipleWrapping", func(t *testing.T) {
		// Test multiple levels of error wrapping
		baseErr := ErrInvalidMinAge
		validationErr := NewValidationError("min_age", "500ms", "invalid duration", baseErr)
		configErr := NewValidationError("config", nil, "configuration error", validationErr)
		
		expectedMsg := "configuration error: invalid duration: minimum age must be at least 1 second"
		assert.Equal(t, expectedMsg, configErr.Error())
		assert.True(t, errors.Is(configErr, ErrInvalidMinAge))
	})
}

// Benchmark tests for error operations
func BenchmarkValidationErrorCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewValidationError("field", "value", "message", nil)
	}
}

func BenchmarkValidationErrorWithCause(b *testing.B) {
	cause := errors.New("test cause")
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = NewValidationError("field", "value", "message", cause)
	}
}

func BenchmarkValidationErrorError(b *testing.B) {
	err := NewValidationError("field", "value", "message", errors.New("cause"))
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}