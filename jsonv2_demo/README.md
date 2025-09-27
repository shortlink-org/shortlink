# Go Experimental Features Demo

This demonstration shows how to enable and use experimental features in Go 1.25, including:
- `encoding/json/v2`: Enhanced JSON processing package
- `greenteagc`: New garbage collector implementation

## Requirements

- Go 1.25.1 or later
- Set `GOEXPERIMENT=greenteagc,jsonv2` environment variable

## What is encoding/json/v2?

The experimental `encoding/json/v2` package is a major revision of the existing `encoding/json` package that addresses various issues and enhances performance. Key improvements include:

- **Better Performance**: Substantially faster decoding in many scenarios
- **Enhanced API**: New marshal/unmarshal options and configuration
- **Lower-level Control**: The `encoding/json/jsontext` package provides fine-grained JSON processing
- **Improved Error Messages**: More detailed and helpful error reporting

## What is Green Tea GC?

The experimental `greenteagc` is a new garbage collector implementation that aims to improve:
- **Reduced Pause Times**: Lower stop-the-world garbage collection pauses
- **Better Memory Utilization**: More efficient memory management patterns
- **Improved Scalability**: Better performance for applications with large heap sizes
- **Enhanced Throughput**: Better performance for allocation-heavy workloads

## Enabling the Experiments

To enable both experimental features, set the `GOEXPERIMENT` environment variable:

```bash
export GOEXPERIMENT=greenteagc,jsonv2
```

This enables:
- `greenteagc`: New garbage collector
- `jsonv2`: Enhanced JSON processing
- `encoding/json/v2`: The main enhanced JSON package (when using builtin)
- `encoding/json/jsontext`: Lower-level JSON processing (when using builtin)

## Running the Demo

### Method 1: Using the provided script
```bash
chmod +x run_demo.sh
./run_demo.sh
```

### Method 2: Manual execution
```bash
export GOEXPERIMENT=greenteagc,jsonv2
go run main.go
```

### Method 3: Build and run
```bash
export GOEXPERIMENT=greenteagc,jsonv2
go build -o demo main.go
./demo
```

### Method 4: Run GC-intensive demo
```bash
export GOEXPERIMENT=greenteagc,jsonv2
go run gc_demo.go
```

### Method 5: Run performance benchmark
```bash
export GOEXPERIMENT=greenteagc,jsonv2
go run benchmark.go
```

## Features Demonstrated

### JSON/v2 Features
1. **Basic Marshaling/Unmarshaling**: Comparison between json/v1 and json/v2
2. **Performance Comparison**: Benchmark showing speed differences
3. **Custom Marshaling**: Implementing custom JSON marshaling methods
4. **Marshal Options**: Using enhanced marshaling capabilities
5. **Low-level Processing**: Complex JSON construction and processing
6. **Streaming**: Building JSON incrementally

### Green Tea GC Features
1. **Memory Usage Monitoring**: Track allocation patterns and GC behavior
2. **GC-Intensive Workloads**: Heavy allocation scenarios to test GC performance
3. **Pause Time Analysis**: Measure garbage collection pause times
4. **Memory Pressure Testing**: Large dataset processing to trigger GC cycles

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