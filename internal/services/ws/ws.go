package ws

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// check headers
	if r.Header.Get("Upgrade") != "websocket" {
		http.Error(w, "Not a websocket handshake", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Connection") != "Upgrade" {
		http.Error(w, "Not a websocket handshake", http.StatusBadRequest)
		return
	}

	k := r.Header.Get("Sec-WebSocket-Key")
	if k == "" {
		http.Error(w, "Not a websocket handshake", http.StatusBadRequest)
		return
	}

	// calculate response
	sum := k + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	hash := sha1.Sum([]byte(sum))
	str := base64.StdEncoding.EncodeToString(hash[:])

	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Not a websocket handshake", http.StatusBadRequest)
		return
	}

	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, "Not a websocket handshake", http.StatusBadRequest)
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	// write headers
	_, _ = bufrw.WriteString("HTTP/1.1 101 Switching Protocols\r\n")
	_, _ = bufrw.WriteString("Upgrade: websocket\r\n")
	_, _ = bufrw.WriteString("Connection: Upgrade\r\n")
	_, _ = bufrw.WriteString("Sec-Websocket-Accept: " + str + "\r\n\r\n")
	_ = bufrw.Flush()

	// respond to client
	buf := make([]byte, 1024)
	for {
		n, err := bufrw.Read(buf)
		if err != nil {
			return
		}

		fmt.Println(buf[:n])
	}
}
