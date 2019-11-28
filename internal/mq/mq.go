package mq

type MQ interface {
	Init() error
	Close() error

	Send(message []byte) error
	Subscribe(message chan []byte)
}
