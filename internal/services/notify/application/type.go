package application

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

type Bot struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[any]

	mq  mq.MQ
	log logger.Logger
}

type Service interface {
	// Observer interface for subscribe on system event
	notify.Subscriber[any]

	Init() error
	Send(message string) error
}
