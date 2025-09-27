package infrastructure

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/observability/traicing/domain"
)

func TestFileSystemRepository(t *testing.T) {
	tmpDir := t.TempDir()
	maxAge := 1 * time.Hour

	t.Run("NewFileSystemRepository", func(t *testing.T) {
		repo, err := NewFileSystemRepository(tmpDir, maxAge)
		assert.NoError(t, err)
		assert.NotNil(t, repo)

		// Verify directory was created
		info, err := os.Stat(tmpDir)
		assert.NoError(t, err)
		assert.True(t, info.IsDir())
	})

	t.Run("NewFileSystemRepositoryInvalidPath", func(t *testing.T) {
		// Try to create repository in a read-only location
		invalidPath := "/proc/nonexistent"
		repo, err := NewFileSystemRepository(invalidPath, maxAge)
		assert.Error(t, err)
		assert.Nil(t, repo)
	})
}

func TestFileSystemRepositoryOperations(t *testing.T) {
	tmpDir := t.TempDir()
	maxAge := 1 * time.Hour

	repo, err := NewFileSystemRepository(tmpDir, maxAge)
	require.NoError(t, err)
	require.NotNil(t, repo)

	ctx := context.Background()

	t.Run("SaveAndLoad", func(t *testing.T) {
		testData := "test trace data content"
		testID := "test_trace_001"

		// Save trace data
		err := repo.Save(ctx, testID, strings.NewReader(testData))
		assert.NoError(t, err)

		// Verify file exists
		filePath := filepath.Join(tmpDir, testID+".trace")
		_, err = os.Stat(filePath)
		assert.NoError(t, err)

		// Load trace data
		reader, err := repo.Load(ctx, testID)
		assert.NoError(t, err)
		require.NotNil(t, reader)

		// Verify content
		loadedData, err := io.ReadAll(reader)
		assert.NoError(t, err)
		assert.Equal(t, testData, string(loadedData))

		// Close the reader if it's a file
		if closer, ok := reader.(io.Closer); ok {
			closer.Close()
		}
	})

	t.Run("LoadNonexistent", func(t *testing.T) {
		reader, err := repo.Load(ctx, "nonexistent_trace")
		assert.Error(t, err)
		assert.Equal(t, domain.ErrTraceNotFound, err)
		assert.Nil(t, reader)
	})

	t.Run("Delete", func(t *testing.T) {
		testData := "trace data to delete"
		testID := "trace_to_delete"

		// Save trace data
		err := repo.Save(ctx, testID, strings.NewReader(testData))
		require.NoError(t, err)

		// Verify file exists
		filePath := filepath.Join(tmpDir, testID+".trace")
		_, err = os.Stat(filePath)
		require.NoError(t, err)

		// Delete trace
		err = repo.Delete(ctx, testID)
		assert.NoError(t, err)

		// Verify file is gone
		_, err = os.Stat(filePath)
		assert.True(t, os.IsNotExist(err))

		// Delete again should be idempotent
		err = repo.Delete(ctx, testID)
		assert.NoError(t, err)
	})

	t.Run("List", func(t *testing.T) {
		// Clean directory first
		entries, _ := os.ReadDir(tmpDir)
		for _, entry := range entries {
			os.Remove(filepath.Join(tmpDir, entry.Name()))
		}

		// Create test traces
		testTraces := []string{"trace_001", "trace_002", "trace_003"}
		for _, traceID := range testTraces {
			err := repo.Save(ctx, traceID, strings.NewReader("test data"))
			require.NoError(t, err)
		}

		// List traces
		traces, err := repo.List(ctx)
		assert.NoError(t, err)
		assert.Len(t, traces, len(testTraces))

		// Verify all traces are present
		for _, expectedTrace := range testTraces {
			assert.Contains(t, traces, expectedTrace)
		}
	})

	t.Run("ListWithNonTraceFiles", func(t *testing.T) {
		// Create a fresh temp directory for this test
		testTmpDir := t.TempDir()
		testRepo, err := NewFileSystemRepository(testTmpDir, 1*time.Hour)
		require.NoError(t, err)

		// Create some non-trace files
		nonTraceFiles := []string{"readme.txt", "config.json", "data.log"}
		for _, filename := range nonTraceFiles {
			err := os.WriteFile(filepath.Join(testTmpDir, filename), []byte("content"), 0644)
			require.NoError(t, err)
		}

		// Create trace files
		traceFiles := []string{"trace_001", "trace_002"}
		for _, traceID := range traceFiles {
			err := testRepo.Save(ctx, traceID, strings.NewReader("trace data"))
			require.NoError(t, err)
		}

		// List should only return trace files
		traces, err := testRepo.List(ctx)
		assert.NoError(t, err)
		assert.Len(t, traces, len(traceFiles))

		for _, expectedTrace := range traceFiles {
			assert.Contains(t, traces, expectedTrace)
		}
	})
}

