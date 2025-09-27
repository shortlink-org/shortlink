package main

import (
	"encoding/json"
	"fmt"
	"time"

	jsonv2 "github.com/go-json-experiment/json"
)

// Benchmark struct for testing
type BenchmarkData struct {
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	Values      []float64              `json:"values"`
	Metadata    map[string]interface{} `json:"metadata"`
	Nested      []NestedStruct         `json:"nested"`
	Timestamp   time.Time              `json:"timestamp"`
	Enabled     bool                   `json:"enabled"`
	Description string                 `json:"description"`
}

type NestedStruct struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
	Tags  []string `json:"tags"`
}

func generateTestData() BenchmarkData {
	return BenchmarkData{
		ID:   12345,
		Name: "Test Benchmark Data with a reasonably long name to simulate real-world usage",
		Values: []float64{
			1.23456789, 2.34567890, 3.45678901, 4.56789012, 5.67890123,
			6.78901234, 7.89012345, 8.90123456, 9.01234567, 0.12345678,
		},
		Metadata: map[string]interface{}{
			"version":     "1.2.3",
			"environment": "production",
			"region":      "us-west-2",
			"datacenter":  "dc-01",
			"cluster":     "web-cluster-prod",
		},
		Nested: []NestedStruct{
			{Key: "config1", Value: 100, Tags: []string{"important", "config", "primary"}},
			{Key: "config2", Value: 200, Tags: []string{"secondary", "backup"}},
			{Key: "config3", Value: 300, Tags: []string{"tertiary", "fallback", "emergency"}},
		},
		Timestamp:   time.Now(),
		Enabled:     true,
		Description: "This is a comprehensive test data structure designed to benchmark JSON marshaling and unmarshaling performance between the standard encoding/json package and the experimental json/v2 implementation. It includes various data types including strings, numbers, arrays, nested objects, and maps to provide a realistic workload.",
	}
}

func benchmarkJSONv1(data BenchmarkData, iterations int) time.Duration {
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		// Marshal
		marshaled, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		
		// Unmarshal
		var unmarshaled BenchmarkData
		err = json.Unmarshal(marshaled, &unmarshaled)
		if err != nil {
			panic(err)
		}
	}
	
	return time.Since(start)
}

func benchmarkJSONv2(data BenchmarkData, iterations int) time.Duration {
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		// Marshal
		marshaled, err := jsonv2.Marshal(data)
		if err != nil {
			panic(err)
		}
		
		// Unmarshal
		var unmarshaled BenchmarkData
		err = jsonv2.Unmarshal(marshaled, &unmarshaled)
		if err != nil {
			panic(err)
		}
	}
	
	return time.Since(start)
}

func main() {
	fmt.Println("=== JSON Performance Benchmark ===")
	fmt.Println()
	
	data := generateTestData()
	
	// Show sample data size
	sampleJSON, _ := json.Marshal(data)
	fmt.Printf("Sample JSON size: %d bytes\n", len(sampleJSON))
	fmt.Printf("Sample JSON preview: %.100s...\n", string(sampleJSON))
	fmt.Println()
	
	// Different iteration counts for different scales
	testCases := []struct {
		name       string
		iterations int
	}{
		{"Small scale", 1000},
		{"Medium scale", 5000},
		{"Large scale", 10000},
	}
	
	for _, tc := range testCases {
		fmt.Printf("=== %s (%d iterations) ===\n", tc.name, tc.iterations)
		
		// Benchmark JSON v1
		v1Duration := benchmarkJSONv1(data, tc.iterations)
		fmt.Printf("encoding/json (v1):    %v\n", v1Duration)
		
		// Benchmark JSON v2
		v2Duration := benchmarkJSONv2(data, tc.iterations)
		fmt.Printf("json/v2 (experimental): %v\n", v2Duration)
		
		// Calculate improvement
		if v2Duration < v1Duration {
			improvement := float64(v1Duration-v2Duration) / float64(v1Duration) * 100
			fmt.Printf("json/v2 is %.1f%% faster\n", improvement)
		} else {
			regression := float64(v2Duration-v1Duration) / float64(v1Duration) * 100
			fmt.Printf("json/v2 is %.1f%% slower\n", regression)
		}
		
		// Operations per second
		v1Ops := float64(tc.iterations) / v1Duration.Seconds()
		v2Ops := float64(tc.iterations) / v2Duration.Seconds()
		fmt.Printf("Operations/sec - v1: %.0f, v2: %.0f\n", v1Ops, v2Ops)
		fmt.Println()
	}
	
	fmt.Println("Note: Performance may vary based on data structure complexity,")
	fmt.Println("system resources, and Go runtime optimizations.")
}