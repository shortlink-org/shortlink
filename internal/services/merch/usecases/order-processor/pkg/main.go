package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body:", err.Error())
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}

	fmt.Println("Order received:", string(data))

	// response
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(data)
	if err != nil {
		log.Println("Error writing the response:", err.Error())
		http.Error(w, "Error writing the response", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Create a new router and respond to POST /orders requests
	r := chi.NewMux()
	r.Post("/orders", getOrder)

	// Start the server listening on port 6001
	// This is a blocking call
	err := http.ListenAndServe(":6006", r)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Println("Error starting HTTP server")
	}
}
