# Go Experimental encoding/json/v2 Demo

This demonstration shows how to enable and use the experimental `encoding/json/v2` package introduced in Go 1.25.

## Requirements

- Go 1.25.1 or later
- Set `GOEXPERIMENT=jsonv2` environment variable

## What is encoding/json/v2?

The experimental `encoding/json/v2` package is a major revision of the existing `encoding/json` package that addresses various issues and enhances performance. Key improvements include:

- **Better Performance**: Substantially faster decoding in many scenarios
- **Enhanced API**: New marshal/unmarshal options and configuration
- **Lower-level Control**: The `encoding/json/jsontext` package provides fine-grained JSON processing
- **Improved Error Messages**: More detailed and helpful error reporting

## Enabling the Experiment

To enable the experimental json/v2 package, set the `GOEXPERIMENT` environment variable:

```bash
export GOEXPERIMENT=jsonv2
```

This enables two new packages:
- `encoding/json/v2`: The main enhanced JSON package
- `encoding/json/jsontext`: Lower-level JSON processing

## Running the Demo

### Method 1: Using the provided script
```bash
chmod +x run_demo.sh
./run_demo.sh
```

### Method 2: Manual execution
```bash
export GOEXPERIMENT=jsonv2
go run main.go
```

### Method 3: Build and run
```bash
export GOEXPERIMENT=jsonv2
go build -o demo main.go
./demo
```

## Features Demonstrated

1. **Basic Marshaling/Unmarshaling**: Comparison between json/v1 and json/v2
2. **Performance Comparison**: Benchmark showing speed differences
3. **Custom Marshaling**: Implementing custom MarshalJSONV2/UnmarshalJSONV2 methods
4. **Marshal Options**: Using the new options system for pretty printing
5. **Low-level Processing**: Manual JSON construction with jsontext package
6. **Streaming**: Building JSON incrementally with the encoder

## Key Differences from encoding/json

- Custom marshal/unmarshal methods use `MarshalJSONV2`/`UnmarshalJSONV2` instead of `MarshalJSON`/`UnmarshalJSON`
- New `MarshalOptions` and `UnmarshalOptions` for configuration
- The `jsontext` package provides token-level JSON processing
- Generally faster decoding performance
- More detailed error messages

## Important Notes

- This is an **experimental** package and not subject to Go 1 compatibility promise
- The API may change in future releases
- Test your existing code with `GOEXPERIMENT=jsonv2` to detect compatibility issues
- Provide feedback to the Go team to help improve the final implementation

## Dependencies

The demo uses the experimental json package:
```
github.com/go-json-experiment/json v0.0.0-20250725192818-e39067aee2d2
```