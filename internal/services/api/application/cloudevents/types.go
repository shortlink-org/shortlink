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
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}
