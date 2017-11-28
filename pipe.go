package queues

import (
	"context"
	"io"

	"github.com/corpix/effects/closer"
	"github.com/corpix/formats"
	"github.com/corpix/loggers"
	"github.com/corpix/stores"

	"github.com/cryptounicorns/queues/consumer"
)

func PipeConsumerToWriterWith(c GenericConfig, ctx context.Context, fn consumer.PrepareForWriterFn, w io.Writer, l loggers.Logger) error {
	var (
		closers = closer.Closers{}
		f       formats.Format
		q       Queue
		cr      consumer.Consumer
		mcr     consumer.Generic
		err     error
	)

	f, err = formats.New(c.Format)
	if err != nil {
		return err
	}

	q, err = New(c.Queue, l)
	if err != nil {
		return err
	}
	closers = append(closers, q)

	cr, err = q.Consumer()
	if err != nil {
		return err
	}
	closers = append(closers, cr)

	mcr = consumer.NewUnmarshal(
		cr,
		new(interface{}),
		f,
	)
	closers = append(closers, mcr)

	go func() {
		select {
		case <-ctx.Done():
			err := closers.Close()
			if err != nil {
				l.Error(err)
			}
			return
		}
	}()

	return consumer.PipeToWriterWith(mcr, fn, w)
}

func PipeConsumerToStoreWith(c GenericConfig, ctx context.Context, fn consumer.PrepareForStoreFn, s stores.Store, l loggers.Logger) error {
	var (
		closers = closer.Closers{}
		f       formats.Format
		q       Queue
		cr      consumer.Consumer
		mcr     consumer.Generic
		err     error
	)

	f, err = formats.New(c.Format)
	if err != nil {
		return err
	}

	q, err = New(c.Queue, l)
	if err != nil {
		return err
	}
	closers = append(closers, q)

	cr, err = q.Consumer()
	if err != nil {
		return err
	}
	closers = append(closers, cr)

	mcr = consumer.NewUnmarshal(
		cr,
		new(interface{}),
		f,
	)
	closers = append(closers, mcr)

	go func() {
		select {
		case <-ctx.Done():
			err := closers.Close()
			if err != nil {
				l.Error(err)
			}
			return
		}
	}()

	return consumer.PipeToStoreWith(mcr, fn, s)
}
