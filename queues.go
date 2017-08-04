package queues

import (
	"strings"

	"github.com/corpix/logger"
	"github.com/fatih/structs"

	"github.com/corpix/queues/errors"
	"github.com/corpix/queues/handler"
	"github.com/corpix/queues/message"
	"github.com/corpix/queues/queue/kafka"
	"github.com/corpix/queues/queue/nsq"
)

const (
	// KafkaQueueType is a Config Queue type for Apache Kafka.
	KafkaQueueType = "kafka"

	// NsqQueueType is a Config Queue type for NSQ.
	NsqQueueType = "nsq"
)

//

// Config is a configuration for Queue.
type Config struct {
	Type  string
	Kafka kafka.Config
	Nsq   nsq.Config
}

//

// Queue is a common interface for message queue.
type Queue interface {
	Produce(message.Message) error
	Consume(handler.Handler) error
	Close() error
}

//

// NewFromConfig creates new Queue from Config.
func NewFromConfig(c Config, l logger.Logger) (Queue, error) {
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	var (
		t = strings.ToLower(c.Type)
	)

	for _, v := range structs.New(c).Fields() {
		if strings.ToLower(v.Name()) != t {
			continue
		}

		switch t {
		case KafkaQueueType:
			return kafka.NewFromConfig(
				v.Value().(kafka.Config),
				l,
			)
		case NsqQueueType:
			return nsq.NewFromConfig(
				v.Value().(nsq.Config),
				l,
			)
		}
	}

	return nil, errors.NewErrUnknownQueueType(c.Type)
}
