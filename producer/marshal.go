package producer

import (
	"github.com/corpix/formats"

	"github.com/cryptounicorns/queues/message"
)

type Marshal struct {
	producer Producer
	format   formats.Format
}

func (p *Marshal) Produce(v interface{}) error {
	var (
		buf message.Message
		err error
	)

	buf, err = p.format.Marshal(v)
	if err != nil {
		return err
	}

	return p.producer.Produce(buf)
}

func (p *Marshal) Close() error {
	return nil
}

func NewMarshal(pr Producer, f formats.Format) *Marshal {
	return &Marshal{
		producer: pr,
		format:   f,
	}
}
