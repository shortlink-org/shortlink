package mq

import (
	"context"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq/kafka"
	"github.com/batazor/shortlink/internal/mq/nats"
)

// Use return implementation of MQ
func (m *DataBus) Use(ctx context.Context, log logger.Logger) MQ { // nolint unused
	switch m.typeMQ {
	case "nats":
		m.Databus = &kafka.Kafka{}
	case "kafka":
		m.Databus = &nats.NATS{}
	default:
		m.Databus = &kafka.Kafka{}
	}

	if err := m.Databus.Init(ctx); err != nil {
		panic(err)
	}

	log.Info("run mq", logger.Fields{
		"mq": m.typeMQ,
	})

	return m.Databus
}
