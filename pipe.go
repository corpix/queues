package queues

import (
	"io"

	"github.com/corpix/formats"
	"github.com/corpix/loggers"
	"github.com/corpix/stores"

	"github.com/cryptounicorns/queues/consumer"
)

func PipeConsumerToWriterWith(c GenericConfig, fn consumer.PrepareForWriterFn, w io.Writer, l loggers.Logger) error {
	var (
		f   formats.Format
		q   Queue
		cr  consumer.Consumer
		mcr consumer.Generic
		err error
	)

	f, err = formats.New(c.Format)
	if err != nil {
		return err
	}

	q, err = New(c.Queue, l)
	if err != nil {
		return err
	}
	defer q.Close()

	cr, err = q.Consumer()
	if err != nil {
		return err
	}
	defer cr.Close()

	mcr = consumer.NewUnmarshal(
		cr,
		new(interface{}),
		f,
	)
	defer mcr.Close()

	return consumer.PipeToWriterWith(mcr, fn, w)
}

func PipeConsumerToStoreWith(c GenericConfig, fn consumer.PrepareForStoreFn, s stores.Store, l loggers.Logger) error {
	var (
		f   formats.Format
		q   Queue
		cr  consumer.Consumer
		mcr consumer.Generic
		err error
	)

	f, err = formats.New(c.Format)
	if err != nil {
		return err
	}

	q, err = New(c.Queue, l)
	if err != nil {
		return err
	}
	defer q.Close()

	cr, err = q.Consumer()
	if err != nil {
		return err
	}
	defer cr.Close()

	mcr = consumer.NewUnmarshal(
		cr,
		new(interface{}),
		f,
	)
	defer mcr.Close()

	return consumer.PipeToStoreWith(mcr, fn, s)
}
