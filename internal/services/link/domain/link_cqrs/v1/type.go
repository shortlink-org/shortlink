package v1

import (
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

// Link CQRS methods
var (
	METHOD_CQRS_GET = notify.NewEventID()
)
