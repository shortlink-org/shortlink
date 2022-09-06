package main

import (
	"net/http"

	"github.com/batazor/shortlink/internal/services/ws"
)

func main() {
	http.HandleFunc("/ws", ws.Handler)
	_ = http.ListenAndServe(":8080", nil)
}
