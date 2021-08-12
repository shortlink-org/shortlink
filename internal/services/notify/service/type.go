package service

import (
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Bot struct {
	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event

	MQ  v1.MQ
	Log logger.Logger
}

type Service interface {
	// Observer interface for subscribe on system event
	notify.Subscriber // Observer interface for subscribe on system event

	Init() error
	Send(message string) error
}
