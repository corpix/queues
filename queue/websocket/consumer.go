package websocket

import (
	"github.com/corpix/loggers"
	websocketHelpers "github.com/cryptounicorns/websocket"
	"github.com/gorilla/websocket"

	"github.com/cryptounicorns/queues/queue/readwriter"
)

type Consumer struct {
	connection *websocket.Conn
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

func ConsumerFromConfigWithDialer(d *websocket.Dialer, c Config, l loggers.Logger) (*Consumer, error) {
	var (
		conn *websocket.Conn
		rwp  *readwriter.Consumer
		err  error
	)

	conn, _, err = d.Dial(c.Addr, nil)
	if err != nil {
		return nil, err
	}

	rwp, err = readwriter.ConsumerFromConfig(
		websocketHelpers.NewReader(conn),
		readwriter.Config{ConsumerBufferSize: c.ConsumerBufferSize},
		l,
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		connection: conn,
		Consumer:   rwp,
	}, nil
}

func ConsumerFromConfig(c Config, l loggers.Logger) (*Consumer, error) {
	return ConsumerFromConfigWithDialer(websocket.DefaultDialer, c, l)
}
