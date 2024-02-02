package events

import "github.com/shortlink-org/shortlink/pkg/notify"

var (
	METHOD_NEW_LINK      = notify.NewEventID()
	METHOD_SEND_NEW_LINK = notify.NewEventID()
)
