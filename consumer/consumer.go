package consumer

import (
	"github.com/corpix/queues/message"
)

type Consumer interface {
	Consume() <-chan message.Message
	Close() error
}
