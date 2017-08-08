package queues

import (
	"github.com/corpix/formats"
	"github.com/corpix/logger"

	"github.com/corpix/queues/errors"
	"github.com/corpix/queues/producer"
)

// MarshalProducer is an addon for Queue which provides
// a marshaling wrapper around Producer.
type MarshalProducer struct {
	producer producer.Producer
	Format   formats.Format
	log      logger.Logger
}

// Produce marshals a message `m` into Format and produces it
// into the queue.
func (c *MarshalProducer) Produce(m interface{}) error {
	var (
		buf []byte
		err error
	)

	buf, err = c.Format.Marshal(m)
	if err != nil {
		return err
	}

	return c.producer.Produce(buf)

}

// Close implements io.Closer and frees resources.
func (c *MarshalProducer) Close() error {
	return nil
}

// NewMarshalProducer creates a new MarshalProducer.
// It receives a producer which will be used to produce
// a marshaled byte-slice.
// It receives f which will be used to marshal the Message.
// It receives l which will be used in case of any errors
// to log them.
func NewMarshalProducer(p producer.Producer, f formats.Format, l logger.Logger) (*MarshalProducer, error) {
	if p == nil {
		return nil, errors.NewErrNilArgument(p)
	}
	if f == nil {
		return nil, errors.NewErrNilArgument(f)
	}
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	return &MarshalProducer{
		producer: p,
		Format:   f,
		log:      l,
	}, nil
}
