package query

type Message struct {
	Key     []byte // routing key
	Payload []byte // payload
}

type Response struct {
	Key  []byte      // routing key
	Chan chan []byte // payload
}
