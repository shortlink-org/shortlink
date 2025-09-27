package fsroot_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/shortlink-org/shortlink/pkg/fsroot"
)

// Example_basicUsage demonstrates basic SafeFS operations
func Example_basicUsage() {
	// Create a temporary directory for the example
	tempDir, err := os.MkdirTemp("", "safefs_example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create some test files
	os.WriteFile(filepath.Join(tempDir, "config.txt"), []byte("config content"), 0644)
	os.Mkdir(filepath.Join(tempDir, "data"), 0755)
	os.WriteFile(filepath.Join(tempDir, "data", "file.txt"), []byte("data content"), 0644)

	// Create a SafeFS rooted at the temporary directory
	fs, err := fsroot.NewSafeFS(tempDir)
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()

	// Read a file
	content, err := fs.ReadFile("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config content: %s\n", content)

	// List directory contents
	entries, err := fs.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Root directory contains %d items:\n", len(entries))
	for _, entry := range entries {
		fmt.Printf("- %s (dir: %v)\n", entry.Name(), entry.IsDir())
	}

	// Read from subdirectory
	subContent, err := fs.ReadFile("data/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data content: %s\n", subContent)

	// Output:
	// Config content: config content
	// Root directory contains 2 items:
	// - config.txt (dir: false)
	// - data (dir: true)
	// Data content: data content
}

// Example_writeOperations demonstrates file writing with SafeFS
func Example_writeOperations() {
	// Create a temporary directory for the example
	tempDir, err := os.MkdirTemp("", "safefs_write_example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a SafeFS rooted at the temporary directory
	fs, err := fsroot.NewSafeFS(tempDir)
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()

	// Write a file using WriteFile
	err = fs.WriteFile("output.txt", []byte("Hello, SafeFS!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Create and write to a file manually
	file, err := fs.Create("manual.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Manually written content")
	if err != nil {
		log.Fatal(err)
	}

	// Read back the files to verify
	content1, _ := fs.ReadFile("output.txt")
	content2, _ := fs.ReadFile("manual.txt")

	fmt.Printf("output.txt: %s\n", content1)
	fmt.Printf("manual.txt: %s\n", content2)

	// Output:
	// output.txt: Hello, SafeFS!
	// manual.txt: Manually written content
}

// Example_securityFeatures demonstrates the security aspects of SafeFS
func Example_securityFeatures() {
	// Create a temporary directory for the example
	tempDir, err := os.MkdirTemp("", "safefs_security_example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test file
	os.WriteFile(filepath.Join(tempDir, "safe.txt"), []byte("safe content"), 0644)

	// Create a SafeFS rooted at the temporary directory
	fs, err := fsroot.NewSafeFS(tempDir)
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()

	// This works - accessing file within root
	content, err := fs.ReadFile("safe.txt")
	if err != nil {
		fmt.Printf("Error reading safe.txt: %v\n", err)
	} else {
		fmt.Printf("Successfully read: %s\n", content)
	}

	// These operations will fail safely due to path traversal attempts
	dangerous_paths := []string{
		"../../../etc/passwd",    // Path traversal
		"/etc/passwd",           // Absolute path
		"../outside.txt",        // Outside root
	}

	for _, path := range dangerous_paths {
		_, err := fs.ReadFile(path)
		if err != nil {
			fmt.Printf("Safely blocked access to: %s\n", path)
		}
	}

	// Output:
	// Successfully read: safe content
	// Safely blocked access to: ../../../etc/passwd
	// Safely blocked access to: /etc/passwd
	// Safely blocked access to: ../outside.txt
}

// Example_openInRoot demonstrates the convenience function
func Example_openInRoot() {
	// Create a temporary directory for the example
	tempDir, err := os.MkdirTemp("", "safefs_openinroot_example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test file
	testContent := []byte("test file content")
	os.WriteFile(filepath.Join(tempDir, "test.txt"), testContent, 0644)

	// Use OpenInRoot for simple file access
	file, err := fsroot.OpenInRoot(tempDir, "test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the content
	buffer := make([]byte, len(testContent))
	n, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])

	// Output:
	// Read 17 bytes: test file content
}