package v1

import (
	"github.com/batazor/shortlink/internal/pkg/notify"
)

// Link CRUD methods
var (
	METHOD_ADD    = notify.NewEventID()
	METHOD_GET    = notify.NewEventID()
	METHOD_LIST   = notify.NewEventID()
	METHOD_UPDATE = notify.NewEventID()
	METHOD_DELETE = notify.NewEventID()
)

// MQ
const (
	// MQ_NEW - subscribe on new link
	MQ_NEW = "shortlink.link.new"
)
