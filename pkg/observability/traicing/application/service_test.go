package application

import (
	"context"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/observability/traicing/domain"
)

// Mock implementations for testing

type MockRecorder struct {
	mock.Mock
}

func (m *MockRecorder) Start(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockRecorder) Stop(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockRecorder) State() domain.RecorderState {
	args := m.Called()
	return args.Get(0).(domain.RecorderState)
}

func (m *MockRecorder) WriteTo(w io.Writer) (int64, error) {
	args := m.Called(w)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRecorder) Configuration() *domain.Configuration {
	args := m.Called()
	return args.Get(0).(*domain.Configuration)
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(ctx context.Context, id string, data io.Reader) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func (m *MockRepository) Load(ctx context.Context, id string) (io.Reader, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(io.Reader), args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepository) List(ctx context.Context) ([]string, error) {
	args := m.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) OnStarted(ctx context.Context, config *domain.Configuration) {
	m.Called(ctx, config)
}

func (m *MockEventHandler) OnStopped(ctx context.Context) {
	m.Called(ctx)
}

func (m *MockEventHandler) OnError(ctx context.Context, err error) {
	m.Called(ctx, err)
}

func (m *MockEventHandler) OnTraceSaved(ctx context.Context, id string, bytes int64) {
	m.Called(ctx, id, bytes)
}

func TestRecorderService(t *testing.T) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(t, err)

	t.Run("NewRecorderService", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)
		assert.NotNil(t, service)
	})
}

func TestRecorderServiceOperations(t *testing.T) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("StartRecording", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		mockRecorder.On("State").Return(domain.StateStopped)
		mockRecorder.On("Start", ctx).Return(nil)
		mockEventHandler.On("OnStarted", ctx, config).Return()

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		err := service.StartRecording(ctx)
		assert.NoError(t, err)

		mockRecorder.AssertExpectations(t)
		mockEventHandler.AssertExpectations(t)
	})

	t.Run("StartRecordingAlreadyRunning", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		mockRecorder.On("State").Return(domain.StateRunning)

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		err := service.StartRecording(ctx)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrAlreadyRunning, err)

		mockRecorder.AssertExpectations(t)
	})

	t.Run("StartRecordingError", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		startError := assert.AnError
		mockRecorder.On("State").Return(domain.StateStopped)
		mockRecorder.On("Start", ctx).Return(startError)
		mockEventHandler.On("OnError", ctx, startError).Return()

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		err := service.StartRecording(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to start flight recorder")

		mockRecorder.AssertExpectations(t)
		mockEventHandler.AssertExpectations(t)
	})

	t.Run("StopRecording", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		mockRecorder.On("State").Return(domain.StateRunning)
		mockRecorder.On("Stop", ctx).Return(nil)
		mockEventHandler.On("OnStopped", ctx).Return()

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		err := service.StopRecording(ctx)
		assert.NoError(t, err)

		mockRecorder.AssertExpectations(t)
		mockEventHandler.AssertExpectations(t)
	})

	t.Run("StopRecordingNotRunning", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		mockRecorder.On("State").Return(domain.StateStopped)

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		err := service.StopRecording(ctx)
		assert.NoError(t, err) // Should be idempotent

		mockRecorder.AssertExpectations(t)
	})
}

func TestRecorderServiceCaptureTrace(t *testing.T) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("CaptureTraceSuccess", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		testData := "test trace data"
		bytesWritten := int64(len(testData))

		mockRecorder.On("State").Return(domain.StateRunning)
		mockRecorder.On("WriteTo", mock.AnythingOfType("*io.PipeWriter")).Return(bytesWritten, nil).Run(func(args mock.Arguments) {
			writer := args.Get(0).(io.Writer)
			writer.Write([]byte(testData))
		})
		mockRepository.On("Save", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*io.PipeReader")).Return(nil)
		mockEventHandler.On("OnTraceSaved", ctx, mock.AnythingOfType("string"), bytesWritten).Return()

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		traceID, err := service.CaptureTrace(ctx, "test_reason")
		assert.NoError(t, err)
		assert.NotEmpty(t, traceID)
		assert.Contains(t, traceID, "test_reason")

		mockRecorder.AssertExpectations(t)
		mockRepository.AssertExpectations(t)
		mockEventHandler.AssertExpectations(t)
	})

	t.Run("CaptureTraceNotRunning", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		mockRecorder.On("State").Return(domain.StateStopped)

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		traceID, err := service.CaptureTrace(ctx, "test_reason")
		assert.Error(t, err)
		assert.Equal(t, domain.ErrNotRunning, err)
		assert.Empty(t, traceID)

		mockRecorder.AssertExpectations(t)
	})
}

func TestRecorderServiceStatus(t *testing.T) {
	config, err := domain.NewConfiguration(true, 5*time.Second, 2*1024*1024)
	require.NoError(t, err)

	t.Run("GetRecorderStatus", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		mockRecorder.On("State").Return(domain.StateRunning)

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		status := service.GetRecorderStatus()
		assert.NotNil(t, status)
		assert.Equal(t, domain.StateRunning, status.State)
		assert.Equal(t, config, status.Configuration)
		assert.True(t, status.Enabled)
		assert.True(t, status.IsHealthy())

		mockRecorder.AssertExpectations(t)
	})
}

