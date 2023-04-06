package application

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

type Bot struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[any]

	mq  *v1.DataBus
	log logger.Logger
}

type Service interface {
	// Observer interface for subscribe on system event
	notify.Subscriber[any]

	Init() error
	Send(message string) error
}
