package consumer

import (
	"github.com/cryptounicorns/queues/result"
)

type Consumer interface {
	Consume() (<-chan result.Result, error)
	Close() error
}

type GenericConsumer interface {
	Consume() (<-chan result.Generic, error)
	Close() error
}
