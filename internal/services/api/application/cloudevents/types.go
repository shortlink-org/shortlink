package cloudevents

import (
	"context"
)

// API ...
type API struct {
	ctx context.Context
}

// Example message
type Example struct {
	Message  string `json:"message"`
	Sequence int    `json:"id"`
}
