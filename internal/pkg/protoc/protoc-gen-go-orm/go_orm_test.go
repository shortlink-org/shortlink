package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/internal/pkg/protoc/protoc-gen-go-orm/fixtures"
)

// TestGenerateFile tests the generateFile function of the plugin
func TestGenerateProto(t *testing.T) {
	// Define the base directory for fixtures
	baseFixturesDir := "fixtures"

	// Define the path to the proto file
	protoPath := filepath.Join(baseFixturesDir, "link.proto")

	// Specify the output directory for the generated Go file
	outputDir := baseFixturesDir

	// Define the command to run protoc with your plugin
	// The output is directed to the fixtures directory
	cmd := exec.Command("protoc", "--go-orm_out="+outputDir, "--proto_path=.", protoPath)

	// Run the command
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "protoc failed with output:\n%s", output)

	// Read the generated file
	generatedFileName := "fixtures/link.orm.go" // Update this with the correct file name
	content, err := os.ReadFile(generatedFileName)
	require.NoError(t, err, "Failed to read generated file")

	// Check if the content of the generated file is as expected
	expectedContent := "type FilterLink struct {" // Update this with the expected content
	require.Contains(t, string(content), expectedContent, "Generated file does not contain expected content")
}

func TestFilterLink_BuildFilter(t *testing.T) {
	tests := []struct {
		name         string
		filter       fixtures.FilterLink
		expectedSQL  string
		expectedArgs []interface{}
	}{
		{
			name: "Equal and Contains",
			filter: fixtures.FilterLink{
				Url:      &fixtures.StringFilterInput{Eq: "https://example.com"},
				Describe: &fixtures.StringFilterInput{Contains: "test"},
			},
			expectedSQL:  "SELECT * FROM links WHERE url = ? AND describe LIKE '%' || ?", // Updated expected SQL
			expectedArgs: []interface{}{"https://example.com", "test"},                   // Args should match the bound variables
		},
		// Add more test cases for different combinations of filter conditions
		// ...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := squirrel.Select("*").From("links")
			query = tt.filter.BuildFilter(query)
			sql, args, err := query.ToSql()

			require.NoError(t, err, "Failed to build SQL query")
			require.Equal(t, tt.expectedSQL, sql, "SQL query does not match expected")
			require.Equal(t, tt.expectedArgs, args, "Query arguments do not match expected")
		})
	}
}
