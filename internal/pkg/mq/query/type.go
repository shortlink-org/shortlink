package query

import (
	"context"
)

type Message struct {
	Key     []byte // routing key
	Payload []byte // payload
}

type ResponseMessage struct {
	Body    []byte
	Context context.Context
}

type Response struct {
	Key  []byte               // routing key
	Chan chan ResponseMessage // payload
}
