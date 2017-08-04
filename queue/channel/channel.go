package channel

import (
	"github.com/corpix/logger"

	"github.com/corpix/queues/errors"
	"github.com/corpix/queues/handler"
	"github.com/corpix/queues/message"
)

type Channel struct {
	logger logger.Logger
	config Config
	feed   chan message.Message
}

func (e *Channel) Produce(m message.Message) error {
	e.feed <- m

	return nil
}

func (e *Channel) Consume(h handler.Handler) error {
	go e.consumeLoop(h)

	return nil
}

func (e *Channel) consumeLoop(h handler.Handler) {
	for m := range e.feed {
		h(m)
	}
}

func (e *Channel) Close() error {
	close(e.feed)

	return nil
}

func NewFromConfig(c Config, l logger.Logger) (*Channel, error) {
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	return &Channel{
		logger: l,
		config: c,
		feed: make(
			chan message.Message,
			c.Capacity,
		),
	}, nil
}
