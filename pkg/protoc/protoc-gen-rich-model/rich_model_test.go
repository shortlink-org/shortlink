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
	cmd := exec.Command("protoc", "--rich-model_out="+outputDir, "--proto_path=.", protoPath)

	// Run the command
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "protoc failed with output:\n%s", output)

	// Read the generated file
	generatedFileName := filepath.Join(outputDir, "link_rich.go") // Update this with the correct file name based on your plugin's output
	content, err := os.ReadFile(generatedFileName)
	require.NoError(t, err, "Failed to read generated file")

	// Check if the content of the generated file is as expected
	expectedContent := "type LinkRich struct {" // Update this with the expected content based on what your plugin should generate
	require.Contains(t, string(content), expectedContent, "Generated file does not contain expected content")
}
