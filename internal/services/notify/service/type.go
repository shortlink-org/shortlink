package service

import (
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Bot struct {
	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event

	MQ  mq.MQ
	Log logger.Logger
}

type Service interface {
	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event

	Init() error
	Send(message string) error
}
