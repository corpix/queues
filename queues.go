package queues

import (
	"strings"

	"github.com/corpix/logger"
	"github.com/fatih/structs"

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
		case ChannelQueueType:
			return channel.NewFromConfig(
				v.Value().(channel.Config),
				l,
			)
		}
	}

	return nil, errors.NewErrUnknownQueueType(c.Type)
}
