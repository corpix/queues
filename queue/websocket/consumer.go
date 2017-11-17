package websocket

import (
	"context"
	"io"

	"github.com/corpix/loggers"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"github.com/cryptounicorns/queues/queue/readwriter"
)

type Consumer struct {
	connection io.ReadWriteCloser
	*readwriter.Consumer
}

func (c *Consumer) Close() error {
	var (
		err error
	)

	err = c.connection.Close()
	if err != nil {
		return err
	}

	return c.Consumer.Close()
}

func NewConsumer(c Config, l loggers.Logger) (*Consumer, error) {
	var (
		r   io.ReadWriteCloser
		rwc *readwriter.Consumer
		err error
	)

	r, _, _, err = ws.DefaultDialer.Dial(
		context.Background(),
		c.Addr,
	)
	if err != nil {
		return nil, err
	}

	rwc, err = readwriter.NewConsumer(
		wsutil.NewReader(r, ws.StateClientSide),
		readwriter.Config{
			ConsumerBufferSize: c.ConsumerBufferSize,
		},
		l,
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		connection: r,
		Consumer:   rwc,
	}, nil
}
