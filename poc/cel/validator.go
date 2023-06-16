/*
Example use CEL (Common Expression Language)
*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Claims ...
type Claims struct {
	Exp int64  `json:"exp"`
	Aud string `json:"aud"`
	// Add more fields as needed
}

// EvaluateInput ...
type EvaluateInput struct {
	Rule   string `json:"rule"`
	Claims Claims `json:"claims"`
	Now    int64  `json:"now"`
}

var expectedValue = "my-audience"

func main() {
	rulesPath := "./pkg/poc/cel/rules"
	rules, err := loadRules(rulesPath)
	if err != nil {
		fmt.Printf("Error loading rules: %v\n", err)
		os.Exit(1)
	}

	compiledRules, err := compileRules(rules)
	if err != nil {
		fmt.Printf("Error compiling rules: %v\n", err)
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", setupRoutes(r, compiledRules))
	http.ListenAndServe(":8080", r)
}
