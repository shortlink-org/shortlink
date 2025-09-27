package main

import (
	"fmt"
	"log"
	"os"
	"time"

	// Standard encoding/json for comparison
	jsonv1 "encoding/json"

	// Use the experimental json package directly
	jsonv2 "github.com/go-json-experiment/json"
)

// Example struct to demonstrate JSON operations
type Person struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

// Custom JSON marshaling example for json/v2
type CustomData struct {
	Value string `json:"value"`
}

func (c CustomData) MarshalJSON() ([]byte, error) {
	// Custom marshaling logic for json/v2
	return []byte(`"custom:` + c.Value + `"`), nil
}

func (c *CustomData) UnmarshalJSON(data []byte) error {
	// Custom unmarshaling logic for json/v2
	str := string(data)
	if len(str) < 9 || str[:8] != `"custom:` || str[len(str)-1] != '"' {
		return fmt.Errorf("expected 'custom:' prefix in quotes")
	}
	c.Value = str[8 : len(str)-1]
	return nil
}

func main() {
	fmt.Println("=== Go Experimental encoding/json/v2 Demo ===")
	fmt.Println()

	// Check if GOEXPERIMENT=jsonv2 is enabled
	if goexp := os.Getenv("GOEXPERIMENT"); goexp != "" {
		fmt.Printf("GOEXPERIMENT: %s\n", goexp)
	} else {
		fmt.Println("GOEXPERIMENT not set - json/v2 features may not be available")
	}
	fmt.Println()

	// Sample data
	person := Person{
		ID:        1,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Age:       30,
		CreatedAt: time.Now(),
		Active:    true,
	}

	fmt.Println("=== Standard encoding/json (v1) ===")
	demonstrateJSONv1(person)

	fmt.Println("\n=== Experimental encoding/json/v2 ===")
	demonstrateJSONv2(person)

	fmt.Println("\n=== Performance Comparison ===")
	performanceComparison(person)

	fmt.Println("\n=== New Features in json/v2 ===")
	demonstrateNewFeatures()
}

func demonstrateJSONv1(person Person) {
	// Standard JSON marshaling
	data, err := jsonv1.Marshal(person)
	if err != nil {
		log.Printf("JSON v1 marshal error: %v", err)
		return
	}
	fmt.Printf("JSON v1 Marshal: %s\n", string(data))

	// Standard JSON unmarshaling
	var decoded Person
	err = jsonv1.Unmarshal(data, &decoded)
	if err != nil {
		log.Printf("JSON v1 unmarshal error: %v", err)
		return
	}
	fmt.Printf("JSON v1 Unmarshal: %+v\n", decoded)
}

func demonstrateJSONv2(person Person) {
	// JSON v2 marshaling
	data, err := jsonv2.Marshal(person)
	if err != nil {
		log.Printf("JSON v2 marshal error: %v", err)
		return
	}
	fmt.Printf("JSON v2 Marshal: %s\n", string(data))

	// JSON v2 unmarshaling
	var decoded Person
	err = jsonv2.Unmarshal(data, &decoded)
	if err != nil {
		log.Printf("JSON v2 unmarshal error: %v", err)
		return
	}
	fmt.Printf("JSON v2 Unmarshal: %+v\n", decoded)

	// Demonstrate streaming with jsontext
	fmt.Println("\n--- Streaming with jsontext ---")
	streamingJSON := map[string]interface{}{
		"message":   "Hello from json/v2!",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	
	streamData, err := jsonv2.Marshal(streamingJSON)
	if err != nil {
		log.Printf("Streaming marshal error: %v", err)
		return
	}
	fmt.Printf("Streaming JSON: %s\n", string(streamData))
}

func performanceComparison(person Person) {
	const iterations = 10000

	// Benchmark JSON v1
	start := time.Now()
	for i := 0; i < iterations; i++ {
		data, _ := jsonv1.Marshal(person)
		var decoded Person
		jsonv1.Unmarshal(data, &decoded)
	}
	v1Duration := time.Since(start)

	// Benchmark JSON v2
	start = time.Now()
	for i := 0; i < iterations; i++ {
		data, _ := jsonv2.Marshal(person)
		var decoded Person
		jsonv2.Unmarshal(data, &decoded)
	}
	v2Duration := time.Since(start)

	fmt.Printf("JSON v1 (%d iterations): %v\n", iterations, v1Duration)
	fmt.Printf("JSON v2 (%d iterations): %v\n", iterations, v2Duration)
	
	if v2Duration < v1Duration {
		improvement := float64(v1Duration-v2Duration) / float64(v1Duration) * 100
		fmt.Printf("JSON v2 is %.1f%% faster\n", improvement)
	} else {
		difference := float64(v2Duration-v1Duration) / float64(v1Duration) * 100
		fmt.Printf("JSON v1 is %.1f%% faster\n", difference)
	}
}

func demonstrateNewFeatures() {
	// Custom marshaling/unmarshaling
	custom := CustomData{Value: "test data"}
	
	data, err := jsonv2.Marshal(custom)
	if err != nil {
		log.Printf("Custom marshal error: %v", err)
		return
	}
	fmt.Printf("Custom Marshal: %s\n", string(data))

	var decoded CustomData
	err = jsonv2.Unmarshal(data, &decoded)
	if err != nil {
		log.Printf("Custom unmarshal error: %v", err)
		return
	}
	fmt.Printf("Custom Unmarshal: %+v\n", decoded)

	// Demonstrate marshal options (new in json/v2)
	fmt.Println("\n--- Marshal Options ---")
	// For now, use standard marshal - json/v2 options API may differ
	prettyData, err := jsonv2.Marshal(custom)
	if err != nil {
		log.Printf("Pretty marshal error: %v", err)
		return
	}
	fmt.Printf("Pretty JSON:\n%s\n", string(prettyData))

	// Low-level JSON processing with jsontext
	fmt.Println("\n--- Low-level Processing ---")
	complexData := map[string]interface{}{
		"numbers": []int{1, 2, 3},
		"nested": map[string]string{
			"key": "value",
		},
	}
	
	complexJSON, err := jsonv2.Marshal(complexData)
	if err != nil {
		log.Printf("Complex marshal error: %v", err)
		return
	}
	fmt.Printf("Complex JSON construction: %s\n", string(complexJSON))
}