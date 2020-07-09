package api_type

import (
	"time"

	"github.com/batazor/shortlink/internal/notify"
)

var (
	METHOD_ADD    = notify.NewEventID()
	METHOD_GET    = notify.NewEventID()
	METHOD_LIST   = notify.NewEventID()
	METHOD_UPDATE = notify.NewEventID()
	METHOD_DELETE = notify.NewEventID()
)

// Config - base configuration for API
type Config struct { // nolint unused
	Port    int
	Timeout time.Duration
}
