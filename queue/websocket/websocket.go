package websocket

import (
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
)

type Websocket struct {
	config Config
	log    loggers.Logger
}

func (q *Websocket) Producer() (producer.Producer, error) {
	return NewProducer(q.config, q.log)
}

func (q *Websocket) Consumer() (consumer.Consumer, error) {
	return NewConsumer(q.config, q.log)
}

func (q *Websocket) Close() error {
	return nil
}

func New(c Config, l loggers.Logger) *Websocket {
	return &Websocket{
		config: c,
		log:    l,
	}
}
