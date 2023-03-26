package query

import (
	"context"
)

type Message struct {
	Key     []byte // routing key
	Payload []byte // payload
}

type ResponseMessage struct {
	Context context.Context
	Body    []byte
}

type Response struct {
	Chan chan ResponseMessage
	Key  []byte
}
