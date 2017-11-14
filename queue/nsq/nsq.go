package nsq

import (
	nsq "github.com/bitly/go-nsq"
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
)

type Nsq struct {
	config Config
	log    loggers.Logger
}

func (q *Nsq) Producer() (producer.Producer, error) {
	return NewProducer(q.config, q.log)
}

func (q *Nsq) Consumer() (consumer.Consumer, error) {
	return NewConsumer(q.config, q.log)
}

func (q *Nsq) Close() error { return nil }

func New(c Config, l loggers.Logger) *Nsq {
	if c.Nsq == nil {
		c.Nsq = nsq.NewConfig()
	}

	return &Nsq{
		config: c,
		log:    l,
	}
}
