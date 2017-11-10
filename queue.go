package queues

import (
	"github.com/corpix/queues/consumer"
	"github.com/corpix/queues/producer"
)

// Queue is a common interface for message queue.
type Queue interface {
	Producer() (producer.Producer, error)
	Consumer() (consumer.Consumer, error)
	Close() error
}
