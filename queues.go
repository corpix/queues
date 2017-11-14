package queues

import (
	"strings"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"

	"github.com/cryptounicorns/queues/errors"
	"github.com/cryptounicorns/queues/queue/channel"
	"github.com/cryptounicorns/queues/queue/kafka"
	"github.com/cryptounicorns/queues/queue/nsq"
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

// New creates new Queue from Config.
func New(c Config, l loggers.Logger) (Queue, error) {
	switch strings.ToLower(c.Type) {
	case KafkaQueueType:
		return kafka.New(
			c.Kafka,
			prefixwrapper.New(
				loggerPrefix(KafkaQueueType),
				l,
			),
		), nil
	case NsqQueueType:
		return nsq.New(
			c.Nsq,
			prefixwrapper.New(
				loggerPrefix(NsqQueueType),
				l,
			),
		), nil
	case ChannelQueueType:
		return channel.New(
			c.Channel,
			prefixwrapper.New(
				loggerPrefix(ChannelQueueType),
				l,
			),
		), nil
	default:
		return nil, errors.NewErrUnknownQueueType(c.Type)
	}
}
