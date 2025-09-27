// Package fsroot provides safe filesystem operations using os.Root to restrict
// file access to specific directory boundaries and prevent path traversal attacks.
package fsroot

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// SafeFS wraps os.Root to provide confined filesystem operations
type SafeFS struct {
	root   *os.Root
	rootDir string
}

// NewSafeFS creates a new SafeFS instance rooted at the specified directory.
// All subsequent file operations will be restricted to this directory and its subdirectories.
func NewSafeFS(rootDir string) (*SafeFS, error) {
	// Clean and validate the root directory path
	cleanRoot := filepath.Clean(rootDir)
	
	// Open the root directory using os.OpenRoot
	root, err := os.OpenRoot(cleanRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to open root directory %q: %w", cleanRoot, err)
	}

	return &SafeFS{
		root:   root,
		rootDir: cleanRoot,
	}, nil
}

// Close closes the underlying root directory.
// This should be called when the SafeFS is no longer needed.
func (fs *SafeFS) Close() error {
	if fs.root != nil {
		return fs.root.Close()
	}
	return nil
}

// RootDir returns the root directory path that this SafeFS is confined to.
func (fs *SafeFS) RootDir() string {
	return fs.rootDir
}

// ReadFile reads the named file within the root directory and returns the contents.
// The filename must be relative to the root directory.
func (fs *SafeFS) ReadFile(filename string) ([]byte, error) {
	data, err := fs.root.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", filename, err)
	}
	return data, nil
}

// ReadDir reads the named directory within the root and returns a list of directory entries.
// The dirname must be relative to the root directory. Use "." to read the root directory itself.
func (fs *SafeFS) ReadDir(dirname string) ([]fs.DirEntry, error) {
	// Open the directory and read its contents
	file, err := fs.root.Open(dirname)
	if err != nil {
		return nil, fmt.Errorf("failed to open directory %q: %w", dirname, err)
	}
	defer file.Close()

	entries, err := file.ReadDir(-1)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %q: %w", dirname, err)
	}
	return entries, nil
}

// Create creates or truncates the named file within the root directory.
// The filename must be relative to the root directory.
func (fs *SafeFS) Create(filename string) (*os.File, error) {
	file, err := fs.root.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %q: %w", filename, err)
	}
	return file, nil
}

// Open opens the named file within the root directory for reading.
// The filename must be relative to the root directory.
func (fs *SafeFS) Open(filename string) (*os.File, error) {
	file, err := fs.root.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q: %w", filename, err)
	}
	return file, nil
}

// Remove removes the named file or directory within the root directory.
// The name must be relative to the root directory.
func (fs *SafeFS) Remove(name string) error {
	err := fs.root.Remove(name)
	if err != nil {
		return fmt.Errorf("failed to remove %q: %w", name, err)
	}
	return nil
}

// Stat returns a FileInfo describing the named file within the root directory.
// The filename must be relative to the root directory.
func (fs *SafeFS) Stat(filename string) (os.FileInfo, error) {
	info, err := fs.root.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to stat file %q: %w", filename, err)
	}
	return info, nil
}

// WriteFile writes data to the named file within the root directory, creating it if necessary.
// The filename must be relative to the root directory.
func (fs *SafeFS) WriteFile(filename string, data []byte, perm os.FileMode) error {
	err := fs.root.WriteFile(filename, data, perm)
	if err != nil {
		return fmt.Errorf("failed to write file %q: %w", filename, err)
	}
	return nil
}

// OpenInRoot is a convenience function that combines opening a root directory
// and accessing a file within it in a single operation.
func OpenInRoot(rootDir, filename string) (*os.File, error) {
	cleanRoot := filepath.Clean(rootDir)
	file, err := os.OpenInRoot(cleanRoot, filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q in root %q: %w", filename, cleanRoot, err)
	}
	return file, nil
}