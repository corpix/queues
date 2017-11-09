package channel

import (
	"github.com/corpix/queues/message"
	"github.com/corpix/queues/result"
)

type Consumer struct {
	channel chan message.Message
	done    chan struct{}
}

func (c *Consumer) Consume() <-chan result.Result {
	var (
		stream = make(chan result.Result)
	)

	go func() {
		for {
			select {
			case <-c.done:
				return
			case m := <-c.channel:
				stream <- result.Result{
					Value: m,
					Err:   nil,
				}
			}
		}
	}()

	return stream
}

func (c *Consumer) Close() error {
	close(c.done)
	return nil
}

func NewConsumer(channel chan message.Message) (*Consumer, error) {
	return &Consumer{
		channel: channel,
		done:    make(chan struct{}),
	}, nil
}
