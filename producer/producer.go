package producer

import (
	"github.com/corpix/queues/message"
)

type Producer interface {
	Produce(m message.Message) error
	Close() error
}
