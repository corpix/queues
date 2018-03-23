package websocket

import (
	"github.com/corpix/loggers"
	websocketHelpers "github.com/cryptounicorns/websocket"
	"github.com/gorilla/websocket"

	"github.com/cryptounicorns/queues/queue/readwriter"
)

type Producer struct {
	connection *websocket.Conn
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

func ProducerFromConfigWithDialer(d *websocket.Dialer, c Config, l loggers.Logger) (*Producer, error) {
	var (
		conn *websocket.Conn
		rwp  *readwriter.Producer
		err  error
	)

	conn, _, err = d.Dial(c.Addr, nil)
	if err != nil {
		return nil, err
	}

	rwp, err = readwriter.ProducerFromConfig(
		websocketHelpers.NewWriter(conn),
		readwriter.Config{ConsumerBufferSize: c.ConsumerBufferSize},
		l,
	)
	if err != nil {
		return nil, err
	}

	return &Producer{
		connection: conn,
		Producer:   rwp,
	}, nil
}

func ProducerFromConfig(c Config, l loggers.Logger) (*Producer, error) {
	return ProducerFromConfigWithDialer(websocket.DefaultDialer, c, l)
}
