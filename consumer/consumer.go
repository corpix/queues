package consumer

import (
	"github.com/cryptounicorns/queues/result"
)

type Consumer interface {
	Consume() (<-chan result.Result, error)
	Close() error
}
