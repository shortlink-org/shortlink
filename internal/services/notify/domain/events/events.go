package events

import "github.com/shortlink-org/shortlink/internal/pkg/notify"

var (
	METHOD_NEW_LINK      = notify.NewEventID()
	METHOD_SEND_NEW_LINK = notify.NewEventID()
)