func TestRecorderServiceTraceManagement(t *testing.T) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("ListTraces", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		expectedTraces := []string{"trace_001", "trace_002", "trace_003"}
		mockRepository.On("List", ctx).Return(expectedTraces, nil)

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		traces, err := service.ListTraces(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedTraces, traces)

		mockRepository.AssertExpectations(t)
	})

	t.Run("LoadTrace", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		testData := strings.NewReader("test trace content")
		traceID := "test_trace_123"

		mockRepository.On("Load", ctx, traceID).Return(testData, nil)

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		reader, err := service.LoadTrace(ctx, traceID)
		assert.NoError(t, err)
		assert.Equal(t, testData, reader)

		mockRepository.AssertExpectations(t)
	})

	t.Run("DeleteTrace", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		traceID := "trace_to_delete"
		mockRepository.On("Delete", ctx, traceID).Return(nil)

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		err := service.DeleteTrace(ctx, traceID)
		assert.NoError(t, err)

		mockRepository.AssertExpectations(t)
	})
}

func TestRecorderServiceDisabled(t *testing.T) {
	config, err := domain.NewConfiguration(false, 0, 0)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("StartRecordingDisabled", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		err := service.StartRecording(ctx)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrRecorderDisabled, err)

		// No mock expectations set - should not be called
		mockRecorder.AssertNotCalled(t, "Start")
	})
}

func TestRecorderStatusHealth(t *testing.T) {
	t.Run("HealthyStates", func(t *testing.T) {
		tests := []struct {
			enabled bool
			state   domain.RecorderState
			healthy bool
		}{
			{false, domain.StateStopped, true},  // Disabled is healthy
			{true, domain.StateStopped, true},   // Stopped is healthy
			{true, domain.StateRunning, true},   // Running is healthy
			{true, domain.StateError, false},    // Error is not healthy
		}

		for _, tt := range tests {
			config, err := domain.NewConfiguration(tt.enabled, 1*time.Second, 1024*1024)
			if tt.enabled {
				require.NoError(t, err)
			}

			status := &RecorderStatus{
				State:         tt.state,
				Configuration: config,
				Enabled:       tt.enabled,
			}

			assert.Equal(t, tt.healthy, status.IsHealthy(),
				"State: %s, Enabled: %v should be healthy: %v", tt.state, tt.enabled, tt.healthy)
		}
	})
}

// Integration tests using real implementations

func TestRecorderServiceIntegration(t *testing.T) {
	// Skip integration tests in short mode
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	_ = t.TempDir() // For future use in real integration tests
	
	// Create real implementations
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(t, err)

	// Note: We can't easily test the real GoFlightRecorder without the logger interface issue
	// So we'll use mocks for now and add integration tests later when the logger issue is resolved
	
	t.Run("FullWorkflow", func(t *testing.T) {
		mockRecorder := &MockRecorder{}
		mockRepository := &MockRepository{}
		mockEventHandler := &MockEventHandler{}

		ctx := context.Background()

		// Setup expectations for full workflow
		mockRecorder.On("State").Return(domain.StateStopped).Once()
		mockRecorder.On("Start", ctx).Return(nil).Once()
		mockEventHandler.On("OnStarted", ctx, config).Return().Once()

		mockRecorder.On("State").Return(domain.StateRunning).Once()
		mockRecorder.On("WriteTo", mock.AnythingOfType("*io.PipeWriter")).Return(int64(1024), nil).Run(func(args mock.Arguments) {
			writer := args.Get(0).(io.Writer)
			writer.Write([]byte("test trace data"))
		}).Once()
		mockRepository.On("Save", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*io.PipeReader")).Return(nil).Once()
		mockEventHandler.On("OnTraceSaved", ctx, mock.AnythingOfType("string"), int64(1024)).Return().Once()

		mockRecorder.On("State").Return(domain.StateRunning).Once()
		mockRecorder.On("Stop", ctx).Return(nil).Once()
		mockEventHandler.On("OnStopped", ctx).Return().Once()

		service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

		// Start recording
		err := service.StartRecording(ctx)
		assert.NoError(t, err)

		// Capture trace
		traceID, err := service.CaptureTrace(ctx, "test_workflow")
		assert.NoError(t, err)
		assert.NotEmpty(t, traceID)

		// Stop recording
		err = service.StopRecording(ctx)
		assert.NoError(t, err)

		// Verify all expectations were met
		mockRecorder.AssertExpectations(t)
		mockRepository.AssertExpectations(t)
		mockEventHandler.AssertExpectations(t)
	})
}

// Benchmark tests
func BenchmarkRecorderServiceStartStop(b *testing.B) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(b, err)

	mockRecorder := &MockRecorder{}
	mockRepository := &MockRepository{}
	mockEventHandler := &MockEventHandler{}

	ctx := context.Background()

	// Setup mock expectations
	mockRecorder.On("State").Return(domain.StateStopped)
	mockRecorder.On("Start", ctx).Return(nil)
	mockEventHandler.On("OnStarted", ctx, config).Return()
	mockRecorder.On("Stop", ctx).Return(nil)
	mockEventHandler.On("OnStopped", ctx).Return()

	service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = service.StartRecording(ctx)
		_ = service.StopRecording(ctx)
	}
}

func BenchmarkRecorderServiceStatus(b *testing.B) {
	config, err := domain.NewConfiguration(true, 1*time.Second, 1024*1024)
	require.NoError(b, err)

	mockRecorder := &MockRecorder{}
	mockRepository := &MockRepository{}
	mockEventHandler := &MockEventHandler{}

	mockRecorder.On("State").Return(domain.StateRunning)

	service := NewRecorderService(mockRecorder, mockRepository, mockEventHandler, config)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = service.GetRecorderStatus()
	}
}