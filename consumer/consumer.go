package consumer

import (
	"github.com/corpix/queues/result"
)

type Consumer interface {
	Consume() <-chan result.Result
	Close() error
}
