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

func NewConsumer(channel chan message.Message, c Config, l loggers.Logger) (*Consumer, error) {
	var (
		capacity = c.Capacity
		consumer *Consumer
	)

	if capacity == 0 {
		capacity = 1
	}

	consumer = &Consumer{
		channel: channel,
		stream: make(
			chan result.Result,
			capacity,
		),
		done: make(chan struct{}),
	}

	go consumerWorker(consumer)

	return consumer, nil
}
