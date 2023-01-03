package v1

import (
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

var (
	// Link CRUD methods
	METHOD_ADD    = notify.NewEventID()
	METHOD_GET    = notify.NewEventID()
	METHOD_LIST   = notify.NewEventID()
	METHOD_UPDATE = notify.NewEventID()
	METHOD_DELETE = notify.NewEventID()
)

const (
	MQ_EVENT_CQRS_NEW = "shortlink.metadata.cqrs.new"
)
