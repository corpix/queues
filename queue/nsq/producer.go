package nsq

import (
	nsq "github.com/bitly/go-nsq"
	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"

	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/producer"
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

func NewProducer(c Config, l loggers.Logger) (producer.Producer, error) {
	var (
		log = prefixwrapper.New(
			"NsqProducer: ",
			l,
		)
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
		NewLogger(log),
		c.LogLevel,
	)

	return &Producer{
		topic:       c.Topic,
		nsqProducer: nsqProducer,
	}, nil
}
