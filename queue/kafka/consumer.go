package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/result"
)

type Consumer struct {
	config                 Config
	client                 sarama.Client
	kafkaConsumer          sarama.Consumer
	kafkaPartitionConsumer sarama.PartitionConsumer
	stream                 chan result.Result
	done                   chan struct{}
}

func consumerWorker(c *Consumer) {
	var (
		msg *sarama.ConsumerMessage
		err error
	)

	for {
		select {
		case <-c.done:
			return
		case err = <-c.kafkaPartitionConsumer.Errors():
			c.stream <- result.Result{
				Value: nil,
				Err:   err,
			}
		case msg = <-c.kafkaPartitionConsumer.Messages():
			c.stream <- result.Result{
				Value: msg.Value,
				Err:   nil,
			}
		}
	}
}

func (c *Consumer) Consume() (<-chan result.Result, error) {
	return c.stream, nil
}

func (c *Consumer) Close() error {
	var (
		err error
	)

	defer close(c.done)
	defer close(c.stream)

	err = c.kafkaPartitionConsumer.Close()
	if err != nil {
		c.kafkaConsumer.Close()
		c.client.Close()
		return err
	}

	err = c.kafkaConsumer.Close()
	if err != nil {
		c.client.Close()
		return err
	}

	return c.client.Close()
}

func NewConsumer(c Config, l loggers.Logger) (*Consumer, error) {
	var (
		config                 = c
		bufSize                = config.ConsumerBufferSize
		stream                 chan result.Result
		client                 sarama.Client
		kafkaConsumer          sarama.Consumer
		kafkaPartitionConsumer sarama.PartitionConsumer
		consumer               *Consumer
		err                    error
	)

	if bufSize == 0 {
		bufSize = 1
	}

	stream = make(
		chan result.Result,
		bufSize,
	)

	if config.Kafka == nil {
		config.Kafka = sarama.NewConfig()
		config.Kafka.Consumer.Return.Errors = true
	}

	client, err = sarama.NewClient(
		config.Addrs,
		config.Kafka,
	)
	if err != nil {
		return nil, err
	}

	kafkaConsumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		client.Close()
		return nil, err
	}

	kafkaPartitionConsumer, err = kafkaConsumer.ConsumePartition(
		config.Topic,
		0,
		sarama.OffsetOldest,
	)
	if err != nil {
		kafkaPartitionConsumer.Close()
		client.Close()
		return nil, err
	}

	consumer = &Consumer{
		config:                 c,
		client:                 client,
		kafkaConsumer:          kafkaConsumer,
		kafkaPartitionConsumer: kafkaPartitionConsumer,
		stream:                 stream,
		done:                   make(chan struct{}),
	}

	go consumerWorker(consumer)

	return consumer, nil
}
