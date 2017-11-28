package queues

import (
	"strings"

	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/errors"
	"github.com/cryptounicorns/queues/queue/channel"
	"github.com/cryptounicorns/queues/queue/kafka"
	"github.com/cryptounicorns/queues/queue/nsq"
	"github.com/cryptounicorns/queues/queue/websocket"
)

const (
	// KafkaQueueType is a Config Queue type for Apache Kafka.
	KafkaQueueType = "kafka"

	// NsqQueueType is a Config Queue type for NSQ.
	NsqQueueType = "nsq"

	// ChannelQueueType is a Config Queue type for go channels.
	// Usually used for testing.
	ChannelQueueType = "channel"

	// WebsocketQueueType is a Config Queue type for websockets.
	// Usually used for testing.
	WebsocketQueueType = "websocket"

	// ReadWriterQueueType is a Config Queue type for go read-writers.
	// Usually used for testing and as part in other queues.
	// It could not be constructed from here because it is
	// more low-level(requires io.ReadWriter) so it should be
	// used in other packages.
	ReadWriterQueueType = "readwriter"
)

// Config is a configuration for Queue.
type Config struct {
	Type      string
	Kafka     kafka.Config
	Nsq       nsq.Config
	Channel   channel.Config
	Websocket websocket.Config
}

type GenericConfig struct {
	Format string
	Queue  Config
}

// New creates new Queue from Config.
func New(c Config, l loggers.Logger) (Queue, error) {
	switch strings.ToLower(c.Type) {
	case KafkaQueueType:
		return kafka.New(
			c.Kafka,
			prefixedLogger(KafkaQueueType, l),
		), nil
	case NsqQueueType:
		return nsq.New(
			c.Nsq,
			prefixedLogger(NsqQueueType, l),
		), nil
	case ChannelQueueType:
		return channel.New(
			c.Channel,
			prefixedLogger(ChannelQueueType, l),
		), nil
	case WebsocketQueueType:
		return websocket.New(
			c.Websocket,
			prefixedLogger(WebsocketQueueType, l),
		), nil
	default:
		return nil, errors.NewErrUnknownQueueType(c.Type)
	}
}
