// Package infrastructure provides storage implementations for trace data.
// This layer abstracts different storage backends (filesystem, cloud storage, etc.).
package infrastructure

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shortlink-org/shortlink/pkg/observability/traicing/domain"
)

// FileSystemRepository implements domain.Repository using the local filesystem.
// This implementation provides reliable local storage with proper error handling.
type FileSystemRepository struct {
	basePath string
	maxAge   time.Duration
}

// NewFileSystemRepository creates a new filesystem-based repository.
// It ensures the base directory exists and is writable.
func NewFileSystemRepository(basePath string, maxAge time.Duration) (*FileSystemRepository, error) {
	// Ensure base path exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create base directory %s: %w", basePath, err)
	}

	// Verify write permissions
	testFile := filepath.Join(basePath, ".write_test")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return nil, fmt.Errorf("directory %s is not writable: %w", basePath, err)
	}
	os.Remove(testFile) // Clean up test file

	return &FileSystemRepository{
		basePath: basePath,
		maxAge:   maxAge,
	}, nil
}

// Save persists trace data to the filesystem with the given identifier.
// The file is written atomically to prevent corruption.
func (r *FileSystemRepository) Save(ctx context.Context, id string, data io.Reader) error {
	if err := r.validateID(id); err != nil {
		return err
	}

	filePath := r.getFilePath(id)
	tempPath := filePath + ".tmp"

	// Create temporary file for atomic write
	file, err := os.Create(tempPath)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}

	// Ensure cleanup on error
	defer func() {
		file.Close()
		if err != nil {
			os.Remove(tempPath)
		}
	}()

	// Copy data to temporary file
	if _, err = io.Copy(file, data); err != nil {
		return fmt.Errorf("failed to write trace data: %w", err)
	}

	// Ensure data is flushed to disk
	if err = file.Sync(); err != nil {
		return fmt.Errorf("failed to sync trace data: %w", err)
	}

	file.Close()

	// Atomic rename to final location
	if err = os.Rename(tempPath, filePath); err != nil {
		return fmt.Errorf("failed to finalize trace file: %w", err)
	}

	return nil
}

// Load retrieves trace data by identifier.
// Returns domain.ErrTraceNotFound if the trace doesn't exist.
func (r *FileSystemRepository) Load(ctx context.Context, id string) (io.Reader, error) {
	if err := r.validateID(id); err != nil {
		return nil, err
	}

	filePath := r.getFilePath(id)
	
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.ErrTraceNotFound
		}
		return nil, fmt.Errorf("failed to open trace file: %w", err)
	}

	return file, nil
}

// Delete removes trace data by identifier.
// This operation is idempotent - no error is returned if the file doesn't exist.
func (r *FileSystemRepository) Delete(ctx context.Context, id string) error {
	if err := r.validateID(id); err != nil {
		return err
	}

	filePath := r.getFilePath(id)
	
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete trace file: %w", err)
	}

	return nil
}

// List returns available trace identifiers.
// It automatically filters out expired traces based on maxAge.
func (r *FileSystemRepository) List(ctx context.Context) ([]string, error) {
	entries, err := os.ReadDir(r.basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read trace directory: %w", err)
	}

	var traces []string
	cutoff := time.Now().Add(-r.maxAge)

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".trace") {
			continue
		}

		// Check file age if maxAge is configured
		if r.maxAge > 0 {
			info, err := entry.Info()
			if err != nil {
				continue // Skip files we can't stat
			}
			if info.ModTime().Before(cutoff) {
				// Optionally clean up old files
				os.Remove(filepath.Join(r.basePath, entry.Name()))
				continue
			}
		}

		// Extract ID from filename
		id := strings.TrimSuffix(entry.Name(), ".trace")
		traces = append(traces, id)
	}

	return traces, nil
}

// validateID ensures the trace identifier is safe for filesystem use.
func (r *FileSystemRepository) validateID(id string) error {
	if id == "" {
		return fmt.Errorf("trace ID cannot be empty")
	}

	// Prevent directory traversal attacks
	if strings.Contains(id, "..") || strings.Contains(id, "/") || strings.Contains(id, "\\") {
		return fmt.Errorf("invalid trace ID: %s", id)
	}

	return nil
}

// getFilePath returns the full file path for a trace identifier.
func (r *FileSystemRepository) getFilePath(id string) string {
	return filepath.Join(r.basePath, id+".trace")
}

// Ensure FileSystemRepository implements domain.Repository
var _ domain.Repository = (*FileSystemRepository)(nil)