package consumer

import (
	"io"

	"github.com/corpix/stores"

	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/result"
)

type PrepareForWriterFn = func(v interface{}) (message.Message, error)
type PrepareForStoreFn = func(v interface{}) (string, interface{}, error)

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

func PipeToStoreWith(c GenericConsumer, fn PrepareForStoreFn, s stores.Store) error {
	var (
		stream <-chan result.Generic
		k      string
		v      interface{}
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

		k, v, err = fn(r.Value)
		if err != nil {
			return err
		}

		err = s.Set(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
