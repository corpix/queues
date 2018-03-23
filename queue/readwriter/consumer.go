package readwriter

import (
	"io"
	"io/ioutil"

	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/result"
)

type Consumer struct {
	reader io.Reader
	stream chan result.Result
	done   chan struct{}
}

func consumerWorker(c *Consumer) {
	var (
		buf []byte
		err error
	)

	for {
		buf, err = ioutil.ReadAll(c.reader)
		// FIXME: Send me an empty buffer and I will
		// return you EOF. You can't know where was a
		// network connection reset.
		if err == nil && len(buf) == 0 {
			err = io.EOF
		}

		select {
		case <-c.done:
			return
		default:
			c.stream <- result.Result{
				Value: buf,
				Err:   err,
			}
			if err != nil {
				return
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

func ConsumerFromConfig(r io.Reader, c Config, l loggers.Logger) (*Consumer, error) {
	var (
		bufSize  = c.ConsumerBufferSize
		consumer *Consumer
	)

	if bufSize == 0 {
		bufSize = 1
	}

	consumer = &Consumer{
		reader: r,
		stream: make(
			chan result.Result,
			bufSize,
		),
		done: make(chan struct{}),
	}

	go consumerWorker(consumer)

	return consumer, nil
}