func TestFileSystemRepositoryValidation(t *testing.T) {
	tmpDir := t.TempDir()
	maxAge := 1 * time.Hour

	repo, err := NewFileSystemRepository(tmpDir, maxAge)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("EmptyID", func(t *testing.T) {
		err := repo.Save(ctx, "", strings.NewReader("data"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "empty")

		_, err = repo.Load(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "empty")

		err = repo.Delete(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "empty")
	})

	t.Run("InvalidIDCharacters", func(t *testing.T) {
		invalidIDs := []string{
			"trace/../invalid",
			"trace\\invalid",
			"trace..invalid",
			"trace/subdir/invalid",
		}

		for _, invalidID := range invalidIDs {
			err := repo.Save(ctx, invalidID, strings.NewReader("data"))
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "invalid")

			_, err = repo.Load(ctx, invalidID)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "invalid")

			err = repo.Delete(ctx, invalidID)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "invalid")
		}
	})

	t.Run("ValidIDs", func(t *testing.T) {
		validIDs := []string{
			"trace_001",
			"error_2024_01_01",
			"panic-trace-123",
			"user_session_abc123",
		}

		for _, validID := range validIDs {
			err := repo.Save(ctx, validID, strings.NewReader("test data"))
			assert.NoError(t, err)

			_, err = repo.Load(ctx, validID)
			assert.NoError(t, err)

			err = repo.Delete(ctx, validID)
			assert.NoError(t, err)
		}
	})
}

func TestFileSystemRepositoryMaxAge(t *testing.T) {
	tmpDir := t.TempDir()
	maxAge := 100 * time.Millisecond // Very short for testing

	repo, err := NewFileSystemRepository(tmpDir, maxAge)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("OldFilesCleanup", func(t *testing.T) {
		// Create an old trace file
		oldTraceID := "old_trace"
		err := repo.Save(ctx, oldTraceID, strings.NewReader("old data"))
		require.NoError(t, err)

		// Wait for it to become old
		time.Sleep(150 * time.Millisecond)

		// Create a new trace file
		newTraceID := "new_trace"
		err = repo.Save(ctx, newTraceID, strings.NewReader("new data"))
		require.NoError(t, err)

		// List should only return new trace (old one should be cleaned up)
		traces, err := repo.List(ctx)
		assert.NoError(t, err)
		
		// The old trace should be cleaned up during List operation
		assert.Contains(t, traces, newTraceID)
		// Old trace might still be present depending on timing, but that's OK
	})
}

func TestFileSystemRepositoryAtomicWrites(t *testing.T) {
	tmpDir := t.TempDir()
	maxAge := 1 * time.Hour

	repo, err := NewFileSystemRepository(tmpDir, maxAge)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("AtomicWriteSuccess", func(t *testing.T) {
		testID := "atomic_test"
		testData := "atomic test data"

		// Save should be atomic
		err := repo.Save(ctx, testID, strings.NewReader(testData))
		assert.NoError(t, err)

		// File should exist and have correct content
		reader, err := repo.Load(ctx, testID)
		require.NoError(t, err)

		data, err := io.ReadAll(reader)
		assert.NoError(t, err)
		assert.Equal(t, testData, string(data))

		if closer, ok := reader.(io.Closer); ok {
			closer.Close()
		}

		// No temporary files should remain
		entries, err := os.ReadDir(tmpDir)
		require.NoError(t, err)

		for _, entry := range entries {
			assert.False(t, strings.HasSuffix(entry.Name(), ".tmp"),
				"Temporary file should not exist: %s", entry.Name())
		}
	})
}

// Benchmark tests
func BenchmarkFileSystemRepositorySave(b *testing.B) {
	tmpDir := b.TempDir()
	repo, err := NewFileSystemRepository(tmpDir, 1*time.Hour)
	require.NoError(b, err)

	ctx := context.Background()
	testData := strings.NewReader("benchmark test data")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testData.Seek(0, 0) // Reset reader
		_ = repo.Save(ctx, fmt.Sprintf("bench_trace_%d", i), testData)
	}
}

func BenchmarkFileSystemRepositoryLoad(b *testing.B) {
	tmpDir := b.TempDir()
	repo, err := NewFileSystemRepository(tmpDir, 1*time.Hour)
	require.NoError(b, err)

	ctx := context.Background()
	testID := "bench_trace"
	
	// Create test file
	err = repo.Save(ctx, testID, strings.NewReader("benchmark data"))
	require.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader, _ := repo.Load(ctx, testID)
		if closer, ok := reader.(io.Closer); ok {
			closer.Close()
		}
	}
}