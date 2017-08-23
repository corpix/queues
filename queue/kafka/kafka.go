package kafka

import (
	"github.com/corpix/loggers"

	"github.com/corpix/queues/consumer"
	"github.com/corpix/queues/errors"
	"github.com/corpix/queues/producer"
)

type Kafka struct {
	config Config
	log    loggers.Logger
}

func (q *Kafka) Producer() (producer.Producer, error) {
	return NewProducer(q.config, q.log)
}

func (q *Kafka) Consumer() (consumer.Consumer, error) {
	return NewConsumer(q.config, q.log)
}

func (q *Kafka) Close() error { return nil }

func NewFromConfig(c Config, l loggers.Logger) (*Kafka, error) {
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	return &Kafka{
		config: c,
		log:    l,
	}, nil
}
