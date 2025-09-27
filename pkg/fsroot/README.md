# SafeFS - Secure Filesystem Operations with os.Root

The `fsroot` package provides a secure wrapper around Go 1.25's `os.Root` type to restrict filesystem operations to specific directory boundaries and prevent path traversal attacks.

## Overview

`os.Root` was introduced in Go 1.25 to enhance filesystem security by confining file operations to a specified directory. This is particularly useful when handling untrusted input or when you need to ensure that file operations cannot escape a designated directory tree.

## Features

- **Path Traversal Protection**: Prevents access to files outside the designated root directory
- **Simple API**: Familiar file operation methods with built-in security
- **Safe Defaults**: All paths are automatically cleaned and validated
- **Comprehensive Coverage**: Supports reading, writing, creating, removing files and directories
- **Error Handling**: Clear error messages with context

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/shortlink-org/shortlink/pkg/fsroot"
)

func main() {
    // Create a SafeFS rooted at a specific directory
    fs, err := fsroot.NewSafeFS("/safe/directory")
    if err != nil {
        log.Fatal(err)
    }
    defer fs.Close()

    // Read a file within the root directory
    content, err := fs.ReadFile("config.json")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Content: %s\n", content)
}
```

### Reading Directory Contents

```go
// List files in the root directory
entries, err := fs.ReadDir(".")
if err != nil {
    log.Fatal(err)
}

for _, entry := range entries {
    fmt.Printf("Found: %s (dir: %v)\n", entry.Name(), entry.IsDir())
}

// List files in a subdirectory
subEntries, err := fs.ReadDir("subdirectory")
if err != nil {
    log.Fatal(err)
}
```

### Writing Files

```go
// Write content to a file
data := []byte("Hello, secure world!")
err := fs.WriteFile("output.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}

// Create and write to a file manually
file, err := fs.Create("manual.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("Manual content")
if err != nil {
    log.Fatal(err)
}
```

### File Operations

```go
// Check if file exists and get info
info, err := fs.Stat("somefile.txt")
if err != nil {
    log.Printf("File does not exist: %v", err)
} else {
    fmt.Printf("File size: %d bytes\n", info.Size())
}

// Remove a file
err = fs.Remove("unwanted.txt")
if err != nil {
    log.Printf("Failed to remove file: %v", err)
}
```

## Convenience Function

For simple operations, you can use the `OpenInRoot` convenience function:

```go
// Open a file directly without creating a SafeFS instance
file, err := fsroot.OpenInRoot("/safe/directory", "config.json")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// Read from the file
content, err := io.ReadAll(file)
if err != nil {
    log.Fatal(err)
}
```

## Security Benefits

### Path Traversal Prevention

The SafeFS wrapper prevents path traversal attacks automatically:

```go
fs, _ := fsroot.NewSafeFS("/safe/directory")

// These operations will fail safely:
_, err := fs.ReadFile("../../../etc/passwd")        // Error: outside root
_, err = fs.ReadFile("/etc/passwd")                 // Error: absolute path
_, err = fs.Open("subdir/../../../sensitive.txt")  // Error: traversal attempt
```

### Safe Relative Paths

All paths are automatically cleaned and validated:

```go
// These are equivalent and safe:
content1, _ := fs.ReadFile("./subdir/file.txt")
content2, _ := fs.ReadFile("subdir/file.txt")
content3, _ := fs.ReadFile("subdir//file.txt")
```

## API Reference

### SafeFS Type

```go
type SafeFS struct {
    // Contains filtered or unexported fields
}
```

### Constructor

```go
func NewSafeFS(rootDir string) (*SafeFS, error)
```
Creates a new SafeFS instance rooted at the specified directory.

### Methods

#### File Reading
- `ReadFile(filename string) ([]byte, error)` - Read entire file contents
- `Open(filename string) (*os.File, error)` - Open file for reading

#### Directory Reading
- `ReadDir(dirname string) ([]fs.DirEntry, error)` - List directory contents

#### File Writing
- `WriteFile(filename string, data []byte, perm os.FileMode) error` - Write data to file
- `Create(filename string) (*os.File, error)` - Create/truncate file for writing

#### File Operations
- `Stat(filename string) (os.FileInfo, error)` - Get file information
- `Remove(name string) error` - Remove file or directory

#### Utility
- `RootDir() string` - Get the root directory path
- `Close() error` - Close the SafeFS instance

### Standalone Functions

```go
func OpenInRoot(rootDir, filename string) (*os.File, error)
```
Convenience function to open a file within a root directory in a single operation.

## Error Handling

All SafeFS operations return descriptive errors with context:

```go
_, err := fs.ReadFile("nonexistent.txt")
if err != nil {
    // Error message includes the filename and operation
    fmt.Printf("Error: %v\n", err)
    // Output: failed to read file "nonexistent.txt": open nonexistent.txt: no such file or directory
}
```

## Best Practices

1. **Always Close SafeFS**: Use `defer fs.Close()` to ensure proper cleanup
2. **Use Relative Paths**: Only use relative paths, never absolute paths
3. **Validate Input**: Validate user-provided paths before using them
4. **Error Handling**: Always check for errors from SafeFS operations
5. **Root Selection**: Choose the most restrictive root directory possible

## Integration Examples

### Web Server File Uploads

```go
func handleUpload(w http.ResponseWriter, r *http.Request) {
    // Create SafeFS for uploads directory
    fs, err := fsroot.NewSafeFS("/var/uploads")
    if err != nil {
        http.Error(w, "Server error", 500)
        return
    }
    defer fs.Close()

    // Save uploaded file safely
    filename := r.FormValue("filename")
    data, _ := io.ReadAll(r.Body)
    
    err = fs.WriteFile(filename, data, 0644)
    if err != nil {
        http.Error(w, "Upload failed", 500)
        return
    }
}
```

### Configuration File Processing

```go
func loadConfig(configDir string) (*Config, error) {
    fs, err := fsroot.NewSafeFS(configDir)
    if err != nil {
        return nil, err
    }
    defer fs.Close()

    // Read main config
    mainConfig, err := fs.ReadFile("config.yaml")
    if err != nil {
        return nil, err
    }

    // Read additional configs from subdirectories
    entries, err := fs.ReadDir("modules")
    if err != nil {
        return nil, err
    }

    // Process each module config safely
    for _, entry := range entries {
        if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") {
            moduleConfig, err := fs.ReadFile("modules/" + entry.Name())
            if err != nil {
                continue // Skip invalid configs
            }
            // Process module config...
        }
    }

    return config, nil
}
```

## Testing

The package includes comprehensive tests demonstrating all functionality:

```bash
go test -tags=unit ./pkg/fsroot -v
```

## Migration Guide

If you're currently using direct `os` package functions, here's how to migrate:

### Before (Unsafe)
```go
files, err := os.ReadDir(userProvidedPath)
content, err := os.ReadFile(filepath.Join(userProvidedPath, filename))
```

### After (Safe)
```go
fs, err := fsroot.NewSafeFS(userProvidedPath)
defer fs.Close()

files, err := fs.ReadDir(".")
content, err := fs.ReadFile(filename)
```

## Compatibility

- Requires Go 1.25 or later
- Compatible with all platforms supported by `os.Root`
- Thread-safe for concurrent operations

## Limitations

- Only works within a single directory tree
- Cannot access files outside the root directory
- Symbolic links are followed but cannot point outside the root
- Some platform-specific limitations apply (see `os.Root` documentation)