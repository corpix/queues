package queues

import (
	"strings"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"

	"github.com/corpix/queues/consumer"
	"github.com/corpix/queues/errors"
	"github.com/corpix/queues/producer"
	"github.com/corpix/queues/queue/channel"
	"github.com/corpix/queues/queue/kafka"
	"github.com/corpix/queues/queue/nsq"
)

const (
	// KafkaQueueType is a Config Queue type for Apache Kafka.
	KafkaQueueType = "kafka"

	// NsqQueueType is a Config Queue type for NSQ.
	NsqQueueType = "nsq"

	// ChannelQueueType is a Config Queue type for go channels.
	// Usually used for testing.
	ChannelQueueType = "channel"
)

// Config is a configuration for Queue.
type Config struct {
	Type    string
	Kafka   kafka.Config
	Nsq     nsq.Config
	Channel channel.Config
}

// Queue is a common interface for message queue.
type Queue interface {
	Producer() (producer.Producer, error)
	Consumer() (consumer.Consumer, error)
	Close() error
}

// NewFromConfig creates new Queue from Config.
func NewFromConfig(c Config, l loggers.Logger) (Queue, error) {
	if l == nil {
		return nil, errors.NewErrNilArgument(l)
	}

	switch strings.ToLower(c.Type) {
	case KafkaQueueType:
		return kafka.NewFromConfig(
			c.Kafka,
			prefixwrapper.New(
				loggerPrefix(KafkaQueueType),
				l,
			),
		)
	case NsqQueueType:
		return nsq.NewFromConfig(
			c.Nsq,
			prefixwrapper.New(
				loggerPrefix(NsqQueueType),
				l,
			),
		)
	case ChannelQueueType:
		return channel.NewFromConfig(
			c.Channel,
			prefixwrapper.New(
				loggerPrefix(ChannelQueueType),
				l,
			),
		)
	default:
		return nil, errors.NewErrUnknownQueueType(c.Type)
	}
}
