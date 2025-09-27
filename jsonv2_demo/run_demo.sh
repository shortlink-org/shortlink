#!/bin/bash

echo "=== Go Experimental Features Demo ==="
echo "Setting GOEXPERIMENT=greenteagc,jsonv2 to enable experimental features"
echo "- greenteagc: New garbage collector"
echo "- jsonv2: Enhanced JSON processing"
echo ""

# Set the experimental flags and run the demo
export GOEXPERIMENT=greenteagc,jsonv2

echo "Building with GOEXPERIMENT=greenteagc,jsonv2..."
go build -o demo main.go

if [ $? -eq 0 ]; then
    echo "Build successful! Running demo..."
    echo ""
    ./demo
else
    echo "Build failed. Trying without experimental flag..."
    unset GOEXPERIMENT
    go build -o demo main.go
    if [ $? -eq 0 ]; then
        echo "Build successful without experimental flag. Running demo..."
        echo ""
        ./demo
    else
        echo "Build failed completely."
        exit 1
    fi
fi

echo ""
echo "=== Alternative: Running directly with go run ==="
export GOEXPERIMENT=greenteagc,jsonv2
go run main.go