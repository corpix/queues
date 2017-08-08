package channel

import (
	"github.com/corpix/queues/message"
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

func NewProducer(channel chan message.Message) (*Producer, error) {
	return &Producer{
		channel: channel,
	}, nil
}
