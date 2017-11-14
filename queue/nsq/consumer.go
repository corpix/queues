package nsq

import (
	nsq "github.com/bitly/go-nsq"
	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/result"
)

type Consumer struct {
	config      Config
	nsqConsumer *nsq.Consumer
	log         loggers.Logger
	stream      chan result.Result
	done        chan struct{}
}

func consumerHandler(channel chan result.Result) Handler {
	return func(m message.Message) {
		channel <- result.Result{
			Value: m,
			Err:   nil,
		}
	}
}

func (c *Consumer) Consume() (<-chan result.Result, error) {
	return c.stream, nil
}

func (c *Consumer) Close() error {
	defer close(c.done)
	defer close(c.stream)

	// Will NOT block until complete
	// Just initiates graceful shutdown.
	// So nsq has this StopChan thing.
	c.nsqConsumer.Stop()
	<-c.nsqConsumer.StopChan

	return nil
}

func NewConsumer(c Config, l loggers.Logger) (consumer.Consumer, error) {
	var (
		bufSize     = c.ConsumerBufferSize
		concurrency = c.Concurrency
		stream      chan result.Result
		log         = prefixwrapper.New(
			"NsqConsumer: ",
			l,
		)
		nsqConsumer *nsq.Consumer
		err         error
	)

	if bufSize == 0 {
		bufSize = 1
	}
	if concurrency == 0 {
		concurrency = 1
	}

	stream = make(
		chan result.Result,
		bufSize,
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
		NewLogger(log),
		c.LogLevel,
	)

	nsqConsumer.AddConcurrentHandlers(
		NewHandler(consumerHandler(stream)),
		int(concurrency),
	)

	err = nsqConsumer.ConnectToNSQD(c.Addr)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		config:      c,
		nsqConsumer: nsqConsumer,
		log:         log,
		stream:      stream,
		done:        make(chan struct{}),
	}, nil
}
