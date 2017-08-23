package consumer

import (
	"reflect"

	"github.com/corpix/formats"
	"github.com/corpix/loggers"

	"github.com/corpix/queues/errors"
)

// UnmarshalConsumer is an addon for Queue which provides
// a unmarshaling wrapper around Consumer.
type UnmarshalConsumer struct {
	Type     reflect.Type
	consumer Consumer
	Format   formats.Format
	channel  chan interface{}
	log      loggers.Logger
}

// Consume returns a read-only version of the channel
// with messages coming into.
// It tries to mimic Consumer interface.
// (but it can't deliver a better UX coz go has no generics ¯\_(ツ)_/¯).
func (c *UnmarshalConsumer) Consume() <-chan interface{} {
	return c.channel
}

func (c *UnmarshalConsumer) consume() {
	var (
		v   interface{}
		err error
	)

	for m := range c.consumer.Consume() {
		v = reflect.New(c.Type).Interface()

		err = c.Format.Unmarshal(m, v)
		if err != nil {
			c.log.Error(err)
			continue
		}

		c.channel <- v
	}
}

// Close implements io.Closer and frees resources.
func (c *UnmarshalConsumer) Close() error {
	close(c.channel)

	return nil
}

// NewUnmarshalConsumer creates a new UnmarshalConsumer.
// It receives a type t in form `Foo{}` which will be used
// when unmarshaling data from the Queue.
// It receives a consumer which will be used to consume initial
// byte-slice.
// It receives f which will be used to unmarshal the Message.
// It receives l which will be used in case of any errors
// to log them.
func NewUnmarshalConsumer(t interface{}, c Consumer, f formats.Format, l loggers.Logger) (*UnmarshalConsumer, error) {
	if t == nil {
		return nil, errors.NewErrNilArgument(t)
	}
	if c == nil {
		return nil, errors.NewErrNilArgument(c)
	}
	if f == nil {
		return nil, errors.NewErrNilArgument(f)
	}
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	consumer := &UnmarshalConsumer{
		Type:     reflect.TypeOf(t),
		consumer: c,
		Format:   f,
		channel:  make(chan interface{}, 128),
		log:      l,
	}

	go consumer.consume()

	return consumer, nil
}
