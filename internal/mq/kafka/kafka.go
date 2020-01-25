package kafka

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/mq/query"
)

type Config struct { // nolint unused
	URI string
}

type Kafka struct { // nolint unused
	*Config
	client   sarama.Client
	producer sarama.AsyncProducer
	consumer sarama.Consumer
}

func (mq *Kafka) Init(ctx context.Context) error { // nolint unparam
	var err error

	// Set configuration
	mq.setConfig()

	config := sarama.NewConfig()

	if mq.client, err = sarama.NewClient([]string{mq.Config.URI}, config); err != nil {
		return err
	}

	if mq.producer, err = sarama.NewAsyncProducerFromClient(mq.client); err != nil {
		return err
	}

	if mq.consumer, err = sarama.NewConsumerFromClient(mq.client); err != nil {
		return err
	}

	return nil
}

func (mq *Kafka) Close() error {
	var err error

	if mq.client != nil {
		err = mq.client.Close()
	}

	if mq.producer != nil {
		err = mq.producer.Close()
	}

	if mq.consumer != nil {
		err = mq.consumer.Close()
	}

	return err
}

func (k *Kafka) Publish(message query.Message) error {
	k.producer.Input() <- &sarama.ProducerMessage{
		Topic: "shortlink",
		Key:   sarama.StringEncoder(message.Key),
		Value: sarama.ByteEncoder(message.Payload),
	}

	return nil
}

func (mq *Kafka) Subscribe(message query.Response) error {
	consumer, err := mq.consumer.ConsumePartition("shortlink", 1, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-consumer.Errors():
			message.Chan <- []byte(err.Error())
		case msg := <-consumer.Messages():
			message.Chan <- msg.Value
		}
	}
}

func (mq *Kafka) UnSubscribe() error {
	panic("implement me!")
}

// setConfig - set configuration
func (mq *Kafka) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_KAFKA_URI", "localhost:9092")
	mq.Config = &Config{
		URI: viper.GetString("MQ_KAFKA_URI"),
	}
}
