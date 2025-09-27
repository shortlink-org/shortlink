package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	jsonv2 "github.com/go-json-experiment/json"
)

// Large data structure to trigger garbage collection
type LargeData struct {
	ID       int                    `json:"id"`
	Records  []Record               `json:"records"`
	Metadata map[string]interface{} `json:"metadata"`
	Payload  string                 `json:"payload"`
}

type Record struct {
	Timestamp time.Time              `json:"timestamp"`
	Values    []float64              `json:"values"`
	Tags      map[string]string      `json:"tags"`
	Content   string                 `json:"content"`
	Nested    map[string]interface{} `json:"nested"`
}

func generateLargeDataset(size int) LargeData {
	records := make([]Record, size)
	
	for i := 0; i < size; i++ {
		record := Record{
			Timestamp: time.Now().Add(time.Duration(i) * time.Second),
			Values:    make([]float64, 50),
			Tags: map[string]string{
				"type":        "measurement",
				"source":      fmt.Sprintf("sensor-%d", i%10),
				"location":    fmt.Sprintf("rack-%d", i%5),
				"environment": "production",
			},
			Content: fmt.Sprintf("This is record number %d with substantial content to increase memory usage and trigger garbage collection during JSON processing operations", i),
			Nested: map[string]interface{}{
				"level1": map[string]interface{}{
					"level2": map[string]interface{}{
						"level3": []string{"data1", "data2", "data3", "data4", "data5"},
						"numbers": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					},
				},
				"config": map[string]interface{}{
					"enabled":    true,
					"timeout":    30,
					"retries":    3,
					"buffer_size": 1024,
				},
			},
		}
		
		// Fill values array
		for j := range record.Values {
			record.Values[j] = float64(i*j) * 3.14159
		}
		
		records[i] = record
	}
	
	// Create large payload
	payload := ""
	for i := 0; i < 1000; i++ {
		payload += fmt.Sprintf("Large payload content block %d. ", i)
	}
	
	return LargeData{
		ID:      12345,
		Records: records,
		Metadata: map[string]interface{}{
			"version":      "2.1.0",
			"generated_at": time.Now(),
			"total_size":   size,
			"compression":  "none",
			"encoding":     "utf-8",
		},
		Payload: payload,
	}
}

func printGCStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("=== %s ===\n", label)
	fmt.Printf("Alloc: %d KB", bToKb(m.Alloc))
	fmt.Printf(", TotalAlloc: %d KB", bToKb(m.TotalAlloc))
	fmt.Printf(", Sys: %d KB", bToKb(m.Sys))
	fmt.Printf(", NumGC: %d\n", m.NumGC)
	
	gcStats := debug.GCStats{}
	debug.ReadGCStats(&gcStats)
	if len(gcStats.Pause) > 0 {
		fmt.Printf("Last GC Pause: %v\n", gcStats.Pause[0])
	}
	fmt.Println()
}

func bToKb(b uint64) uint64 {
	return b / 1024
}

func runGCIntensiveWorkload() {
	fmt.Println("=== GC-Intensive JSON Workload ===")
	
	// Check experimental features
	if goexp := os.Getenv("GOEXPERIMENT"); goexp != "" {
		fmt.Printf("GOEXPERIMENT: %s\n", goexp)
		if containsGreenTeaGC(goexp) {
			fmt.Println("✓ Green Tea GC is enabled")
		} else {
			fmt.Println("✗ Green Tea GC is not enabled")
		}
	} else {
		fmt.Println("GOEXPERIMENT not set - using default GC")
	}
	fmt.Println()
	
	printGCStats("Initial State")
	
	// Generate large dataset
	fmt.Println("Generating large dataset...")
	data := generateLargeDataset(1000)
	printGCStats("After Dataset Generation")
	
	// Force garbage collection
	runtime.GC()
	printGCStats("After Manual GC")
	
	// JSON processing loop that creates memory pressure
	fmt.Println("Starting JSON processing loop...")
	start := time.Now()
	
	for i := 0; i < 100; i++ {
		// Marshal large data
		jsonData, err := jsonv2.Marshal(data)
		if err != nil {
			panic(err)
		}
		
		// Unmarshal back
		var newData LargeData
		err = jsonv2.Unmarshal(jsonData, &newData)
		if err != nil {
			panic(err)
		}
		
		// Create some temporary allocations
		temp := make([]byte, 10*1024) // 10KB temporary allocation
		_ = temp
		
		if i%20 == 0 {
			printGCStats(fmt.Sprintf("Iteration %d", i))
		}
	}
	
	duration := time.Since(start)
	fmt.Printf("Processing completed in: %v\n", duration)
	
	printGCStats("Final State")
	
	// Final GC to clean up
	runtime.GC()
	printGCStats("After Final GC")
}

func containsGreenTeaGC(goexp string) bool {
	// Check if greenteagc is in the GOEXPERIMENT string
	return contains(goexp, "greenteagc")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && (s[:len(substr)+1] == substr+"," || 
		s[len(s)-len(substr)-1:] == ","+substr || 
		containsMiddle(s, ","+substr+","))))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("=== Go Experimental GC Demo ===")
	fmt.Println("This demo shows garbage collection behavior with heavy JSON processing")
	fmt.Println()
	
	// Set GC target percentage lower to trigger more frequent GC
	debug.SetGCPercent(50)
	fmt.Println("Set GC target to 50% to increase GC frequency")
	fmt.Println()
	
	runGCIntensiveWorkload()
	
	fmt.Println("=== Summary ===")
	fmt.Println("The Green Tea GC experiment aims to improve:")
	fmt.Println("- Reduced garbage collection pause times")
	fmt.Println("- Better memory utilization patterns") 
	fmt.Println("- Improved performance for allocation-heavy workloads")
	fmt.Println("- Enhanced scalability for large heap sizes")
	fmt.Println()
	fmt.Println("To see the difference, run this demo with and without GOEXPERIMENT=greenteagc")
}