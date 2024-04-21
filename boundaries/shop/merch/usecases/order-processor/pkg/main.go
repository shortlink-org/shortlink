package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"

	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
)

//nolint:revive // ignore
func getOrder(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

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
	ctx := context.Background()

	viper.SetDefault("HTTP_SERVER_TIMEOUT", "30s")
	viper.SetDefault("HTTP_SERVER_PORT", 6006) //nolint:revive,mnd // ignore

	// Create a new router and respond to POST /orders requests
	r := chi.NewMux()
	r.Post("/orders", getOrder)

	config := http_server.Config{
		Port:    viper.GetInt("HTTP_SERVER_PORT"),
		Timeout: viper.GetDuration("HTTP_SERVER_TIMEOUT"),
	}
	server := http_server.New(ctx, r, config, nil)

	// Start the server listening on port 6001
	// This is a blocking call
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Println("Error starting HTTP server")
	}
}
