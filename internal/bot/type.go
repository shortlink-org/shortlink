package bot

import (
	"github.com/batazor/shortlink/internal/notify"
)

const (
	METHOD_NEW_LINK = 10 + iota // nolint unused
)

type Bot struct {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event

	services []Service
}

type Service interface {
	Init() error
	Send(message string) error
}
