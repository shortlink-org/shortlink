package rabbit

import (
	"context"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"github.com/batazor/shortlink/internal/mq/query"
)

type RabbitMQ struct {
	URI  string
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func (mq *RabbitMQ) Init(ctx context.Context) error {
	var err error

	// Set configuration
	mq.setConfig()

	// connect to RabbitMQ server
	mq.conn, err = amqp.Dial(mq.URI)
	if err != nil {
		return err
	}

	// create a channel
	mq.ch, err = mq.conn.Channel()
	if err != nil {
		return err
	}

	// create a queue
	mq.q, err = mq.ch.QueueDeclare(
		"shortlink", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	return nil
}

func (mq *RabbitMQ) Publish(message query.Message) error {
	if err := mq.ch.Publish(
		string(message.Key), // exchange
		mq.q.Name,           // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message.Payload,
		}); err != nil {
		return err
	}

	return nil
}

func (mq *RabbitMQ) Subscribe(message query.Response) error {
	msgs, err := mq.ch.Consume(
		mq.q.Name, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			message.Chan <- d.Body
		}
	}()

	return nil
}

func (mq *RabbitMQ) UnSubscribe() error {
	return nil
}

func (mq *RabbitMQ) Close() error {
	if err := mq.conn.Close(); err != nil {
		return err
	}

	if err := mq.ch.Close(); err != nil {
		return err
	}

	return nil
}

// setConfig - Construct a new RabbitMQ configuration.
func (mq *RabbitMQ) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_RABBIT_URI", "amqp://localhost:5672") // RabbitMQ URI

	mq.URI = viper.GetString("MQ_RABBIT_URI")

	return
}
