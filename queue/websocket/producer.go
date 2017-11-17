package websocket

import (
	"context"
	"io"

	"github.com/corpix/loggers"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"github.com/cryptounicorns/queues/queue/readwriter"
)

type Producer struct {
	connection io.ReadWriteCloser
	*readwriter.Producer
}

func (p *Producer) Close() error {
	var (
		err error
	)

	err = p.connection.Close()
	if err != nil {
		return err
	}

	return p.Producer.Close()
}

func NewProducer(c Config, l loggers.Logger) (*Producer, error) {
	var (
		r   io.ReadWriteCloser
		rwp *readwriter.Producer
		err error
	)

	r, _, _, err = ws.DefaultDialer.Dial(
		context.Background(),
		c.Addr,
	)
	if err != nil {
		return nil, err
	}

	rwp, err = readwriter.NewProducer(
		wsutil.NewWriter(
			r,
			ws.StateClientSide,
			ws.OpBinary,
		),
		readwriter.Config{
			ConsumerBufferSize: c.ConsumerBufferSize,
		},
		l,
	)
	if err != nil {
		return nil, err
	}

	return &Producer{
		connection: r,
		Producer:   rwp,
	}, nil
}
