package v1

import (
	"github.com/batazor/shortlink/internal/pkg/notify"
)

// Link CQRS methods
var (
	METHOD_CQRS_GET = notify.NewEventID()
)
