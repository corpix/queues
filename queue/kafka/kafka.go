package kafka

import (
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/errors"
	"github.com/cryptounicorns/queues/producer"
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
