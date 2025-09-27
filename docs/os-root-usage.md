# Using os.Root for Restricted Filesystem Operations

This document explains how to use Go 1.25's `os.Root` type for secure filesystem operations that prevent path traversal vulnerabilities.

## Overview

`os.Root` was introduced in Go 1.25 to provide directory-limited filesystem access. It restricts all file operations to a specific directory tree, preventing unauthorized access to files outside the designated root directory.

## Basic Usage

### Opening a Root Directory

```go
package main

import (
    "log"
    "os"
)

func main() {
    // Open a root directory - all operations will be restricted to this directory and its subdirectories
    root, err := os.OpenRoot("/safe/directory")
    if err != nil {
        log.Fatal(err)
    }
    defer root.Close() // Always close the root when done
}
```

### Reading Files

```go
// Read a file within the root directory
content, err := root.ReadFile("config.txt")
if err != nil {
    log.Fatal(err)
}

// Read a file in a subdirectory
subContent, err := root.ReadFile("subdir/file.txt")
if err != nil {
    log.Fatal(err)
}
```

### Reading Directories

```go
// Open the root directory itself
dir, err := root.Open(".")
if err != nil {
    log.Fatal(err)
}
defer dir.Close()

// Read directory contents
entries, err := dir.ReadDir(-1) // -1 means read all entries
if err != nil {
    log.Fatal(err)
}

for _, entry := range entries {
    fmt.Printf("Found: %s (dir: %v)\n", entry.Name(), entry.IsDir())
}
```

### Writing Files

```go
// Write data to a file
data := []byte("Hello, secure world!")
err := root.WriteFile("output.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}

// Create a file for manual writing
file, err := root.Create("manual.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("Manual content")
if err != nil {
    log.Fatal(err)
}
```

### Other Operations

```go
// Get file information
info, err := root.Stat("somefile.txt")
if err != nil {
    log.Printf("File does not exist: %v", err)
} else {
    fmt.Printf("File size: %d bytes\n", info.Size())
}

// Remove a file
err = root.Remove("unwanted.txt")
if err != nil {
    log.Printf("Failed to remove file: %v", err)
}

// Create directories
err = root.Mkdir("newdir", 0755)
if err != nil {
    log.Printf("Failed to create directory: %v", err)
}
```

## Security Benefits

### Path Traversal Prevention

`os.Root` automatically prevents path traversal attacks:

```go
root, _ := os.OpenRoot("/safe/directory")
defer root.Close()

// These operations will fail safely:
_, err := root.ReadFile("../../../etc/passwd")        // Error: outside root
_, err = root.ReadFile("/etc/passwd")                 // Error: absolute path
_, err = root.Open("subdir/../../../sensitive.txt")  // Error: traversal attempt
```

### Safe Relative Paths

All paths are automatically cleaned and validated:

```go
// These are equivalent and safe:
content1, _ := root.ReadFile("./subdir/file.txt")
content2, _ := root.ReadFile("subdir/file.txt")
content3, _ := root.ReadFile("subdir//file.txt")
```

## Implementation Examples from the Codebase

### CEL Rules Loader (`poc/cel/rules.go`)

```go
func loadRules(path string) (map[string]string, error) {
    // Use os.OpenRoot to restrict file access to the specified directory
    root, err := os.OpenRoot(path)
    if err != nil {
        return nil, err
    }
    defer root.Close()

    // Open the directory and read its contents
    dir, err := root.Open(".")
    if err != nil {
        return nil, err
    }
    defer dir.Close()

    files, err := dir.ReadDir(-1)
    if err != nil {
        return nil, err
    }

    rules := make(map[string]string)
    for _, file := range files {
        if !file.IsDir() {
            // Read file content using os.Root
            content, errReadFile := root.ReadFile(file.Name())
            if errReadFile != nil {
                return nil, errReadFile
            }
            rules[file.Name()] = string(content)
        }
    }

    return rules, nil
}
```

### Test File Reading (Protoc Tests)

```go
func TestGenerateModel(t *testing.T) {
    // ... generate files ...

    // Use os.OpenRoot to restrict file access to the output directory
    root, err := os.OpenRoot(outputDir)
    require.NoError(t, err, "Failed to open root for output directory")
    defer root.Close()

    // Read the generated file using os.Root
    content, err := root.ReadFile("generated_file.go")
    require.NoError(t, err, "Failed to read generated file")

    // ... verify content ...
}
```

### Fixture File Operations (S3 Tests)

```go
func TestS3Operations(t *testing.T) {
    // ... setup test ...

    t.Cleanup(func() {
        // Use os.OpenRoot to restrict file access to fixture directory
        root, err := os.OpenRoot("./fixtures")
        if err != nil {
            t.Fatal(err)
        }
        defer root.Close()

        // Safely remove test files
        err = root.Remove("download.json")
        if err != nil {
            t.Fatal(err)
        }
    })

    // Read test fixtures
    root, err := os.OpenRoot("./fixtures")
    if err != nil {
        t.Fatal(err)
    }
    defer root.Close()

    file, err := root.Open("test.json")
    if err != nil {
        t.Fatal(err)
    }
    defer file.Close()

    // ... use file ...
}
```

## Best Practices

1. **Always Close the Root**: Use `defer root.Close()` immediately after opening
2. **Use Relative Paths**: Only use relative paths, never absolute paths
3. **Validate Input**: Validate user-provided paths before using them with os.Root
4. **Error Handling**: Always check for errors from os.Root operations
5. **Root Selection**: Choose the most restrictive root directory possible

## Error Handling

os.Root operations return descriptive errors:

```go
_, err := root.ReadFile("nonexistent.txt")
if err != nil {
    // Error will indicate the specific issue
    fmt.Printf("Error: %v\n", err)
}
```

## Limitations

- Only works within a single directory tree
- Cannot access files outside the root directory
- Symbolic links are followed but cannot point outside the root
- Some platform-specific limitations apply (see Go documentation)

## Available Methods on os.Root

- `Open(name string) (*File, error)` - Open file/directory for reading
- `Create(name string) (*File, error)` - Create/truncate file for writing
- `ReadFile(name string) ([]byte, error)` - Read entire file
- `WriteFile(name string, data []byte, perm FileMode) error` - Write data to file
- `Stat(name string) (FileInfo, error)` - Get file information
- `Remove(name string) error` - Remove file/directory
- `Mkdir(name string, perm FileMode) error` - Create directory
- `MkdirAll(name string, perm FileMode) error` - Create directory tree
- `Close() error` - Close the root

For the complete API, see the [official Go documentation](https://pkg.go.dev/os#Root).

## Migration from Custom Wrappers

If you were previously using a custom wrapper around os.Root, migration is straightforward:

### Before (Custom Wrapper)
```go
fs, err := customwrapper.NewSafeFS(path)
defer fs.Close()
content, err := fs.ReadFile("file.txt")
```

### After (Standard os.Root)
```go
root, err := os.OpenRoot(path)
defer root.Close()
content, err := root.ReadFile("file.txt")
```

## Compatibility

- Requires Go 1.25 or later (Go 1.24 for initial implementation)
- Available on all platforms supported by Go
- Thread-safe for concurrent operations