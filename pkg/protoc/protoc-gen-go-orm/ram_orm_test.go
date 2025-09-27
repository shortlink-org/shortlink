//go:build unit

package main_test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/shortlink-org/shortlink/pkg/fsroot"
)

func TestRAMORMGeneration(t *testing.T) {
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

	// Create SafeFS rooted at the output directory to safely read generated files
	fs, err := fsroot.NewSafeFS("./output")
	if err != nil {
		t.Fatalf("Failed to create SafeFS for output directory: %v", err)
	}
	defer fs.Close()

	// Check if the output file exists and contains RAM-specific ORM code
	data, err := fs.ReadFile("link.ram.orm.go")
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
