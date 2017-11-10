package nsq

import (
	nsq "github.com/bitly/go-nsq"
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/errors"
	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/result"
)

type Consumer struct {
	nsqConsumer *nsq.Consumer
	channel     chan result.Result
}

func (c *Consumer) handler(m message.Message) {
	c.channel <- result.Result{
		Value: m,
		Err:   nil,
	}
}

func (c *Consumer) Consume() <-chan result.Result {
	return c.channel
}

func (c *Consumer) Close() error {
	// Will NOT block until complete
	// Just initiates graceful shutdown.
	// So nsq has this StopChan thing.
	c.nsqConsumer.Stop()
	<-c.nsqConsumer.StopChan
	close(c.channel)

	return nil
}

func NewConsumer(c Config, l loggers.Logger) (consumer.Consumer, error) {
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	var (
		nsqConsumer *nsq.Consumer
		consumer    *Consumer
		err         error
	)

	nsqConsumer, err = nsq.NewConsumer(
		c.Topic,
		c.Channel,
		c.Nsq,
	)
	if err != nil {
		return nil, err
	}
	nsqConsumer.SetLogger(
		NewLogger(l),
		c.LogLevel,
	)

	consumer = &Consumer{
		nsqConsumer: nsqConsumer,
		channel: make(
			chan result.Result,
			c.ConsumerBufferSize,
		),
	}

	nsqConsumer.AddHandler(
		NewHandler(consumer.handler),
	)

	err = nsqConsumer.ConnectToNSQD(c.Addr)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
