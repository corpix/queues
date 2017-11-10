package queues

import (
	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
)

// Queue is a common interface for message queue.
type Queue interface {
	Producer() (producer.Producer, error)
	Consumer() (consumer.Consumer, error)
	Close() error
}
