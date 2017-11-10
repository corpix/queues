package producer

import (
	"github.com/cryptounicorns/queues/message"
)

type Producer interface {
	Produce(m message.Message) error
	Close() error
}
