package channel

import (
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/producer"
)

const (
	Name = "channel"
)

type Channel struct {
	config  Config
	log     loggers.Logger
	channel chan message.Message
}

func (q *Channel) Producer() (producer.Producer, error) {
	return NewProducer(q.channel, q.config, q.log)
}

func (q *Channel) Consumer() (consumer.Consumer, error) {
	return NewConsumer(q.channel, q.config, q.log)
}

func (q *Channel) Close() error {
	close(q.channel)

	return nil
}

func New(c Config, l loggers.Logger) *Channel {
	return &Channel{
		config: c,
		log:    l,
		channel: make(
			chan message.Message,
			c.Capacity,
		),
	}
}
