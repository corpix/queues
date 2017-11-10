package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/errors"
	"github.com/cryptounicorns/queues/message"
)

type Producer struct {
	topic         string
	client        sarama.Client
	kafkaProducer sarama.SyncProducer
	log           loggers.Logger
}

func (p *Producer) Produce(m message.Message) error {
	_, _, err := p.kafkaProducer.SendMessage(
		&sarama.ProducerMessage{
			Topic: p.topic,
			Value: sarama.StringEncoder(m),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) Close() error {
	var (
		err error
	)

	err = p.kafkaProducer.Close()
	if err != nil {
		return err
	}

	err = p.client.Close()
	if err != nil {
		return err
	}

	return nil
}

func NewProducer(c Config, l loggers.Logger) (*Producer, error) {
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	var (
		client        sarama.Client
		kafkaProducer sarama.SyncProducer
		err           error
	)

	if c.Kafka == nil {
		c.Kafka = sarama.NewConfig()
		c.Kafka.Producer.Return.Successes = true
	}

	client, err = sarama.NewClient(
		c.Addrs,
		c.Kafka,
	)
	if err != nil {
		return nil, err
	}

	kafkaProducer, err = sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	return &Producer{
		topic:         c.Topic,
		client:        client,
		kafkaProducer: kafkaProducer,
		log:           l,
	}, nil
}
