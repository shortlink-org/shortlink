#!/bin/bash

echo "=== Go Experimental Features Complete Demo Suite ==="
echo "This script runs all demonstrations with GOEXPERIMENT=greenteagc,jsonv2"
echo ""

# Set experimental features
export GOEXPERIMENT=greenteagc,jsonv2

echo "Enabled experimental features: $GOEXPERIMENT"
echo "- greenteagc: New garbage collector with improved pause times"
echo "- jsonv2: Enhanced JSON processing with better performance"
echo ""

# Function to run a demo with error handling
run_demo() {
    local demo_name="$1"
    local demo_file="$2"
    
    echo "==============================================="
    echo "Running: $demo_name"
    echo "==============================================="
    
    if [ -f "$demo_file" ]; then
        go run "$demo_file"
        if [ $? -eq 0 ]; then
            echo "✓ $demo_name completed successfully"
        else
            echo "✗ $demo_name failed"
        fi
    else
        echo "✗ Demo file $demo_file not found"
    fi
    
    echo ""
    echo "Press Enter to continue to next demo..."
    read
    echo ""
}

# Run all demonstrations
echo "Starting complete demo suite..."
echo ""

run_demo "Main JSON/v2 Features Demo" "main.go"
run_demo "GC-Intensive Workload Demo" "gc_demo.go" 
run_demo "Performance Benchmark" "benchmark.go"
run_demo "Builtin JSON/v2 Example" "builtin_example.go"

echo "==============================================="
echo "All demos completed!"
echo "==============================================="
echo ""
echo "Summary of what was demonstrated:"
echo "1. JSON/v2 enhanced marshaling and unmarshaling"
echo "2. Performance comparison between json/v1 and json/v2"
echo "3. Garbage collection behavior with Green Tea GC"
echo "4. Memory usage patterns and GC pause times"
echo "5. Large-scale JSON processing workloads"
echo ""
echo "To compare with default behavior, run without GOEXPERIMENT:"
echo "  unset GOEXPERIMENT"
echo "  go run main.go"
echo "  go run gc_demo.go"