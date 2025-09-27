#!/bin/bash

echo "=== Go Experimental encoding/json/v2 Demo ==="
echo "Setting GOEXPERIMENT=jsonv2 to enable experimental features"
echo ""

# Set the experimental flag and run the demo
export GOEXPERIMENT=jsonv2

echo "Building with GOEXPERIMENT=jsonv2..."
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
export GOEXPERIMENT=jsonv2
go run main.go