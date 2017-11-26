package consumer

import (
	"io"

	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/result"
)

type PrepareForWriterFn = func(v interface{}) (message.Message, error)

func PipeToWriter(c Consumer, w io.Writer) error {
	var (
		stream <-chan result.Result
		err    error
	)

	stream, err = c.Consume()
	if err != nil {
		return err
	}

	for r := range stream {
		if r.Err != nil {
			return r.Err
		}

		_, err = w.Write(r.Value)
		if err != nil {
			return nil
		}
	}

	return nil
}

func PipeToWriterWith(c GenericConsumer, fn PrepareForWriterFn, w io.Writer) error {
	var (
		stream <-chan result.Generic
		buf    message.Message
		err    error
	)

	stream, err = c.Consume()
	if err != nil {
		return err
	}

	for r := range stream {
		if r.Err != nil {
			return r.Err
		}

		buf, err = fn(r.Value)
		if err != nil {
			return err
		}

		_, err = w.Write(buf)
		if err != nil {
			return nil
		}
	}

	return nil
}
