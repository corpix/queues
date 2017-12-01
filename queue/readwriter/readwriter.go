package readwriter

import (
	"io"

	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
)

const (
	Name = "readwriter"
)

type ReadWriter struct {
	config Config
	reader io.Reader
	writer io.Writer
	log    loggers.Logger
}

func (q *ReadWriter) Producer() (producer.Producer, error) {
	return NewProducer(q.writer, q.config, q.log)
}

func (q *ReadWriter) Consumer() (consumer.Consumer, error) {
	return NewConsumer(q.reader, q.config, q.log)
}

func (q *ReadWriter) Close() error {
	return nil
}

func New(rw io.ReadWriter, c Config, l loggers.Logger) *ReadWriter {
	return &ReadWriter{
		config: c,
		reader: rw,
		writer: rw,
		log:    l,
	}
}
