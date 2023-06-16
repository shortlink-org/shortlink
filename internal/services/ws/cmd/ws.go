package main

import (
	"net/http"

	"github.com/shortlink-org/shortlink/internal/services/ws"
)

func main() {
	http.HandleFunc("/ws", ws.Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
