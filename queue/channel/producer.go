package channel

import (
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/message"
)

type Producer struct {
	channel chan message.Message
}

func (p *Producer) Produce(m message.Message) error {
	p.channel <- m
	return nil
}

func (p *Producer) Close() error {
	return nil
}

func ProducerFromConfig(channel chan message.Message, c Config, l loggers.Logger) (*Producer, error) {
	return &Producer{channel: channel}, nil
}
