package nsq

import (
	nsq "github.com/bitly/go-nsq"
	"github.com/corpix/logger"

	"github.com/corpix/queues/handler"
	"github.com/corpix/queues/message"
)

//

type Nsq struct {
	logger   logger.Logger
	config   Config
	producer *nsq.Producer
	consumer *nsq.Consumer
}

func (e *Nsq) Produce(m message.Message) error {
	return e.producer.Publish(
		e.config.Topic,
		[]byte(m),
	)
}

func (e *Nsq) Consume(h handler.Handler) error {
	e.consumer.AddHandler(NewHandler(h))
	return e.consumer.ConnectToNSQD(e.config.Addr)
}

func (e *Nsq) Close() error {
	// Will block until complete
	e.producer.Stop()

	// Will NOT block until complete
	// Just initiates graceful shutdown
	e.consumer.Stop()
	<-e.consumer.StopChan

	return nil
}

//

func NewFromConfig(c Config, l logger.Logger) (*Nsq, error) {
	var (
		config = nsq.NewConfig()
		logger = NewLogger(l)

		producer *nsq.Producer
		consumer *nsq.Consumer
		err      error
	)

	producer, err = nsq.NewProducer(c.Addr, config)
	if err != nil {
		return nil, err
	}
	producer.SetLogger(
		logger,
		c.LogLevel,
	)

	consumer, err = nsq.NewConsumer(
		c.Topic,
		c.Channel,
		config,
	)
	if err != nil {
		return nil, err
	}
	consumer.SetLogger(
		logger,
		c.LogLevel,
	)

	return &Nsq{
		l,
		c,
		producer,
		consumer,
	}, nil
}
