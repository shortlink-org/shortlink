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

// MQ events
const (
	// MQ_EVENT_LINK_NEW - subscribe on request created link
	MQ_EVENT_LINK_NEW = "shortlink.link.event.new"

	// MQ_EVENT_LINK_CREATED - subscribe on created link
	MQ_EVENT_LINK_CREATED = "shortlink.link.event.created"
)
