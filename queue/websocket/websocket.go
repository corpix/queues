package websocket

import (
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
)

const (
	Name = "websocket"
)

type Websocket struct {
	config Config
	log    loggers.Logger
}

func (q *Websocket) Producer() (producer.Producer, error) {
	return ProducerFromConfig(q.config, q.log)
}

func (q *Websocket) Consumer() (consumer.Consumer, error) {
	return ConsumerFromConfig(q.config, q.log)
}

func (q *Websocket) Close() error {
	return nil
}

func FromConfig(c Config, l loggers.Logger) *Websocket {
	return &Websocket{
		config: c,
		log:    l,
	}
}
