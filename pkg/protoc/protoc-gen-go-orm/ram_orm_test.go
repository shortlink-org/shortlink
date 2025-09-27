//go:build unit

package main_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestRAMORMGeneration(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "protoc-gen-go-orm")
	t.Attr("component", "protoc")

		t.Attr("type", "unit")
		t.Attr("package", "protoc-gen-go-orm")
		t.Attr("component", "protoc")
	
	// Path to the proto file
	protoPath := "fixtures/link.proto"

	// Running protoc with the go-orm plugin and postgres flag
	cmd := exec.Command("protoc",
		"--go-orm_out=./output",
		"--go-orm_opt=orm=ram,pkg=example,filter=Link,common_path=github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1",
		"--proto_path=.",
		protoPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("protoc failed: %s, %v", string(output), err)
	}

	// Check if the output file exists and contains PostgreSQL-specific ORM code
	// You would specify the expected output filename based on your plugin's file naming scheme
	expectedFile := "./output/link.ram.orm.go"
	data, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Examples of PostgreSQL-specific checks you might perform
	expectedContents := []string{
		"\"reflect\"", // Check for PostgreSQL specific library import
		// Add more PostgreSQL-specific code snippets to check for
	}

	for _, content := range expectedContents {
		if !strings.Contains(string(data), content) {
			t.Errorf("Generated file does not contain expected PostgreSQL content: %s", content)
		}
	}
}
