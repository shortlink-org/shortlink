package auth

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/require"
)

func TestGetPermissions(t *testing.T) {
	// Create a mock file system
	mockFS := fstest.MapFS{
		"test1.zed": &fstest.MapFile{Data: []byte("content1")},
		"test2.zed": &fstest.MapFile{Data: []byte("content2")},
		"test3.txt": &fstest.MapFile{Data: []byte("content3")},
	}

	permissionsData, err := GetPermissions(mockFS)
	require.NoError(t, err)

	// Expecting 2 files with .zed extension
	require.Len(t, permissionsData, 2)

	// Check the content of the first file
	require.Equal(t, "content1", string(permissionsData[0]))

	// Check the content of the second file
	require.Equal(t, "content2", string(permissionsData[1]))
}
