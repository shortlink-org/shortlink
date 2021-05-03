package cloudevents

import (
	"context"
)

// API ...
type API struct { // nolint unused
	ctx context.Context
}

// Example message
type Example struct { // nolint unused
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}
