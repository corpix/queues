package nsq

import (
	nsq "github.com/bitly/go-nsq"
	"github.com/corpix/logger"

	"github.com/corpix/queues/errors"
	"github.com/corpix/queues/message"
	"github.com/corpix/queues/producer"
)

type Producer struct {
	topic       string
	nsqProducer *nsq.Producer
}

func (p *Producer) Produce(m message.Message) error {
	return p.nsqProducer.Publish(
		p.topic,
		m,
	)
}

func (p *Producer) Close() error {
	p.nsqProducer.Stop()

	return nil
}

func NewProducer(c Config, l logger.Logger) (producer.Producer, error) {
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	var (
		nsqProducer *nsq.Producer
		err         error
	)

	nsqProducer, err = nsq.NewProducer(
		c.Addr,
		c.Nsq,
	)
	if err != nil {
		return nil, err
	}

	nsqProducer.SetLogger(
		NewLogger(l),
		c.LogLevel,
	)

	return &Producer{
		topic:       c.Topic,
		nsqProducer: nsqProducer,
	}, nil
}
