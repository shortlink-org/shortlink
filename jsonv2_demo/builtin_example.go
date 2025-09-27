package main

import (
	"fmt"
	"log"
	"os"
	"time"

	// When GOEXPERIMENT=jsonv2 is set, these imports will work
	// Standard encoding/json for comparison
	jsonv1 "encoding/json"
)

// Example showing how to use builtin json/v2 when GOEXPERIMENT=jsonv2 is enabled
// Note: This file demonstrates the syntax, but may not compile without GOEXPERIMENT=jsonv2

// Example struct
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Profile   Profile   `json:"profile"`
}

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio,omitempty"`
}

func main() {
	fmt.Println("=== Builtin encoding/json/v2 Example ===")
	fmt.Println("Note: This requires GOEXPERIMENT=jsonv2 to be set")
	fmt.Println()

	// Check environment
	if goexp := os.Getenv("GOEXPERIMENT"); goexp != "" {
		fmt.Printf("GOEXPERIMENT: %s\n", goexp)
	} else {
		fmt.Println("GOEXPERIMENT not set - builtin json/v2 not available")
		fmt.Println("Use 'export GOEXPERIMENT=jsonv2' to enable")
	}
	fmt.Println()

	// Sample data
	user := User{
		ID:        123,
		Username:  "johndoe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		Profile: Profile{
			FirstName: "John",
			LastName:  "Doe",
			Bio:       "Software developer interested in Go and JSON processing",
		},
	}

	// Standard JSON v1 example
	fmt.Println("=== Standard encoding/json (v1) ===")
	data, err := jsonv1.Marshal(user)
	if err != nil {
		log.Printf("JSON v1 marshal error: %v", err)
		return
	}
	fmt.Printf("JSON v1 output:\n%s\n\n", string(data))

	// When GOEXPERIMENT=jsonv2 is enabled, you can use:
	// 
	// import (
	//     jsonv2 "encoding/json/v2"
	//     "encoding/json/v2/jsontext"
	// )
	//
	// And then use the new API:
	//
	// fmt.Println("=== Builtin encoding/json/v2 ===")
	// 
	// // New marshal options
	// opts := jsonv2.MarshalOptions{
	//     Indent: "  ",
	// }
	// data2, err := opts.Marshal(user)
	// if err != nil {
	//     log.Printf("JSON v2 marshal error: %v", err)
	//     return
	// }
	// fmt.Printf("JSON v2 with options:\n%s\n", string(data2))
	// 
	// // Low-level streaming with jsontext
	// var buf []byte
	// enc := jsontext.NewEncoder(&buf)
	// enc.WriteToken(jsontext.ObjectStart)
	// enc.WriteToken(jsontext.String("message"))
	// enc.WriteToken(jsontext.String("Hello from builtin json/v2!"))
	// enc.WriteToken(jsontext.ObjectEnd)
	// fmt.Printf("Streaming: %s\n", string(buf))

	fmt.Println("To see the full json/v2 features, run with GOEXPERIMENT=jsonv2")
	fmt.Println("The experimental package used in main.go provides similar functionality")
}