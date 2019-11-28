package main

import (
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/mq/kafka"
)

func main() {
	var service mq.MQ

	service = &kafka.Kafka{}

	if err := service.Init(); err != nil {

	}
}
