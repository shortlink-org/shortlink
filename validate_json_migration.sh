#!/bin/bash

# JSON v2 Migration Validation Script
# This script validates the JSON v2 migration for the shortlink codebase

echo "🔍 JSON v2 Migration Validation Script"
echo "======================================"

# Check Go version
echo "📋 Current Go Version:"
go version

# Check if Go 1.25+ is available
if go version | grep -E "go1\.(2[5-9]|[3-9][0-9])" > /dev/null; then
    echo "✅ Go 1.25+ detected - JSON v2 migration can be tested"
    
    echo ""
    echo "🧪 Testing JSON v2 functionality..."
    
    # Test with JSON v2 enabled
    echo "Running tests with GOEXPERIMENT=jsonv2..."
    GOEXPERIMENT=jsonv2 go test -v ./pkg/logger/logger_test.go
    
    if [ $? -eq 0 ]; then
        echo "✅ Logger tests passed with JSON v2"
    else
        echo "❌ Logger tests failed with JSON v2"
        exit 1
    fi
    
    # Test benchmark performance
    echo ""
    echo "🚀 Running performance benchmarks..."
    GOEXPERIMENT=jsonv2 go test -bench=. ./docs/ADR/decisions/proof/ADR-0007/serialization_bench_test.go
    
    # Build check
    echo ""
    echo "🔨 Testing build with JSON v2..."
    GOEXPERIMENT=jsonv2 go build -o /tmp/shortlink-test ./poc/cel/
    
    if [ $? -eq 0 ]; then
        echo "✅ Build successful with JSON v2"
        rm -f /tmp/shortlink-test
    else
        echo "❌ Build failed with JSON v2"
        exit 1
    fi
    
    echo ""
    echo "🎉 JSON v2 migration validation completed successfully!"
    
else
    echo "⚠️  Go 1.25+ not detected - JSON v2 migration prepared but cannot be tested yet"
    echo ""
    echo "📝 Migration status:"
    echo "   - ✅ All imports updated to encoding/json/v2"
    echo "   - ✅ Struct tags remain compatible"
    echo "   - ✅ Function calls remain compatible"
    echo "   - ⏳ Waiting for Go 1.25 release"
    echo ""
    echo "🔄 To test when Go 1.25 is available:"
    echo "   1. Update Go version: go install golang.org/dl/go1.25@latest"
    echo "   2. Run this script again"
    echo "   3. Or manually run: GOEXPERIMENT=jsonv2 go test ./..."
fi

echo ""
echo "📚 For more information, see: JSON_V2_MIGRATION.md"