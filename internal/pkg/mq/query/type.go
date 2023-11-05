package query

import (
	"context"
)

type ResponseMessage struct {
	Context context.Context
	Body    []byte
}

type Response struct {
	Chan chan ResponseMessage
}
