package kafka

import (
	"github.com/Shopify/sarama"

	"github.com/corpix/logger"

	"github.com/corpix/queues/handler"
	"github.com/corpix/queues/message"
)

type Kafka struct {
	logger       logger.Logger
	config       Config
	producer     sarama.SyncProducer
	consumer     sarama.Consumer
	consumerDone chan bool
}

func (e *Kafka) Produce(m message.Message) error {
	_, _, err := e.producer.SendMessage(
		&sarama.ProducerMessage{
			Topic: e.config.Topic,
			Value: sarama.StringEncoder(m),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (e *Kafka) Consume(h handler.Handler) error {
	var (
		c   sarama.PartitionConsumer
		err error
	)

	c, err = e.consumer.ConsumePartition(
		e.config.Topic,
		0,
		sarama.OffsetOldest,
	)
	if err != nil {
		return err
	}

	go e.consumeLoop(c, h)

	return nil
}

func (e *Kafka) consumeLoop(c sarama.PartitionConsumer, h handler.Handler) {
	var (
		msg *sarama.ConsumerMessage
		err error
	)

	for {
		select {
		case err = <-c.Errors():
			e.logger.Error(err)
		case msg = <-c.Messages():
			h(message.Message(msg.Value))
		case <-e.consumerDone:
			err = c.Close()
			if err != nil {
				e.logger.Error(err)
			}
			return
		}
	}
}

func (e *Kafka) Close() error {
	var (
		err error
	)

	err = e.producer.Close()
	if err != nil {
		return err
	}

	err = e.consumer.Close()
	if err != nil {
		return err
	}

	e.consumerDone <- true
	close(e.consumerDone)

	return nil
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
		make(chan bool),
	}, nil
}
