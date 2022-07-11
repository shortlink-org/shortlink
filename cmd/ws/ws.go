package main

import (
	"net/http"

	"github.com/batazor/shortlink/internal/services/ws"
)

func main() {
	http.HandleFunc("/ws", ws.Handler)
	http.ListenAndServe(":8080", nil)
}
