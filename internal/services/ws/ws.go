package ws

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	BUFFER_SIZE = 1024
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
	hash := func(key string) string {
		h := sha1.New()
		h.Write([]byte(key))                                    // nolint:errcheck
		h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")) // nolint:errcheck

		return base64.StdEncoding.EncodeToString(h.Sum(nil))
	}(k)

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
	_, _ = bufrw.WriteString("HTTP/1.1 101 Switching Protocols\r\n")       // nolint:errcheck
	_, _ = bufrw.WriteString("Upgrade: websocket\r\n")                     // nolint:errcheck
	_, _ = bufrw.WriteString("Connection: Upgrade\r\n")                    // nolint:errcheck
	_, _ = bufrw.WriteString("Sec-Websocket-Accept: " + hash + "\r\n\r\n") // nolint:errcheck
	_ = bufrw.Flush()

	// respond to client
	buf := make([]byte, BUFFER_SIZE)
	for {
		n, err := bufrw.Read(buf)
		if err != nil {
			return
		}

		fmt.Println(buf[:n])
	}
}
