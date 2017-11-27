package producer

import (
	"github.com/cryptounicorns/queues/message"
)

type Producer interface {
	Produce(m message.Message) error
	Close() error
}

type Generic interface {
	Produce(m interface{}) error
	Close() error
}
