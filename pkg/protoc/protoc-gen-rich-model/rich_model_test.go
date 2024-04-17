//go:build unit

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGenerateRichModel tests the generateRichModel function of the plugin
func TestGenerateRichModel(t *testing.T) {
	// Define the base directory for fixtures
	baseFixturesDir := "fixtures"

	// Define the path to the proto file
	protoPath := filepath.Join(baseFixturesDir, "link.proto")

	// Specify the output directory for the generated Go file
	outputDir := baseFixturesDir

	// Define the command to run protoc with your plugin
	// The output is directed to the fixtures directory
	cmd := exec.Command(
		"protoc",
		"--rich-model_out="+outputDir,
		"--rich-model_opt=filter=Link",
		"--proto_path=.", protoPath,
	)

	// Run the command
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "protoc failed with output:\n%s", output)

	// Read the generated file for Link
	linkGeneratedFileName := filepath.Join(outputDir, "Link.ddd.go")
	linkContent, err := os.ReadFile(linkGeneratedFileName)
	require.NoError(t, err, "Failed to read generated file for Link")

	// Check if the content of the generated file is as expected
	expectedContent := []string{
		"package fixtures",   // Check for correct package name
		"import (",           // Check for import block
		"\"time\"",           // Check for specific imports
		"type link struct {", // Check for correct struct definition
		"url       string",   // Check for correct field definition
	}
	for _, exp := range expectedContent {
		require.Contains(t, string(linkContent), exp, "Generated file does not contain expected content: "+exp)
	}
}
