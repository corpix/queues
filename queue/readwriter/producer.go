package readwriter

import (
	"io"

	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/message"
)

type Producer struct {
	writer io.Writer
}

func (p *Producer) Produce(m message.Message) error {
	var (
		err error
	)

	_, err = p.writer.Write(m)
	return err
}

func (p *Producer) Close() error {
	return nil
}

func NewProducer(w io.Writer, c Config, l loggers.Logger) (*Producer, error) {
	return &Producer{
		writer: w,
	}, nil
}
