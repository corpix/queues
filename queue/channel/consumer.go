package channel

import (
	"github.com/corpix/queues/message"
)

type Consumer struct {
	channel chan message.Message
}

func (c *Consumer) Consume() <-chan message.Message {
	return c.channel
}

func (c *Consumer) Close() error {
	return nil
}

func NewConsumer(channel chan message.Message) (*Consumer, error) {
	return &Consumer{
		channel: channel,
	}, nil
}
