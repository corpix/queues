package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/corpix/loggers"

	"github.com/corpix/queues/errors"
	"github.com/corpix/queues/result"
)

type Consumer struct {
	client                 sarama.Client
	kafkaConsumer          sarama.Consumer
	kafkaPartitionConsumer sarama.PartitionConsumer
	log                    loggers.Logger
	channel                chan result.Result
	done                   chan bool
}

func (c *Consumer) Consume() <-chan result.Result {
	return c.channel
}

func (c *Consumer) consume() {
	var (
		msg *sarama.ConsumerMessage
		err error
	)

	for {
		select {
		case <-c.done:
			return
		case err = <-c.kafkaPartitionConsumer.Errors():
			c.channel <- result.Result{
				Value: nil,
				Err:   err,
			}
		case msg = <-c.kafkaPartitionConsumer.Messages():
			c.channel <- result.Result{
				Value: msg.Value,
				Err:   nil,
			}
		}
	}
}

func (c *Consumer) Close() error {
	var (
		err error
	)

	err = c.kafkaPartitionConsumer.Close()
	if err != nil {
		return err
	}

	err = c.kafkaConsumer.Close()
	if err != nil {
		return err
	}

	err = c.client.Close()
	if err != nil {
		return err
	}

	close(c.done)
	// XXX: c.channel will be GC'ed, not closing it
	// to mitigate write to closed channel in case of race.

	return nil
}

func NewConsumer(c Config, l loggers.Logger) (*Consumer, error) {
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	var (
		client                 sarama.Client
		kafkaConsumer          sarama.Consumer
		kafkaPartitionConsumer sarama.PartitionConsumer
		consumer               *Consumer
		err                    error
	)

	if c.Kafka == nil {
		c.Kafka = sarama.NewConfig()
		c.Kafka.Consumer.Return.Errors = true
	}

	client, err = sarama.NewClient(
		c.Addrs,
		c.Kafka,
	)
	if err != nil {
		return nil, err
	}

	kafkaConsumer, err = sarama.NewConsumerFromClient(
		client,
	)
	if err != nil {
		return nil, err
	}

	kafkaPartitionConsumer, err = kafkaConsumer.ConsumePartition(
		c.Topic,
		0,
		sarama.OffsetOldest,
	)
	if err != nil {
		return nil, err
	}

	consumer = &Consumer{
		client:                 client,
		kafkaConsumer:          kafkaConsumer,
		kafkaPartitionConsumer: kafkaPartitionConsumer,
		log: l,
		channel: make(
			chan result.Result,
			c.ConsumerBufferSize,
		),
		done: make(chan bool),
	}

	go consumer.consume()

	return consumer, nil
}
