package kafka

import (
	"github.com/Shopify/sarama"

	"github.com/corpix/logger"
)

type Config struct {
	Addrs  []string
	Topic  string
	Sarama *sarama.Config
}

type Kafka struct {
	logger   logger.Logger
	config   Config
	producer sarama.SyncProducer
	consumer sarama.Consumer
}

func (e *Kafka) Produce(payload []byte) error {
	_, _, err := e.producer.SendMessage(
		&sarama.ProducerMessage{
			Topic: e.config.Topic,
			Value: sarama.StringEncoder(payload),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (e *Kafka) Consume() ([]byte, error) {
	pc, err := e.consumer.ConsumePartition(
		e.config.Topic,
		0,
		sarama.OffsetOldest,
	)
	select {
	case err := <-pc.Errors():
	case msg := <-consumer.Messages():
	}
	return nil, nil
}

func (e *Kafka) Close() error {
	var (
		err error
	)

	err = e.producer.Close()
	if err != nil {
		return err
	}
	return e.consumer.Close()
}

func NewFromConfig(l logger.Logger, c Config) (*Kafka, error) {
	var (
		config   = c.Sarama
		client   sarama.Client
		producer sarama.SyncProducer
		consumer sarama.Consumer
		err      error
	)

	if config == nil {
		config = sarama.NewConfig()
		config.Consumer.Return.Errors = true
		config.Producer.Return.Successes = true
	}

	client, err = sarama.NewClient(
		c.Addrs,
		config,
	)
	if err != nil {
		return nil, err
	}
	producer, err = sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}
	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}

	return &Kafka{
		l,
		c,
		producer,
		consumer,
	}, nil
}
