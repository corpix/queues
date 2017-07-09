package nsq

import (
	"github.com/bitly/go-nsq"
	"github.com/sirupsen/logrus"

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
	e.producer.Stop()
	e.consumer.Stop()
	return nil
}

//

func NewFromConfig(l logger.Logger, c Config) (*Nsq, error) {
	var (
		config      = nsq.NewConfig()
		logger      = NewLogger(l)
		loggerLevel = NewLevel(l.Level().(logrus.Level))

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
		loggerLevel,
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
		loggerLevel,
	)

	return &Nsq{
		l,
		c,
		producer,
		consumer,
	}, nil
}
