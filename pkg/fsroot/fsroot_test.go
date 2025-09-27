//go:build unit

package fsroot

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSafeFS(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "safefs_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create test files and directories
	testFile := filepath.Join(tempDir, "test.txt")
	testContent := []byte("Hello, SafeFS!")
	err = os.WriteFile(testFile, testContent, 0644)
	require.NoError(t, err)

	subDir := filepath.Join(tempDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	require.NoError(t, err)

	subFile := filepath.Join(subDir, "subfile.txt")
	subContent := []byte("Subdirectory content")
	err = os.WriteFile(subFile, subContent, 0644)
	require.NoError(t, err)

	t.Run("NewSafeFS", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		require.NotNil(t, fs)
		require.Equal(t, tempDir, fs.RootDir())
		defer fs.Close()
	})

	t.Run("NewSafeFS_InvalidDir", func(t *testing.T) {
		_, err := NewSafeFS("/nonexistent/directory")
		require.Error(t, err)
	})

	t.Run("ReadFile", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		data, err := fs.ReadFile("test.txt")
		require.NoError(t, err)
		require.Equal(t, testContent, data)
	})

	t.Run("ReadFile_Subdirectory", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		data, err := fs.ReadFile("subdir/subfile.txt")
		require.NoError(t, err)
		require.Equal(t, subContent, data)
	})

	t.Run("ReadFile_NonExistent", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		_, err = fs.ReadFile("nonexistent.txt")
		require.Error(t, err)
	})

	t.Run("ReadDir", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		entries, err := fs.ReadDir(".")
		require.NoError(t, err)
		require.Len(t, entries, 2) // test.txt and subdir

		// Verify entries
		var foundFile, foundDir bool
		for _, entry := range entries {
			switch entry.Name() {
			case "test.txt":
				foundFile = true
				require.False(t, entry.IsDir())
			case "subdir":
				foundDir = true
				require.True(t, entry.IsDir())
			}
		}
		require.True(t, foundFile, "test.txt not found")
		require.True(t, foundDir, "subdir not found")
	})

	t.Run("ReadDir_Subdirectory", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		entries, err := fs.ReadDir("subdir")
		require.NoError(t, err)
		require.Len(t, entries, 1)
		require.Equal(t, "subfile.txt", entries[0].Name())
		require.False(t, entries[0].IsDir())
	})

	t.Run("Create_And_WriteFile", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		// Test Create
		file, err := fs.Create("new_file.txt")
		require.NoError(t, err)
		require.NotNil(t, file)
		file.Close()

		// Test WriteFile
		newContent := []byte("New file content")
		err = fs.WriteFile("written_file.txt", newContent, 0644)
		require.NoError(t, err)

		// Verify both files exist and have correct content
		data, err := fs.ReadFile("written_file.txt")
		require.NoError(t, err)
		require.Equal(t, newContent, data)
	})

	t.Run("Open", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		file, err := fs.Open("test.txt")
		require.NoError(t, err)
		require.NotNil(t, file)
		defer file.Close()

		// Read a few bytes to verify it works
		buffer := make([]byte, 5)
		n, err := file.Read(buffer)
		require.NoError(t, err)
		require.Equal(t, 5, n)
		require.Equal(t, []byte("Hello"), buffer)
	})

	t.Run("Stat", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		info, err := fs.Stat("test.txt")
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, "test.txt", info.Name())
		require.False(t, info.IsDir())
		require.Equal(t, int64(len(testContent)), info.Size())
	})

	t.Run("Remove", func(t *testing.T) {
		fs, err := NewSafeFS(tempDir)
		require.NoError(t, err)
		defer fs.Close()

		// Create a file to remove
		testRemoveFile := "to_remove.txt"
		err = fs.WriteFile(testRemoveFile, []byte("remove me"), 0644)
		require.NoError(t, err)

		// Verify file exists
		_, err = fs.Stat(testRemoveFile)
		require.NoError(t, err)

		// Remove the file
		err = fs.Remove(testRemoveFile)
		require.NoError(t, err)

		// Verify file no longer exists
		_, err = fs.Stat(testRemoveFile)
		require.Error(t, err)
	})
}

func TestOpenInRoot(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "openinroot_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a test file
	testFile := filepath.Join(tempDir, "test.txt")
	testContent := []byte("OpenInRoot test")
	err = os.WriteFile(testFile, testContent, 0644)
	require.NoError(t, err)

	t.Run("OpenInRoot_Success", func(t *testing.T) {
		file, err := OpenInRoot(tempDir, "test.txt")
		require.NoError(t, err)
		require.NotNil(t, file)
		defer file.Close()

		// Read content to verify
		buffer := make([]byte, len(testContent))
		n, err := file.Read(buffer)
		require.NoError(t, err)
		require.Equal(t, len(testContent), n)
		require.Equal(t, testContent, buffer)
	})

	t.Run("OpenInRoot_NonExistent", func(t *testing.T) {
		_, err := OpenInRoot(tempDir, "nonexistent.txt")
		require.Error(t, err)
	})

	t.Run("OpenInRoot_InvalidRoot", func(t *testing.T) {
		_, err := OpenInRoot("/nonexistent/directory", "test.txt")
		require.Error(t, err)
	})
}