package consumer

import (
	"reflect"

	"github.com/corpix/formats"
	"github.com/corpix/logger"

	"github.com/corpix/queues/errors"
	"github.com/corpix/queues/message"
)

// Consumer is an addon for Queue which provides
// a decoder Handler and read-only channel.
type Consumer struct {
	Type   reflect.Type
	Format formats.Format
	feed   chan interface{}
	log    logger.Logger
}

// GetFeed returns a read-only version of the feed
// with messages coming into.
func (c *Consumer) GetFeed() <-chan interface{} {
	return c.feed
}

// Handler is a Handler func which could be used as Queue.Consume(...)
// argument.
func (c *Consumer) Handler(m message.Message) {
	var (
		v   = reflect.New(c.Type).Interface()
		err error
	)

	err = c.Format.Unmarshal(m, v)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.feed <- v
}

// Close implements io.Closer and frees resources.
func (c *Consumer) Close() error {
	close(c.feed)
	return nil
}

// New creates a new Consumer.
// It receives a type t in form `Foo{}` which will be used
// when unmarshaling data from the Queue.
// It receives f which will be used to unmarshal the Message.
// It receives l which will be used in case of any errors
// to log them.
func New(t interface{}, f formats.Format, l logger.Logger) (*Consumer, error) {
	if t == nil {
		return nil, errors.NewErrNilArgument(t)
	}
	if f == nil {
		return nil, errors.NewErrNilArgument(f)
	}
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	return &Consumer{
		Type:   reflect.TypeOf(t),
		Format: f,
		feed:   make(chan interface{}, 128),
		log:    l,
	}, nil
}
