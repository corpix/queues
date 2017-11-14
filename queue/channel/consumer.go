package channel

import (
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/result"
)

type Consumer struct {
	channel chan message.Message
	stream  chan result.Result
	done    chan struct{}
}

func consumerWorker(c *Consumer) {
	for {
		select {
		case <-c.done:
			return
		case m := <-c.channel:
			c.stream <- result.Result{
				Value: m,
				Err:   nil,
			}
		}
	}
}

func (c *Consumer) Consume() (<-chan result.Result, error) {
	return c.stream, nil
}

func (c *Consumer) Close() error {
	defer close(c.done)
	defer close(c.stream)

	return nil
}

func NewConsumer(channel chan message.Message, l loggers.Logger) (*Consumer, error) {
	var (
		consumer = &Consumer{
			channel: channel,
			stream:  make(chan result.Result),
			done:    make(chan struct{}),
		}
	)

	go consumerWorker(consumer)

	return consumer, nil
}
