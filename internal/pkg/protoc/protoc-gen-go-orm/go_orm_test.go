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
	generatedFileName := "fixtures/github.com/shortlink-org/link.orm.go" // Update this with the correct file name
	content, err := os.ReadFile(generatedFileName)
	require.NoError(t, err, "Failed to read generated file")

	// Check if the content of the generated file is as expected
	expectedContent := "type FilterLink struct {" // Update this with the expected content
	require.Contains(t, string(content), expectedContent, "Generated file does not contain expected content")
}

func TestFilterLink_BuildFilter(t *testing.T) {
	// Create an instance of FilterLink with some criteria
	filter := fixtures.FilterLink{
		Url:      &fixtures.StringFilterInput{Eq: "https://example.com"},
		Describe: &fixtures.StringFilterInput{Contains: "test"},
		// Add more fields as necessary...
	}

	// Create a new Squirrel select builder
	query := squirrel.Select("*").From("links")

	// Build the query using the filter
	query = filter.BuildFilter(query)

	// Convert the query to SQL and arguments
	sql, args, err := query.ToSql()
	require.NoError(t, err, "Failed to build SQL query")

	// Check if the SQL query and arguments are as expected
	expectedSQL := "SELECT * FROM links WHERE url = ? AND description LIKE ?"
	expectedArgs := []interface{}{"https://example.com", "%test%"}
	require.Equal(t, expectedSQL, sql, "SQL query does not match expected")
	require.Equal(t, expectedArgs, args, "Query arguments do not match expected")
}
